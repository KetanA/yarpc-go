// Code generated by MockGen. DO NOT EDIT.
// Source: go.uber.org/yarpc/api/transport (interfaces: UnaryOutbound,OnewayOutbound,StreamOutbound)

// Package transporttest is a generated GoMock package.
package transporttest

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	transport "go.uber.org/yarpc/api/transport"
	reflect "reflect"
)

// MockUnaryOutbound is a mock of UnaryOutbound interface
type MockUnaryOutbound struct {
	ctrl     *gomock.Controller
	recorder *MockUnaryOutboundMockRecorder
}

// MockUnaryOutboundMockRecorder is the mock recorder for MockUnaryOutbound
type MockUnaryOutboundMockRecorder struct {
	mock *MockUnaryOutbound
}

// NewMockUnaryOutbound creates a new mock instance
func NewMockUnaryOutbound(ctrl *gomock.Controller) *MockUnaryOutbound {
	mock := &MockUnaryOutbound{ctrl: ctrl}
	mock.recorder = &MockUnaryOutboundMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUnaryOutbound) EXPECT() *MockUnaryOutboundMockRecorder {
	return m.recorder
}

// Call mocks base method
func (m *MockUnaryOutbound) Call(arg0 context.Context, arg1 *transport.Request) (*transport.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Call", arg0, arg1)
	ret0, _ := ret[0].(*transport.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Call indicates an expected call of Call
func (mr *MockUnaryOutboundMockRecorder) Call(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Call", reflect.TypeOf((*MockUnaryOutbound)(nil).Call), arg0, arg1)
}

// IsRunning mocks base method
func (m *MockUnaryOutbound) IsRunning() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsRunning")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsRunning indicates an expected call of IsRunning
func (mr *MockUnaryOutboundMockRecorder) IsRunning() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsRunning", reflect.TypeOf((*MockUnaryOutbound)(nil).IsRunning))
}

// Start mocks base method
func (m *MockUnaryOutbound) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockUnaryOutboundMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockUnaryOutbound)(nil).Start))
}

// Stop mocks base method
func (m *MockUnaryOutbound) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockUnaryOutboundMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockUnaryOutbound)(nil).Stop))
}

// Transports mocks base method
func (m *MockUnaryOutbound) Transports() []transport.Transport {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transports")
	ret0, _ := ret[0].([]transport.Transport)
	return ret0
}

// Transports indicates an expected call of Transports
func (mr *MockUnaryOutboundMockRecorder) Transports() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transports", reflect.TypeOf((*MockUnaryOutbound)(nil).Transports))
}

// MockOnewayOutbound is a mock of OnewayOutbound interface
type MockOnewayOutbound struct {
	ctrl     *gomock.Controller
	recorder *MockOnewayOutboundMockRecorder
}

// MockOnewayOutboundMockRecorder is the mock recorder for MockOnewayOutbound
type MockOnewayOutboundMockRecorder struct {
	mock *MockOnewayOutbound
}

// NewMockOnewayOutbound creates a new mock instance
func NewMockOnewayOutbound(ctrl *gomock.Controller) *MockOnewayOutbound {
	mock := &MockOnewayOutbound{ctrl: ctrl}
	mock.recorder = &MockOnewayOutboundMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOnewayOutbound) EXPECT() *MockOnewayOutboundMockRecorder {
	return m.recorder
}

// CallOneway mocks base method
func (m *MockOnewayOutbound) CallOneway(arg0 context.Context, arg1 *transport.Request) (transport.Ack, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CallOneway", arg0, arg1)
	ret0, _ := ret[0].(transport.Ack)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CallOneway indicates an expected call of CallOneway
func (mr *MockOnewayOutboundMockRecorder) CallOneway(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallOneway", reflect.TypeOf((*MockOnewayOutbound)(nil).CallOneway), arg0, arg1)
}

// IsRunning mocks base method
func (m *MockOnewayOutbound) IsRunning() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsRunning")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsRunning indicates an expected call of IsRunning
func (mr *MockOnewayOutboundMockRecorder) IsRunning() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsRunning", reflect.TypeOf((*MockOnewayOutbound)(nil).IsRunning))
}

// Start mocks base method
func (m *MockOnewayOutbound) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockOnewayOutboundMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockOnewayOutbound)(nil).Start))
}

// Stop mocks base method
func (m *MockOnewayOutbound) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockOnewayOutboundMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockOnewayOutbound)(nil).Stop))
}

// Transports mocks base method
func (m *MockOnewayOutbound) Transports() []transport.Transport {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transports")
	ret0, _ := ret[0].([]transport.Transport)
	return ret0
}

// Transports indicates an expected call of Transports
func (mr *MockOnewayOutboundMockRecorder) Transports() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transports", reflect.TypeOf((*MockOnewayOutbound)(nil).Transports))
}

// MockStreamOutbound is a mock of StreamOutbound interface
type MockStreamOutbound struct {
	ctrl     *gomock.Controller
	recorder *MockStreamOutboundMockRecorder
}

// MockStreamOutboundMockRecorder is the mock recorder for MockStreamOutbound
type MockStreamOutboundMockRecorder struct {
	mock *MockStreamOutbound
}

// NewMockStreamOutbound creates a new mock instance
func NewMockStreamOutbound(ctrl *gomock.Controller) *MockStreamOutbound {
	mock := &MockStreamOutbound{ctrl: ctrl}
	mock.recorder = &MockStreamOutboundMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStreamOutbound) EXPECT() *MockStreamOutboundMockRecorder {
	return m.recorder
}

// CallStream mocks base method
func (m *MockStreamOutbound) CallStream(arg0 context.Context, arg1 *transport.StreamRequest) (*transport.ClientStream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CallStream", arg0, arg1)
	ret0, _ := ret[0].(*transport.ClientStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CallStream indicates an expected call of CallStream
func (mr *MockStreamOutboundMockRecorder) CallStream(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallStream", reflect.TypeOf((*MockStreamOutbound)(nil).CallStream), arg0, arg1)
}

// IsRunning mocks base method
func (m *MockStreamOutbound) IsRunning() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsRunning")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsRunning indicates an expected call of IsRunning
func (mr *MockStreamOutboundMockRecorder) IsRunning() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsRunning", reflect.TypeOf((*MockStreamOutbound)(nil).IsRunning))
}

// Start mocks base method
func (m *MockStreamOutbound) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockStreamOutboundMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockStreamOutbound)(nil).Start))
}

// Stop mocks base method
func (m *MockStreamOutbound) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockStreamOutboundMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockStreamOutbound)(nil).Stop))
}

// Transports mocks base method
func (m *MockStreamOutbound) Transports() []transport.Transport {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transports")
	ret0, _ := ret[0].([]transport.Transport)
	return ret0
}

// Transports indicates an expected call of Transports
func (mr *MockStreamOutboundMockRecorder) Transports() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transports", reflect.TypeOf((*MockStreamOutbound)(nil).Transports))
}
