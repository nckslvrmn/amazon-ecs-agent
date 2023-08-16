// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/amazon-ecs-agent/agent/engine/dockerstate (interfaces: TaskEngineState)

// Package mock_dockerstate is a generated GoMock package.
package mock_dockerstate

import (
	reflect "reflect"

	container "github.com/aws/amazon-ecs-agent/agent/api/container"
	task "github.com/aws/amazon-ecs-agent/agent/api/task"
	image "github.com/aws/amazon-ecs-agent/agent/engine/image"
	resource "github.com/aws/amazon-ecs-agent/ecs-agent/api/resource"
	networkinterface "github.com/aws/amazon-ecs-agent/ecs-agent/netlib/model/networkinterface"
	gomock "github.com/golang/mock/gomock"
)

// MockTaskEngineState is a mock of TaskEngineState interface.
type MockTaskEngineState struct {
	ctrl     *gomock.Controller
	recorder *MockTaskEngineStateMockRecorder
}

// MockTaskEngineStateMockRecorder is the mock recorder for MockTaskEngineState.
type MockTaskEngineStateMockRecorder struct {
	mock *MockTaskEngineState
}

// NewMockTaskEngineState creates a new mock instance.
func NewMockTaskEngineState(ctrl *gomock.Controller) *MockTaskEngineState {
	mock := &MockTaskEngineState{ctrl: ctrl}
	mock.recorder = &MockTaskEngineStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskEngineState) EXPECT() *MockTaskEngineStateMockRecorder {
	return m.recorder
}

// AddContainer mocks base method.
func (m *MockTaskEngineState) AddContainer(arg0 *container.DockerContainer, arg1 *task.Task) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddContainer", arg0, arg1)
}

// AddContainer indicates an expected call of AddContainer.
func (mr *MockTaskEngineStateMockRecorder) AddContainer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddContainer", reflect.TypeOf((*MockTaskEngineState)(nil).AddContainer), arg0, arg1)
}

// AddEBSAttachment mocks base method.
func (m *MockTaskEngineState) AddEBSAttachment(arg0 *resource.ResourceAttachment) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddEBSAttachment", arg0)
}

// AddEBSAttachment indicates an expected call of AddEBSAttachment.
func (mr *MockTaskEngineStateMockRecorder) AddEBSAttachment(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEBSAttachment", reflect.TypeOf((*MockTaskEngineState)(nil).AddEBSAttachment), arg0)
}

// AddENIAttachment mocks base method.
func (m *MockTaskEngineState) AddENIAttachment(arg0 *networkinterface.ENIAttachment) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddENIAttachment", arg0)
}

// AddENIAttachment indicates an expected call of AddENIAttachment.
func (mr *MockTaskEngineStateMockRecorder) AddENIAttachment(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddENIAttachment", reflect.TypeOf((*MockTaskEngineState)(nil).AddENIAttachment), arg0)
}

// AddImageState mocks base method.
func (m *MockTaskEngineState) AddImageState(arg0 *image.ImageState) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddImageState", arg0)
}

// AddImageState indicates an expected call of AddImageState.
func (mr *MockTaskEngineStateMockRecorder) AddImageState(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddImageState", reflect.TypeOf((*MockTaskEngineState)(nil).AddImageState), arg0)
}

// AddPulledContainer mocks base method.
func (m *MockTaskEngineState) AddPulledContainer(arg0 *container.DockerContainer, arg1 *task.Task) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddPulledContainer", arg0, arg1)
}

// AddPulledContainer indicates an expected call of AddPulledContainer.
func (mr *MockTaskEngineStateMockRecorder) AddPulledContainer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPulledContainer", reflect.TypeOf((*MockTaskEngineState)(nil).AddPulledContainer), arg0, arg1)
}

// AddTask mocks base method.
func (m *MockTaskEngineState) AddTask(arg0 *task.Task) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddTask", arg0)
}

// AddTask indicates an expected call of AddTask.
func (mr *MockTaskEngineStateMockRecorder) AddTask(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTask", reflect.TypeOf((*MockTaskEngineState)(nil).AddTask), arg0)
}

// AddTaskIPAddress mocks base method.
func (m *MockTaskEngineState) AddTaskIPAddress(arg0, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddTaskIPAddress", arg0, arg1)
}

// AddTaskIPAddress indicates an expected call of AddTaskIPAddress.
func (mr *MockTaskEngineStateMockRecorder) AddTaskIPAddress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTaskIPAddress", reflect.TypeOf((*MockTaskEngineState)(nil).AddTaskIPAddress), arg0, arg1)
}

// AllENIAttachments mocks base method.
func (m *MockTaskEngineState) AllENIAttachments() []*networkinterface.ENIAttachment {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllENIAttachments")
	ret0, _ := ret[0].([]*networkinterface.ENIAttachment)
	return ret0
}

// AllENIAttachments indicates an expected call of AllENIAttachments.
func (mr *MockTaskEngineStateMockRecorder) AllENIAttachments() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllENIAttachments", reflect.TypeOf((*MockTaskEngineState)(nil).AllENIAttachments))
}

// AllExternalTasks mocks base method.
func (m *MockTaskEngineState) AllExternalTasks() []*task.Task {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllExternalTasks")
	ret0, _ := ret[0].([]*task.Task)
	return ret0
}

// AllExternalTasks indicates an expected call of AllExternalTasks.
func (mr *MockTaskEngineStateMockRecorder) AllExternalTasks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllExternalTasks", reflect.TypeOf((*MockTaskEngineState)(nil).AllExternalTasks))
}

// AllImageStates mocks base method.
func (m *MockTaskEngineState) AllImageStates() []*image.ImageState {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllImageStates")
	ret0, _ := ret[0].([]*image.ImageState)
	return ret0
}

// AllImageStates indicates an expected call of AllImageStates.
func (mr *MockTaskEngineStateMockRecorder) AllImageStates() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllImageStates", reflect.TypeOf((*MockTaskEngineState)(nil).AllImageStates))
}

// AllTasks mocks base method.
func (m *MockTaskEngineState) AllTasks() []*task.Task {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllTasks")
	ret0, _ := ret[0].([]*task.Task)
	return ret0
}

// AllTasks indicates an expected call of AllTasks.
func (mr *MockTaskEngineStateMockRecorder) AllTasks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllTasks", reflect.TypeOf((*MockTaskEngineState)(nil).AllTasks))
}

// ContainerByID mocks base method.
func (m *MockTaskEngineState) ContainerByID(arg0 string) (*container.DockerContainer, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerByID", arg0)
	ret0, _ := ret[0].(*container.DockerContainer)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// ContainerByID indicates an expected call of ContainerByID.
func (mr *MockTaskEngineStateMockRecorder) ContainerByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerByID", reflect.TypeOf((*MockTaskEngineState)(nil).ContainerByID), arg0)
}

// ContainerMapByArn mocks base method.
func (m *MockTaskEngineState) ContainerMapByArn(arg0 string) (map[string]*container.DockerContainer, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerMapByArn", arg0)
	ret0, _ := ret[0].(map[string]*container.DockerContainer)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// ContainerMapByArn indicates an expected call of ContainerMapByArn.
func (mr *MockTaskEngineStateMockRecorder) ContainerMapByArn(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerMapByArn", reflect.TypeOf((*MockTaskEngineState)(nil).ContainerMapByArn), arg0)
}

// DockerIDByV3EndpointID mocks base method.
func (m *MockTaskEngineState) DockerIDByV3EndpointID(arg0 string) (string, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DockerIDByV3EndpointID", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// DockerIDByV3EndpointID indicates an expected call of DockerIDByV3EndpointID.
func (mr *MockTaskEngineStateMockRecorder) DockerIDByV3EndpointID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DockerIDByV3EndpointID", reflect.TypeOf((*MockTaskEngineState)(nil).DockerIDByV3EndpointID), arg0)
}

// ENIByMac mocks base method.
func (m *MockTaskEngineState) ENIByMac(arg0 string) (*networkinterface.ENIAttachment, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ENIByMac", arg0)
	ret0, _ := ret[0].(*networkinterface.ENIAttachment)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// ENIByMac indicates an expected call of ENIByMac.
func (mr *MockTaskEngineStateMockRecorder) ENIByMac(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ENIByMac", reflect.TypeOf((*MockTaskEngineState)(nil).ENIByMac), arg0)
}

// GetAllContainerIDs mocks base method.
func (m *MockTaskEngineState) GetAllContainerIDs() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllContainerIDs")
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetAllContainerIDs indicates an expected call of GetAllContainerIDs.
func (mr *MockTaskEngineStateMockRecorder) GetAllContainerIDs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllContainerIDs", reflect.TypeOf((*MockTaskEngineState)(nil).GetAllContainerIDs))
}

// GetAllEBSAttachments mocks base method.
func (m *MockTaskEngineState) GetAllEBSAttachments() []*resource.ResourceAttachment {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllEBSAttachments")
	ret0, _ := ret[0].([]*resource.ResourceAttachment)
	return ret0
}

// GetAllEBSAttachments indicates an expected call of GetAllEBSAttachments.
func (mr *MockTaskEngineStateMockRecorder) GetAllEBSAttachments() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllEBSAttachments", reflect.TypeOf((*MockTaskEngineState)(nil).GetAllEBSAttachments))
}

// GetAllPendingEBSAttachments mocks base method.
func (m *MockTaskEngineState) GetAllPendingEBSAttachments() []*resource.ResourceAttachment {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllPendingEBSAttachments")
	ret0, _ := ret[0].([]*resource.ResourceAttachment)
	return ret0
}

// GetAllPendingEBSAttachments indicates an expected call of GetAllPendingEBSAttachments.
func (mr *MockTaskEngineStateMockRecorder) GetAllPendingEBSAttachments() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllPendingEBSAttachments", reflect.TypeOf((*MockTaskEngineState)(nil).GetAllPendingEBSAttachments))
}

// GetEBSByVolumeId mocks base method.
func (m *MockTaskEngineState) GetEBSByVolumeId(arg0 string) (*resource.ResourceAttachment, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEBSByVolumeId", arg0)
	ret0, _ := ret[0].(*resource.ResourceAttachment)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetEBSByVolumeId indicates an expected call of GetEBSByVolumeId.
func (mr *MockTaskEngineStateMockRecorder) GetEBSByVolumeId(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEBSByVolumeId", reflect.TypeOf((*MockTaskEngineState)(nil).GetEBSByVolumeId), arg0)
}

// GetIPAddressByTaskARN mocks base method.
func (m *MockTaskEngineState) GetIPAddressByTaskARN(arg0 string) (string, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIPAddressByTaskARN", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetIPAddressByTaskARN indicates an expected call of GetIPAddressByTaskARN.
func (mr *MockTaskEngineStateMockRecorder) GetIPAddressByTaskARN(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIPAddressByTaskARN", reflect.TypeOf((*MockTaskEngineState)(nil).GetIPAddressByTaskARN), arg0)
}

// GetTaskByIPAddress mocks base method.
func (m *MockTaskEngineState) GetTaskByIPAddress(arg0 string) (string, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTaskByIPAddress", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetTaskByIPAddress indicates an expected call of GetTaskByIPAddress.
func (mr *MockTaskEngineStateMockRecorder) GetTaskByIPAddress(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskByIPAddress", reflect.TypeOf((*MockTaskEngineState)(nil).GetTaskByIPAddress), arg0)
}

// MarshalJSON mocks base method.
func (m *MockTaskEngineState) MarshalJSON() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarshalJSON")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalJSON indicates an expected call of MarshalJSON.
func (mr *MockTaskEngineStateMockRecorder) MarshalJSON() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarshalJSON", reflect.TypeOf((*MockTaskEngineState)(nil).MarshalJSON))
}

// PulledContainerMapByArn mocks base method.
func (m *MockTaskEngineState) PulledContainerMapByArn(arg0 string) (map[string]*container.DockerContainer, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PulledContainerMapByArn", arg0)
	ret0, _ := ret[0].(map[string]*container.DockerContainer)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// PulledContainerMapByArn indicates an expected call of PulledContainerMapByArn.
func (mr *MockTaskEngineStateMockRecorder) PulledContainerMapByArn(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PulledContainerMapByArn", reflect.TypeOf((*MockTaskEngineState)(nil).PulledContainerMapByArn), arg0)
}

// RemoveEBSAttachment mocks base method.
func (m *MockTaskEngineState) RemoveEBSAttachment(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveEBSAttachment", arg0)
}

// RemoveEBSAttachment indicates an expected call of RemoveEBSAttachment.
func (mr *MockTaskEngineStateMockRecorder) RemoveEBSAttachment(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveEBSAttachment", reflect.TypeOf((*MockTaskEngineState)(nil).RemoveEBSAttachment), arg0)
}

// RemoveENIAttachment mocks base method.
func (m *MockTaskEngineState) RemoveENIAttachment(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveENIAttachment", arg0)
}

// RemoveENIAttachment indicates an expected call of RemoveENIAttachment.
func (mr *MockTaskEngineStateMockRecorder) RemoveENIAttachment(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveENIAttachment", reflect.TypeOf((*MockTaskEngineState)(nil).RemoveENIAttachment), arg0)
}

// RemoveImageState mocks base method.
func (m *MockTaskEngineState) RemoveImageState(arg0 *image.ImageState) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveImageState", arg0)
}

// RemoveImageState indicates an expected call of RemoveImageState.
func (mr *MockTaskEngineStateMockRecorder) RemoveImageState(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveImageState", reflect.TypeOf((*MockTaskEngineState)(nil).RemoveImageState), arg0)
}

// RemoveTask mocks base method.
func (m *MockTaskEngineState) RemoveTask(arg0 *task.Task) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveTask", arg0)
}

// RemoveTask indicates an expected call of RemoveTask.
func (mr *MockTaskEngineStateMockRecorder) RemoveTask(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTask", reflect.TypeOf((*MockTaskEngineState)(nil).RemoveTask), arg0)
}

// Reset mocks base method.
func (m *MockTaskEngineState) Reset() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Reset")
}

// Reset indicates an expected call of Reset.
func (mr *MockTaskEngineStateMockRecorder) Reset() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockTaskEngineState)(nil).Reset))
}

// TaskARNByV3EndpointID mocks base method.
func (m *MockTaskEngineState) TaskARNByV3EndpointID(arg0 string) (string, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TaskARNByV3EndpointID", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// TaskARNByV3EndpointID indicates an expected call of TaskARNByV3EndpointID.
func (mr *MockTaskEngineStateMockRecorder) TaskARNByV3EndpointID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TaskARNByV3EndpointID", reflect.TypeOf((*MockTaskEngineState)(nil).TaskARNByV3EndpointID), arg0)
}

// TaskByArn mocks base method.
func (m *MockTaskEngineState) TaskByArn(arg0 string) (*task.Task, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TaskByArn", arg0)
	ret0, _ := ret[0].(*task.Task)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// TaskByArn indicates an expected call of TaskByArn.
func (mr *MockTaskEngineStateMockRecorder) TaskByArn(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TaskByArn", reflect.TypeOf((*MockTaskEngineState)(nil).TaskByArn), arg0)
}

// TaskByID mocks base method.
func (m *MockTaskEngineState) TaskByID(arg0 string) (*task.Task, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TaskByID", arg0)
	ret0, _ := ret[0].(*task.Task)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// TaskByID indicates an expected call of TaskByID.
func (mr *MockTaskEngineStateMockRecorder) TaskByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TaskByID", reflect.TypeOf((*MockTaskEngineState)(nil).TaskByID), arg0)
}

// TaskByShortID mocks base method.
func (m *MockTaskEngineState) TaskByShortID(arg0 string) ([]*task.Task, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TaskByShortID", arg0)
	ret0, _ := ret[0].([]*task.Task)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// TaskByShortID indicates an expected call of TaskByShortID.
func (mr *MockTaskEngineStateMockRecorder) TaskByShortID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TaskByShortID", reflect.TypeOf((*MockTaskEngineState)(nil).TaskByShortID), arg0)
}

// UnmarshalJSON mocks base method.
func (m *MockTaskEngineState) UnmarshalJSON(arg0 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnmarshalJSON", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnmarshalJSON indicates an expected call of UnmarshalJSON.
func (mr *MockTaskEngineStateMockRecorder) UnmarshalJSON(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnmarshalJSON", reflect.TypeOf((*MockTaskEngineState)(nil).UnmarshalJSON), arg0)
}
