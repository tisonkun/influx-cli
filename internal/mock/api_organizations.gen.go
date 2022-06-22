// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/influxdata/influx-cli/v2/api (interfaces: OrganizationsApi)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	api "github.com/influxdata/influx-cli/v2/api"
)

// MockOrganizationsApi is a mock of OrganizationsApi interface.
type MockOrganizationsApi struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationsApiMockRecorder
}

// MockOrganizationsApiMockRecorder is the mock recorder for MockOrganizationsApi.
type MockOrganizationsApiMockRecorder struct {
	mock *MockOrganizationsApi
}

// NewMockOrganizationsApi creates a new mock instance.
func NewMockOrganizationsApi(ctrl *gomock.Controller) *MockOrganizationsApi {
	mock := &MockOrganizationsApi{ctrl: ctrl}
	mock.recorder = &MockOrganizationsApiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationsApi) EXPECT() *MockOrganizationsApiMockRecorder {
	return m.recorder
}

// DeleteOrgsID mocks base method.
func (m *MockOrganizationsApi) DeleteOrgsID(arg0 context.Context, arg1 string) api.ApiDeleteOrgsIDRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrgsID", arg0, arg1)
	ret0, _ := ret[0].(api.ApiDeleteOrgsIDRequest)
	return ret0
}

// DeleteOrgsID indicates an expected call of DeleteOrgsID.
func (mr *MockOrganizationsApiMockRecorder) DeleteOrgsID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrgsID", reflect.TypeOf((*MockOrganizationsApi)(nil).DeleteOrgsID), arg0, arg1)
}

// DeleteOrgsIDExecute mocks base method.
func (m *MockOrganizationsApi) DeleteOrgsIDExecute(arg0 api.ApiDeleteOrgsIDRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrgsIDExecute", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOrgsIDExecute indicates an expected call of DeleteOrgsIDExecute.
func (mr *MockOrganizationsApiMockRecorder) DeleteOrgsIDExecute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrgsIDExecute", reflect.TypeOf((*MockOrganizationsApi)(nil).DeleteOrgsIDExecute), arg0)
}

// DeleteOrgsIDExecuteWithHttpInfo mocks base method.
func (m *MockOrganizationsApi) DeleteOrgsIDExecuteWithHttpInfo(arg0 api.ApiDeleteOrgsIDRequest) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrgsIDExecuteWithHttpInfo", arg0)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteOrgsIDExecuteWithHttpInfo indicates an expected call of DeleteOrgsIDExecuteWithHttpInfo.
func (mr *MockOrganizationsApiMockRecorder) DeleteOrgsIDExecuteWithHttpInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrgsIDExecuteWithHttpInfo", reflect.TypeOf((*MockOrganizationsApi)(nil).DeleteOrgsIDExecuteWithHttpInfo), arg0)
}

// DeleteOrgsIDMembersID mocks base method.
func (m *MockOrganizationsApi) DeleteOrgsIDMembersID(arg0 context.Context, arg1, arg2 string) api.ApiDeleteOrgsIDMembersIDRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrgsIDMembersID", arg0, arg1, arg2)
	ret0, _ := ret[0].(api.ApiDeleteOrgsIDMembersIDRequest)
	return ret0
}

// DeleteOrgsIDMembersID indicates an expected call of DeleteOrgsIDMembersID.
func (mr *MockOrganizationsApiMockRecorder) DeleteOrgsIDMembersID(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrgsIDMembersID", reflect.TypeOf((*MockOrganizationsApi)(nil).DeleteOrgsIDMembersID), arg0, arg1, arg2)
}

// DeleteOrgsIDMembersIDExecute mocks base method.
func (m *MockOrganizationsApi) DeleteOrgsIDMembersIDExecute(arg0 api.ApiDeleteOrgsIDMembersIDRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrgsIDMembersIDExecute", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOrgsIDMembersIDExecute indicates an expected call of DeleteOrgsIDMembersIDExecute.
func (mr *MockOrganizationsApiMockRecorder) DeleteOrgsIDMembersIDExecute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrgsIDMembersIDExecute", reflect.TypeOf((*MockOrganizationsApi)(nil).DeleteOrgsIDMembersIDExecute), arg0)
}

// DeleteOrgsIDMembersIDExecuteWithHttpInfo mocks base method.
func (m *MockOrganizationsApi) DeleteOrgsIDMembersIDExecuteWithHttpInfo(arg0 api.ApiDeleteOrgsIDMembersIDRequest) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrgsIDMembersIDExecuteWithHttpInfo", arg0)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteOrgsIDMembersIDExecuteWithHttpInfo indicates an expected call of DeleteOrgsIDMembersIDExecuteWithHttpInfo.
func (mr *MockOrganizationsApiMockRecorder) DeleteOrgsIDMembersIDExecuteWithHttpInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrgsIDMembersIDExecuteWithHttpInfo", reflect.TypeOf((*MockOrganizationsApi)(nil).DeleteOrgsIDMembersIDExecuteWithHttpInfo), arg0)
}

// GetOrgs mocks base method.
func (m *MockOrganizationsApi) GetOrgs(arg0 context.Context) api.ApiGetOrgsRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrgs", arg0)
	ret0, _ := ret[0].(api.ApiGetOrgsRequest)
	return ret0
}

// GetOrgs indicates an expected call of GetOrgs.
func (mr *MockOrganizationsApiMockRecorder) GetOrgs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrgs", reflect.TypeOf((*MockOrganizationsApi)(nil).GetOrgs), arg0)
}

// GetOrgsExecute mocks base method.
func (m *MockOrganizationsApi) GetOrgsExecute(arg0 api.ApiGetOrgsRequest) (api.Organizations, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrgsExecute", arg0)
	ret0, _ := ret[0].(api.Organizations)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrgsExecute indicates an expected call of GetOrgsExecute.
func (mr *MockOrganizationsApiMockRecorder) GetOrgsExecute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrgsExecute", reflect.TypeOf((*MockOrganizationsApi)(nil).GetOrgsExecute), arg0)
}

// GetOrgsExecuteWithHttpInfo mocks base method.
func (m *MockOrganizationsApi) GetOrgsExecuteWithHttpInfo(arg0 api.ApiGetOrgsRequest) (api.Organizations, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrgsExecuteWithHttpInfo", arg0)
	ret0, _ := ret[0].(api.Organizations)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetOrgsExecuteWithHttpInfo indicates an expected call of GetOrgsExecuteWithHttpInfo.
func (mr *MockOrganizationsApiMockRecorder) GetOrgsExecuteWithHttpInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrgsExecuteWithHttpInfo", reflect.TypeOf((*MockOrganizationsApi)(nil).GetOrgsExecuteWithHttpInfo), arg0)
}

// GetOrgsID mocks base method.
func (m *MockOrganizationsApi) GetOrgsID(arg0 context.Context, arg1 string) api.ApiGetOrgsIDRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrgsID", arg0, arg1)
	ret0, _ := ret[0].(api.ApiGetOrgsIDRequest)
	return ret0
}

// GetOrgsID indicates an expected call of GetOrgsID.
func (mr *MockOrganizationsApiMockRecorder) GetOrgsID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrgsID", reflect.TypeOf((*MockOrganizationsApi)(nil).GetOrgsID), arg0, arg1)
}

// GetOrgsIDExecute mocks base method.
func (m *MockOrganizationsApi) GetOrgsIDExecute(arg0 api.ApiGetOrgsIDRequest) (api.Organization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrgsIDExecute", arg0)
	ret0, _ := ret[0].(api.Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrgsIDExecute indicates an expected call of GetOrgsIDExecute.
func (mr *MockOrganizationsApiMockRecorder) GetOrgsIDExecute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrgsIDExecute", reflect.TypeOf((*MockOrganizationsApi)(nil).GetOrgsIDExecute), arg0)
}

// GetOrgsIDExecuteWithHttpInfo mocks base method.
func (m *MockOrganizationsApi) GetOrgsIDExecuteWithHttpInfo(arg0 api.ApiGetOrgsIDRequest) (api.Organization, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrgsIDExecuteWithHttpInfo", arg0)
	ret0, _ := ret[0].(api.Organization)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetOrgsIDExecuteWithHttpInfo indicates an expected call of GetOrgsIDExecuteWithHttpInfo.
func (mr *MockOrganizationsApiMockRecorder) GetOrgsIDExecuteWithHttpInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrgsIDExecuteWithHttpInfo", reflect.TypeOf((*MockOrganizationsApi)(nil).GetOrgsIDExecuteWithHttpInfo), arg0)
}

// GetOrgsIDMembers mocks base method.
func (m *MockOrganizationsApi) GetOrgsIDMembers(arg0 context.Context, arg1 string) api.ApiGetOrgsIDMembersRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrgsIDMembers", arg0, arg1)
	ret0, _ := ret[0].(api.ApiGetOrgsIDMembersRequest)
	return ret0
}

// GetOrgsIDMembers indicates an expected call of GetOrgsIDMembers.
func (mr *MockOrganizationsApiMockRecorder) GetOrgsIDMembers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrgsIDMembers", reflect.TypeOf((*MockOrganizationsApi)(nil).GetOrgsIDMembers), arg0, arg1)
}

// GetOrgsIDMembersExecute mocks base method.
func (m *MockOrganizationsApi) GetOrgsIDMembersExecute(arg0 api.ApiGetOrgsIDMembersRequest) (api.ResourceMembers, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrgsIDMembersExecute", arg0)
	ret0, _ := ret[0].(api.ResourceMembers)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrgsIDMembersExecute indicates an expected call of GetOrgsIDMembersExecute.
func (mr *MockOrganizationsApiMockRecorder) GetOrgsIDMembersExecute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrgsIDMembersExecute", reflect.TypeOf((*MockOrganizationsApi)(nil).GetOrgsIDMembersExecute), arg0)
}

// GetOrgsIDMembersExecuteWithHttpInfo mocks base method.
func (m *MockOrganizationsApi) GetOrgsIDMembersExecuteWithHttpInfo(arg0 api.ApiGetOrgsIDMembersRequest) (api.ResourceMembers, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrgsIDMembersExecuteWithHttpInfo", arg0)
	ret0, _ := ret[0].(api.ResourceMembers)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetOrgsIDMembersExecuteWithHttpInfo indicates an expected call of GetOrgsIDMembersExecuteWithHttpInfo.
func (mr *MockOrganizationsApiMockRecorder) GetOrgsIDMembersExecuteWithHttpInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrgsIDMembersExecuteWithHttpInfo", reflect.TypeOf((*MockOrganizationsApi)(nil).GetOrgsIDMembersExecuteWithHttpInfo), arg0)
}

// PatchOrgsID mocks base method.
func (m *MockOrganizationsApi) PatchOrgsID(arg0 context.Context, arg1 string) api.ApiPatchOrgsIDRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PatchOrgsID", arg0, arg1)
	ret0, _ := ret[0].(api.ApiPatchOrgsIDRequest)
	return ret0
}

// PatchOrgsID indicates an expected call of PatchOrgsID.
func (mr *MockOrganizationsApiMockRecorder) PatchOrgsID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchOrgsID", reflect.TypeOf((*MockOrganizationsApi)(nil).PatchOrgsID), arg0, arg1)
}

// PatchOrgsIDExecute mocks base method.
func (m *MockOrganizationsApi) PatchOrgsIDExecute(arg0 api.ApiPatchOrgsIDRequest) (api.Organization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PatchOrgsIDExecute", arg0)
	ret0, _ := ret[0].(api.Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PatchOrgsIDExecute indicates an expected call of PatchOrgsIDExecute.
func (mr *MockOrganizationsApiMockRecorder) PatchOrgsIDExecute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchOrgsIDExecute", reflect.TypeOf((*MockOrganizationsApi)(nil).PatchOrgsIDExecute), arg0)
}

// PatchOrgsIDExecuteWithHttpInfo mocks base method.
func (m *MockOrganizationsApi) PatchOrgsIDExecuteWithHttpInfo(arg0 api.ApiPatchOrgsIDRequest) (api.Organization, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PatchOrgsIDExecuteWithHttpInfo", arg0)
	ret0, _ := ret[0].(api.Organization)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// PatchOrgsIDExecuteWithHttpInfo indicates an expected call of PatchOrgsIDExecuteWithHttpInfo.
func (mr *MockOrganizationsApiMockRecorder) PatchOrgsIDExecuteWithHttpInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchOrgsIDExecuteWithHttpInfo", reflect.TypeOf((*MockOrganizationsApi)(nil).PatchOrgsIDExecuteWithHttpInfo), arg0)
}

// PostOrgs mocks base method.
func (m *MockOrganizationsApi) PostOrgs(arg0 context.Context) api.ApiPostOrgsRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostOrgs", arg0)
	ret0, _ := ret[0].(api.ApiPostOrgsRequest)
	return ret0
}

// PostOrgs indicates an expected call of PostOrgs.
func (mr *MockOrganizationsApiMockRecorder) PostOrgs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostOrgs", reflect.TypeOf((*MockOrganizationsApi)(nil).PostOrgs), arg0)
}

// PostOrgsExecute mocks base method.
func (m *MockOrganizationsApi) PostOrgsExecute(arg0 api.ApiPostOrgsRequest) (api.Organization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostOrgsExecute", arg0)
	ret0, _ := ret[0].(api.Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostOrgsExecute indicates an expected call of PostOrgsExecute.
func (mr *MockOrganizationsApiMockRecorder) PostOrgsExecute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostOrgsExecute", reflect.TypeOf((*MockOrganizationsApi)(nil).PostOrgsExecute), arg0)
}

// PostOrgsExecuteWithHttpInfo mocks base method.
func (m *MockOrganizationsApi) PostOrgsExecuteWithHttpInfo(arg0 api.ApiPostOrgsRequest) (api.Organization, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostOrgsExecuteWithHttpInfo", arg0)
	ret0, _ := ret[0].(api.Organization)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// PostOrgsExecuteWithHttpInfo indicates an expected call of PostOrgsExecuteWithHttpInfo.
func (mr *MockOrganizationsApiMockRecorder) PostOrgsExecuteWithHttpInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostOrgsExecuteWithHttpInfo", reflect.TypeOf((*MockOrganizationsApi)(nil).PostOrgsExecuteWithHttpInfo), arg0)
}

// PostOrgsIDMembers mocks base method.
func (m *MockOrganizationsApi) PostOrgsIDMembers(arg0 context.Context, arg1 string) api.ApiPostOrgsIDMembersRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostOrgsIDMembers", arg0, arg1)
	ret0, _ := ret[0].(api.ApiPostOrgsIDMembersRequest)
	return ret0
}

// PostOrgsIDMembers indicates an expected call of PostOrgsIDMembers.
func (mr *MockOrganizationsApiMockRecorder) PostOrgsIDMembers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostOrgsIDMembers", reflect.TypeOf((*MockOrganizationsApi)(nil).PostOrgsIDMembers), arg0, arg1)
}

// PostOrgsIDMembersExecute mocks base method.
func (m *MockOrganizationsApi) PostOrgsIDMembersExecute(arg0 api.ApiPostOrgsIDMembersRequest) (api.ResourceMember, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostOrgsIDMembersExecute", arg0)
	ret0, _ := ret[0].(api.ResourceMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostOrgsIDMembersExecute indicates an expected call of PostOrgsIDMembersExecute.
func (mr *MockOrganizationsApiMockRecorder) PostOrgsIDMembersExecute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostOrgsIDMembersExecute", reflect.TypeOf((*MockOrganizationsApi)(nil).PostOrgsIDMembersExecute), arg0)
}

// PostOrgsIDMembersExecuteWithHttpInfo mocks base method.
func (m *MockOrganizationsApi) PostOrgsIDMembersExecuteWithHttpInfo(arg0 api.ApiPostOrgsIDMembersRequest) (api.ResourceMember, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostOrgsIDMembersExecuteWithHttpInfo", arg0)
	ret0, _ := ret[0].(api.ResourceMember)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// PostOrgsIDMembersExecuteWithHttpInfo indicates an expected call of PostOrgsIDMembersExecuteWithHttpInfo.
func (mr *MockOrganizationsApiMockRecorder) PostOrgsIDMembersExecuteWithHttpInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostOrgsIDMembersExecuteWithHttpInfo", reflect.TypeOf((*MockOrganizationsApi)(nil).PostOrgsIDMembersExecuteWithHttpInfo), arg0)
}
