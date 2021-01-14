// Code generated by MockGen. DO NOT EDIT.
// Source: ./client.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	compute "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	network "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-12-01/network"
	resources "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-05-01/resources"
	subscriptions "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-06-01/subscriptions"
	gomock "github.com/golang/mock/gomock"
)

// MockAPI is a mock of API interface.
type MockAPI struct {
	ctrl     *gomock.Controller
	recorder *MockAPIMockRecorder
}

// MockAPIMockRecorder is the mock recorder for MockAPI.
type MockAPIMockRecorder struct {
	mock *MockAPI
}

// NewMockAPI creates a new mock instance.
func NewMockAPI(ctrl *gomock.Controller) *MockAPI {
	mock := &MockAPI{ctrl: ctrl}
	mock.recorder = &MockAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAPI) EXPECT() *MockAPIMockRecorder {
	return m.recorder
}

// GetComputeSubnet mocks base method.
func (m *MockAPI) GetComputeSubnet(ctx context.Context, resourceGroupName, virtualNetwork, subnet string) (*network.Subnet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComputeSubnet", ctx, resourceGroupName, virtualNetwork, subnet)
	ret0, _ := ret[0].(*network.Subnet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetComputeSubnet indicates an expected call of GetComputeSubnet.
func (mr *MockAPIMockRecorder) GetComputeSubnet(ctx, resourceGroupName, virtualNetwork, subnet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComputeSubnet", reflect.TypeOf((*MockAPI)(nil).GetComputeSubnet), ctx, resourceGroupName, virtualNetwork, subnet)
}

// GetControlPlaneSubnet mocks base method.
func (m *MockAPI) GetControlPlaneSubnet(ctx context.Context, resourceGroupName, virtualNetwork, subnet string) (*network.Subnet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetControlPlaneSubnet", ctx, resourceGroupName, virtualNetwork, subnet)
	ret0, _ := ret[0].(*network.Subnet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetControlPlaneSubnet indicates an expected call of GetControlPlaneSubnet.
func (mr *MockAPIMockRecorder) GetControlPlaneSubnet(ctx, resourceGroupName, virtualNetwork, subnet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetControlPlaneSubnet", reflect.TypeOf((*MockAPI)(nil).GetControlPlaneSubnet), ctx, resourceGroupName, virtualNetwork, subnet)
}

// GetDiskSkus mocks base method.
func (m *MockAPI) GetDiskSkus(ctx context.Context, region string) ([]compute.ResourceSku, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDiskSkus", ctx, region)
	ret0, _ := ret[0].([]compute.ResourceSku)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDiskSkus indicates an expected call of GetDiskSkus.
func (mr *MockAPIMockRecorder) GetDiskSkus(ctx, region interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDiskSkus", reflect.TypeOf((*MockAPI)(nil).GetDiskSkus), ctx, region)
}

// GetGroup mocks base method.
func (m *MockAPI) GetGroup(ctx context.Context, groupName string) (*resources.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroup", ctx, groupName)
	ret0, _ := ret[0].(*resources.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroup indicates an expected call of GetGroup.
func (mr *MockAPIMockRecorder) GetGroup(ctx, groupName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroup", reflect.TypeOf((*MockAPI)(nil).GetGroup), ctx, groupName)
}

// GetResourcesProvider mocks base method.
func (m *MockAPI) GetResourcesProvider(ctx context.Context, resourceProviderNamespace string) (*resources.Provider, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResourcesProvider", ctx, resourceProviderNamespace)
	ret0, _ := ret[0].(*resources.Provider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResourcesProvider indicates an expected call of GetResourcesProvider.
func (mr *MockAPIMockRecorder) GetResourcesProvider(ctx, resourceProviderNamespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourcesProvider", reflect.TypeOf((*MockAPI)(nil).GetResourcesProvider), ctx, resourceProviderNamespace)
}

// GetVirtualMachineSku mocks base method.
func (m *MockAPI) GetVirtualMachineSku(ctx context.Context, name, region string) (*compute.ResourceSku, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVirtualMachineSku", ctx, name, region)
	ret0, _ := ret[0].(*compute.ResourceSku)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVirtualMachineSku indicates an expected call of GetVirtualMachineSku.
func (mr *MockAPIMockRecorder) GetVirtualMachineSku(ctx, name, region interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVirtualMachineSku", reflect.TypeOf((*MockAPI)(nil).GetVirtualMachineSku), ctx, name, region)
}

// GetVirtualNetwork mocks base method.
func (m *MockAPI) GetVirtualNetwork(ctx context.Context, resourceGroupName, virtualNetwork string) (*network.VirtualNetwork, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVirtualNetwork", ctx, resourceGroupName, virtualNetwork)
	ret0, _ := ret[0].(*network.VirtualNetwork)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVirtualNetwork indicates an expected call of GetVirtualNetwork.
func (mr *MockAPIMockRecorder) GetVirtualNetwork(ctx, resourceGroupName, virtualNetwork interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVirtualNetwork", reflect.TypeOf((*MockAPI)(nil).GetVirtualNetwork), ctx, resourceGroupName, virtualNetwork)
}

// ListLocations mocks base method.
func (m *MockAPI) ListLocations(ctx context.Context) (*[]subscriptions.Location, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListLocations", ctx)
	ret0, _ := ret[0].(*[]subscriptions.Location)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListLocations indicates an expected call of ListLocations.
func (mr *MockAPIMockRecorder) ListLocations(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLocations", reflect.TypeOf((*MockAPI)(nil).ListLocations), ctx)
}

// ListResourceIDsByGroup mocks base method.
func (m *MockAPI) ListResourceIDsByGroup(ctx context.Context, groupName string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListResourceIDsByGroup", ctx, groupName)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListResourceIDsByGroup indicates an expected call of ListResourceIDsByGroup.
func (mr *MockAPIMockRecorder) ListResourceIDsByGroup(ctx, groupName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListResourceIDsByGroup", reflect.TypeOf((*MockAPI)(nil).ListResourceIDsByGroup), ctx, groupName)
}
