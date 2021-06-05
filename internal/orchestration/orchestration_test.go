package orchestration_test

import (
	"net"
	"testing"
	"time"

	"github.com/relab/gorums"
	"github.com/relab/hotstuff/internal/orchestration"
	"github.com/relab/hotstuff/internal/proto/orchestrationpb"
)

func TestOrchestration(t *testing.T) {
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatal(err)
	}

	worker := orchestration.NewWorker()
	srv := gorums.NewServer()
	orchestrationpb.RegisterOrchestratorServer(srv, worker)
	go func() {
		err := srv.Serve(lis)
		if err != nil {
			panic(err)
		}
	}()

	addr := lis.Addr().String()
	experiment := &orchestration.Experiment{
		NumReplicas:    4,
		NumClients:     1,
		BatchSize:      100,
		PayloadSize:    100,
		Duration:       5 * time.Second,
		Consensus:      "chainedhotstuff",
		Crypto:         "ecdsa",
		LeaderRotation: "round-robin",
	}

	err = experiment.Run([]string{addr})
	if err != nil {
		t.Fatal(err)
	}
}
