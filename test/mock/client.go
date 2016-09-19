// Automatically generated by MockGen. DO NOT EDIT!
// Source: ../../client.go

package mock

import (
	gomock "github.com/golang/mock/gomock"
	hrpc "github.com/tsuna/gohbase/hrpc"
	context "golang.org/x/net/context"
)

// Mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *_MockClientRecorder
}

// Recorder for MockClient (not exported)
type _MockClientRecorder struct {
	mock *MockClient
}

func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &_MockClientRecorder{mock}
	return mock
}

func (_m *MockClient) EXPECT() *_MockClientRecorder {
	return _m.recorder
}

func (_m *MockClient) CheckTable(ctx context.Context, table string) error {
	ret := _m.ctrl.Call(_m, "CheckTable", ctx, table)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientRecorder) CheckTable(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CheckTable", arg0, arg1)
}

func (_m *MockClient) Scan(s *hrpc.Scan) ([]*hrpc.Result, error) {
	ret := _m.ctrl.Call(_m, "Scan", s)
	ret0, _ := ret[0].([]*hrpc.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientRecorder) Scan(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Scan", arg0)
}

func (_m *MockClient) Get(g *hrpc.Get) (*hrpc.Result, error) {
	ret := _m.ctrl.Call(_m, "Get", g)
	ret0, _ := ret[0].(*hrpc.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientRecorder) Get(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get", arg0)
}

func (_m *MockClient) Put(p *hrpc.Mutate) (*hrpc.Result, error) {
	ret := _m.ctrl.Call(_m, "Put", p)
	ret0, _ := ret[0].(*hrpc.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientRecorder) Put(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Put", arg0)
}

func (_m *MockClient) Delete(d *hrpc.Mutate) (*hrpc.Result, error) {
	ret := _m.ctrl.Call(_m, "Delete", d)
	ret0, _ := ret[0].(*hrpc.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientRecorder) Delete(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Delete", arg0)
}

func (_m *MockClient) Append(a *hrpc.Mutate) (*hrpc.Result, error) {
	ret := _m.ctrl.Call(_m, "Append", a)
	ret0, _ := ret[0].(*hrpc.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientRecorder) Append(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Append", arg0)
}

func (_m *MockClient) Increment(i *hrpc.Mutate) (int64, error) {
	ret := _m.ctrl.Call(_m, "Increment", i)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientRecorder) Increment(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Increment", arg0)
}

func (_m *MockClient) CheckAndPut(p *hrpc.Mutate, family string, qualifier string, expectedValue []byte) (bool, error) {
	ret := _m.ctrl.Call(_m, "CheckAndPut", p, family, qualifier, expectedValue)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientRecorder) CheckAndPut(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CheckAndPut", arg0, arg1, arg2, arg3)
}

func (_m *MockClient) Close() {
	_m.ctrl.Call(_m, "Close")
}

func (_mr *_MockClientRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

// Mock of AdminClient interface
type MockAdminClient struct {
	ctrl     *gomock.Controller
	recorder *_MockAdminClientRecorder
}

// Recorder for MockAdminClient (not exported)
type _MockAdminClientRecorder struct {
	mock *MockAdminClient
}

func NewMockAdminClient(ctrl *gomock.Controller) *MockAdminClient {
	mock := &MockAdminClient{ctrl: ctrl}
	mock.recorder = &_MockAdminClientRecorder{mock}
	return mock
}

func (_m *MockAdminClient) EXPECT() *_MockAdminClientRecorder {
	return _m.recorder
}

func (_m *MockAdminClient) CreateTable(t *hrpc.CreateTable) error {
	ret := _m.ctrl.Call(_m, "CreateTable", t)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAdminClientRecorder) CreateTable(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateTable", arg0)
}

func (_m *MockAdminClient) DeleteTable(t *hrpc.DeleteTable) error {
	ret := _m.ctrl.Call(_m, "DeleteTable", t)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAdminClientRecorder) DeleteTable(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteTable", arg0)
}

func (_m *MockAdminClient) EnableTable(t *hrpc.EnableTable) error {
	ret := _m.ctrl.Call(_m, "EnableTable", t)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAdminClientRecorder) EnableTable(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "EnableTable", arg0)
}

func (_m *MockAdminClient) DisableTable(t *hrpc.DisableTable) error {
	ret := _m.ctrl.Call(_m, "DisableTable", t)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAdminClientRecorder) DisableTable(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DisableTable", arg0)
}
