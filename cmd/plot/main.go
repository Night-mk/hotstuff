package main

import (
	"bufio"
	"fmt"
	"image/color"
	"io/ioutil"
	"math"
	"os"
	"path"
	"sort"
	"strings"
	"time"

	"go-hep.org/x/hep/hplot"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

type rawThroughput struct {
	deltaTime   float64
	numCommands float64
}

type measurement struct {
	rate       int
	latency    float64
	throughput float64
}

type benchmark struct {
	measurements []measurement
	name         string
	batchSize    int
	payloadSize  int
}

func (b benchmark) String() string {
	var ret strings.Builder
	ret.WriteString(fmt.Sprintf("%s-b%d-p%d: ", b.name, b.batchSize, b.payloadSize))
	for _, m := range b.measurements {
		ret.WriteString(fmt.Sprintf("(%f, %f), ", m.throughput, m.latency))
	}
	return ret.String()
}

func (b *benchmark) Len() int {
	return len(b.measurements)
}

func (b *benchmark) XY(i int) (x, y float64) {
	m := b.measurements[i]
	return m.throughput, m.latency
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s [paths to benchmark files] [output file]\n", os.Args[0])
		os.Exit(1)
	}

	p, err := plot.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create plot: %v\n", err)
		os.Exit(1)
	}

	grid := plotter.NewGrid()
	grid.Horizontal.Color = color.Gray{Y: 200}
	grid.Horizontal.Dashes = plotutil.Dashes(2)
	grid.Vertical.Color = color.Gray{Y: 200}
	grid.Vertical.Dashes = plotutil.Dashes(2)
	p.Add(grid)

	var plots []interface{}
	benchmarks := make(chan *benchmark, 1)
	errors := make(chan error, 1)
	n := 0

	for _, benchFolder := range os.Args[1 : len(os.Args)-1] {
		dir, err := ioutil.ReadDir(benchFolder)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to open benchmark directory: %v\n", err)
			os.Exit(1)
		}

		for _, f := range dir {
			if f.IsDir() {
				go func(dir string) {
					b, err := processBenchmark(dir)
					benchmarks <- b
					errors <- err
				}(path.Join(benchFolder, f.Name()))
				n++
			}
		}
	}

	for i := 0; i < n; i++ {
		err := <-errors
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to read benchmark: %v\n", err)
			os.Exit(1)
		}

		b := <-benchmarks
		label := fmt.Sprintf("%s-b%d-p%d", b.name, b.batchSize, b.payloadSize)

		// insert sorted for deterministic output
		index := len(plots)
		for j, v := range plots {
			if s, ok := v.(string); ok {
				if label > s {
					index = j
				}
			}
		}
		// grow size by 2
		plots = append(plots, struct{}{}, struct{}{})
		copy(plots[index+2:], plots[index:])
		plots[index] = label
		plots[index+1] = b
	}

	err = plotutil.AddLinePoints(p, plots...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to add plots: %v\n", err)
		os.Exit(1)
	}

	/* p.Legend.Left = true */
	p.Legend.Top = true
	p.X.Label.Text = "Throughput Kops/sec"
	p.X.Tick.Marker = hplot.Ticks{N: 10}
	p.Y.Label.Text = "Latency ms"
	p.Y.Tick.Marker = hplot.Ticks{N: 10}

	if err := p.Save(6*vg.Inch, 6*vg.Inch, os.Args[len(os.Args)-1]); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to save plot: %v\n", err)
		os.Exit(1)
	}
}

func processBenchmark(dirPath string) (*benchmark, error) {
	measurements := make(map[int][]measurement)
	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}
	b := &benchmark{}
	if strings.HasPrefix(path.Base(dirPath), "lhs-") {
		b.name = "libhotstuff"
	} else {
		b.name = "relab/hotstuff"
	}
	fmt.Sscanf(strings.TrimPrefix(path.Base(dirPath), "lhs-"), "b%d-p%d", &b.batchSize, &b.payloadSize)
	for _, f := range dir {
		if f.IsDir() {
			if err := processRun(path.Join(dirPath, f.Name()), measurements, b.name); err != nil {
				return nil, err
			}
		}
	}
	b.measurements = make([]measurement, 0, len(measurements))
	for _, ms := range measurements {
		latencyTotal, throughputTotal := 0.0, 0.0
		rate := ms[0].rate
		for _, m := range ms {
			if m.rate != rate {
				panic("Rate mismatch while combining measurements!")
			}
			latencyTotal += m.latency
			throughputTotal += m.throughput
		}
		m := measurement{
			rate:       rate,
			latency:    latencyTotal / float64(len(ms)),
			throughput: throughputTotal / float64(len(ms)),
		}
		i := sort.Search(len(b.measurements), func(i int) bool {
			return m.rate < b.measurements[i].rate
		})
		b.measurements = append(b.measurements, measurement{})
		copy(b.measurements[i+1:], b.measurements[i:])
		b.measurements[i] = m
	}
	return b, nil
}

func processRun(dirPath string, measurements map[int][]measurement, benchType string) error {
	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
	}

	ms := make(chan measurement, len(dir))
	errs := make(chan error, len(dir))
	n := 0

	for _, f := range dir {
		if f.IsDir() {
			go func(dirPath string) {
				var (
					m   measurement
					err error
				)

				if benchType == "libhotstuff" {
					m, err = readInMeasurement(dirPath, false)
				} else {
					m, err = readInMeasurement(dirPath, true)
				}

				fmt.Sscanf(path.Base(dirPath), "t%d", &m.rate)

				errs <- err
				ms <- m

			}(path.Join(dirPath, f.Name()))
			n++
		}
	}

	for i := 0; i < n; i++ {
		err := <-errs
		if err != nil {
			fmt.Fprintf(os.Stderr, "WARNING: error while processing measurement (ignoring): %v\n", err)
			<-ms
			continue
		}
		m := <-ms
		s, _ := measurements[m.rate]
		s = append(s, m)
		measurements[m.rate] = s
	}

	return nil
}

func readInMeasurement(dirPath string, nanoseconds bool) (measurement, error) {
	var numberFormat string
	if nanoseconds {
		numberFormat = "%sns"
	} else {
		numberFormat = "%sus"
	}

	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return measurement{}, err
	}

	var measurements []measurement
	for _, file := range dir {
		if file.IsDir() {
			m, err := readInMeasurement(path.Join(dirPath, file.Name()), nanoseconds)
			if err != nil {
				return measurement{}, err
			}
			measurements = append(measurements, m)
			continue
		}

		f, err := os.Open(path.Join(dirPath, file.Name()))
		if err != nil {
			return measurement{}, fmt.Errorf("Failed to read libhotstuff measurement: %w", err)
		}

		var totalLatency time.Duration
		var totalTime time.Duration
		numCommands := 0
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			l := scanner.Text()
			matches := strings.Split(l, " ")
			if len(matches) < 2 {
				continue
			}

			t, err := time.ParseDuration(fmt.Sprintf("%sns", matches[0]))
			if err != nil {
				return measurement{}, fmt.Errorf("Failed to read libhotstuff measurement: %w", err)
			}
			lat, err := time.ParseDuration(fmt.Sprintf(numberFormat, matches[1]))
			if err != nil {
				return measurement{}, fmt.Errorf("Failed to read libhotstuff measurement: %w", err)
			}
			numCommands++
			totalLatency += lat
			totalTime += t
		}
		latency := (float64(totalLatency) / float64(numCommands)) / float64(time.Millisecond)
		throughput := (float64(numCommands) / 1000) / totalTime.Seconds()
		measurements = append(measurements, measurement{latency: latency, throughput: throughput})
	}

	if len(measurements) == 0 {
		return measurement{}, fmt.Errorf("No measurements")
	}

	sumLatencies := 0.0
	sumThroughput := 0.0
	for _, m := range measurements {
		if math.IsNaN(m.throughput) || math.IsNaN(m.latency) {
			fmt.Fprintf(os.Stderr, "WARNING: Ignoring NaN measurement in '%s'\n", dirPath)
			continue
		}
		sumLatencies += m.latency
		sumThroughput += m.throughput
	}

	return measurement{
		latency:    sumLatencies / float64(len(measurements)),
		throughput: sumThroughput,
	}, nil
}
