// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/relab/hotstuff (interfaces: Consensus)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	hotstuff "github.com/relab/hotstuff"
	reflect "reflect"
)

// MockConsensus is a mock of Consensus interface
type MockConsensus struct {
	ctrl     *gomock.Controller
	recorder *MockConsensusMockRecorder
}

// MockConsensusMockRecorder is the mock recorder for MockConsensus
type MockConsensusMockRecorder struct {
	mock *MockConsensus
}

// NewMockConsensus creates a new mock instance
func NewMockConsensus(ctrl *gomock.Controller) *MockConsensus {
	mock := &MockConsensus{ctrl: ctrl}
	mock.recorder = &MockConsensusMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConsensus) EXPECT() *MockConsensusMockRecorder {
	return m.recorder
}

// BlockChain mocks base method
func (m *MockConsensus) BlockChain() hotstuff.BlockChain {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockChain")
	ret0, _ := ret[0].(hotstuff.BlockChain)
	return ret0
}

// BlockChain indicates an expected call of BlockChain
func (mr *MockConsensusMockRecorder) BlockChain() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockChain", reflect.TypeOf((*MockConsensus)(nil).BlockChain))
}

// Config mocks base method
func (m *MockConsensus) Config() hotstuff.Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(hotstuff.Config)
	return ret0
}

// Config indicates an expected call of Config
func (mr *MockConsensusMockRecorder) Config() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockConsensus)(nil).Config))
}

// CreateDummy mocks base method
func (m *MockConsensus) CreateDummy() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CreateDummy")
}

// CreateDummy indicates an expected call of CreateDummy
func (mr *MockConsensusMockRecorder) CreateDummy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDummy", reflect.TypeOf((*MockConsensus)(nil).CreateDummy))
}

// HighQC mocks base method
func (m *MockConsensus) HighQC() hotstuff.QuorumCert {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HighQC")
	ret0, _ := ret[0].(hotstuff.QuorumCert)
	return ret0
}

// HighQC indicates an expected call of HighQC
func (mr *MockConsensusMockRecorder) HighQC() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HighQC", reflect.TypeOf((*MockConsensus)(nil).HighQC))
}

// LastVote mocks base method
func (m *MockConsensus) LastVote() hotstuff.View {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastVote")
	ret0, _ := ret[0].(hotstuff.View)
	return ret0
}

// LastVote indicates an expected call of LastVote
func (mr *MockConsensusMockRecorder) LastVote() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastVote", reflect.TypeOf((*MockConsensus)(nil).LastVote))
}

// Leaf mocks base method
func (m *MockConsensus) Leaf() *hotstuff.Block {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Leaf")
	ret0, _ := ret[0].(*hotstuff.Block)
	return ret0
}

// Leaf indicates an expected call of Leaf
func (mr *MockConsensusMockRecorder) Leaf() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Leaf", reflect.TypeOf((*MockConsensus)(nil).Leaf))
}

// NewView mocks base method
func (m *MockConsensus) NewView() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NewView")
}

// NewView indicates an expected call of NewView
func (mr *MockConsensusMockRecorder) NewView() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewView", reflect.TypeOf((*MockConsensus)(nil).NewView))
}

// OnNewView mocks base method
func (m *MockConsensus) OnNewView(arg0 hotstuff.QuorumCert) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnNewView", arg0)
}

// OnNewView indicates an expected call of OnNewView
func (mr *MockConsensusMockRecorder) OnNewView(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnNewView", reflect.TypeOf((*MockConsensus)(nil).OnNewView), arg0)
}

// OnPropose mocks base method
func (m *MockConsensus) OnPropose(arg0 *hotstuff.Block) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnPropose", arg0)
}

// OnPropose indicates an expected call of OnPropose
func (mr *MockConsensusMockRecorder) OnPropose(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnPropose", reflect.TypeOf((*MockConsensus)(nil).OnPropose), arg0)
}

// OnVote mocks base method
func (m *MockConsensus) OnVote(arg0 hotstuff.PartialCert) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnVote", arg0)
}

// OnVote indicates an expected call of OnVote
func (mr *MockConsensusMockRecorder) OnVote(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnVote", reflect.TypeOf((*MockConsensus)(nil).OnVote), arg0)
}

// Propose mocks base method
func (m *MockConsensus) Propose() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Propose")
}

// Propose indicates an expected call of Propose
func (mr *MockConsensusMockRecorder) Propose() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Propose", reflect.TypeOf((*MockConsensus)(nil).Propose))
}
