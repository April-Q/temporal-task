// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by protoc-gen-go. DO NOT EDIT.
// plugins:
// 	protoc-gen-go
// 	protoc
// source: temporal/api/enums/v1/update.proto

package enums

import (
	reflect "reflect"
	"strconv"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// UpdateWorkflowExecutionLifecycleStage is specified by clients invoking
// workflow execution updates and used to indicate to the server how long the
// client wishes to wait for a return value from the RPC. If any value other
// than UPDATE_WORKFLOW_EXECUTION_LIFECYCLE_STAGE_COMPLETED is sent by the
// client then the RPC will complete before the update is finished and will
// return a handle to the running update so that it can later be polled for
// completion.
type UpdateWorkflowExecutionLifecycleStage int32

const (
	// An unspecified vale for this enum.
	UPDATE_WORKFLOW_EXECUTION_LIFECYCLE_STAGE_UNSPECIFIED UpdateWorkflowExecutionLifecycleStage = 0
	// The gRPC call will not return until the update request has been admitted
	// by the server - it may be the case that due to a considerations like load
	// or resource limits that an update is made to wait before the server will
	// indicate that it has been received and will be processed. This value
	// does not wait for any sort of acknowledgement from a worker.
	UPDATE_WORKFLOW_EXECUTION_LIFECYCLE_STAGE_ADMITTED UpdateWorkflowExecutionLifecycleStage = 1
	// The gRPC call will not return until the update has passed validation on
	// a worker.
	UPDATE_WORKFLOW_EXECUTION_LIFECYCLE_STAGE_ACCEPTED UpdateWorkflowExecutionLifecycleStage = 2
	// The gRPC call will not return until the update has executed to completion
	// on a worker and has either been rejected or returned a value or an error.
	UPDATE_WORKFLOW_EXECUTION_LIFECYCLE_STAGE_COMPLETED UpdateWorkflowExecutionLifecycleStage = 3
)

// Enum value maps for UpdateWorkflowExecutionLifecycleStage.
var (
	UpdateWorkflowExecutionLifecycleStage_name = map[int32]string{
		0: "UPDATE_WORKFLOW_EXECUTION_LIFECYCLE_STAGE_UNSPECIFIED",
		1: "UPDATE_WORKFLOW_EXECUTION_LIFECYCLE_STAGE_ADMITTED",
		2: "UPDATE_WORKFLOW_EXECUTION_LIFECYCLE_STAGE_ACCEPTED",
		3: "UPDATE_WORKFLOW_EXECUTION_LIFECYCLE_STAGE_COMPLETED",
	}
	UpdateWorkflowExecutionLifecycleStage_value = map[string]int32{
		"UPDATE_WORKFLOW_EXECUTION_LIFECYCLE_STAGE_UNSPECIFIED": 0,
		"UPDATE_WORKFLOW_EXECUTION_LIFECYCLE_STAGE_ADMITTED":    1,
		"UPDATE_WORKFLOW_EXECUTION_LIFECYCLE_STAGE_ACCEPTED":    2,
		"UPDATE_WORKFLOW_EXECUTION_LIFECYCLE_STAGE_COMPLETED":   3,
	}
)

func (x UpdateWorkflowExecutionLifecycleStage) Enum() *UpdateWorkflowExecutionLifecycleStage {
	p := new(UpdateWorkflowExecutionLifecycleStage)
	*p = x
	return p
}

func (x UpdateWorkflowExecutionLifecycleStage) String() string {
	switch x {
	case UPDATE_WORKFLOW_EXECUTION_LIFECYCLE_STAGE_UNSPECIFIED:
		return "Unspecified"
	case UPDATE_WORKFLOW_EXECUTION_LIFECYCLE_STAGE_ADMITTED:
		return "Admitted"
	case UPDATE_WORKFLOW_EXECUTION_LIFECYCLE_STAGE_ACCEPTED:
		return "Accepted"
	case UPDATE_WORKFLOW_EXECUTION_LIFECYCLE_STAGE_COMPLETED:
		return "Completed"
	default:
		return strconv.Itoa(int(x))
	}

}

func (UpdateWorkflowExecutionLifecycleStage) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_enums_v1_update_proto_enumTypes[0].Descriptor()
}

func (UpdateWorkflowExecutionLifecycleStage) Type() protoreflect.EnumType {
	return &file_temporal_api_enums_v1_update_proto_enumTypes[0]
}

func (x UpdateWorkflowExecutionLifecycleStage) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UpdateWorkflowExecutionLifecycleStage.Descriptor instead.
func (UpdateWorkflowExecutionLifecycleStage) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_enums_v1_update_proto_rawDescGZIP(), []int{0}
}

// Records why a WorkflowExecutionUpdateRequestedEvent was written to history.
// Note that not all update requests result in this event.
type UpdateRequestedEventOrigin int32

const (
	UPDATE_REQUESTED_EVENT_ORIGIN_UNSPECIFIED UpdateRequestedEventOrigin = 0
	// The UpdateRequested event was created when reapplying events during reset
	// or replication. I.e. an accepted update on one branch of workflow history
	// was converted into a requested update on a different branch.
	UPDATE_REQUESTED_EVENT_ORIGIN_REAPPLY UpdateRequestedEventOrigin = 1
)

// Enum value maps for UpdateRequestedEventOrigin.
var (
	UpdateRequestedEventOrigin_name = map[int32]string{
		0: "UPDATE_REQUESTED_EVENT_ORIGIN_UNSPECIFIED",
		1: "UPDATE_REQUESTED_EVENT_ORIGIN_REAPPLY",
	}
	UpdateRequestedEventOrigin_value = map[string]int32{
		"UPDATE_REQUESTED_EVENT_ORIGIN_UNSPECIFIED": 0,
		"UPDATE_REQUESTED_EVENT_ORIGIN_REAPPLY":     1,
	}
)

func (x UpdateRequestedEventOrigin) Enum() *UpdateRequestedEventOrigin {
	p := new(UpdateRequestedEventOrigin)
	*p = x
	return p
}

func (x UpdateRequestedEventOrigin) String() string {
	switch x {
	case UPDATE_REQUESTED_EVENT_ORIGIN_UNSPECIFIED:
		return "Unspecified"
	case UPDATE_REQUESTED_EVENT_ORIGIN_REAPPLY:
		return "Reapply"
	default:
		return strconv.Itoa(int(x))
	}

}

func (UpdateRequestedEventOrigin) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_enums_v1_update_proto_enumTypes[1].Descriptor()
}

func (UpdateRequestedEventOrigin) Type() protoreflect.EnumType {
	return &file_temporal_api_enums_v1_update_proto_enumTypes[1]
}

func (x UpdateRequestedEventOrigin) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UpdateRequestedEventOrigin.Descriptor instead.
func (UpdateRequestedEventOrigin) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_enums_v1_update_proto_rawDescGZIP(), []int{1}
}

var File_temporal_api_enums_v1_update_proto protoreflect.FileDescriptor

var file_temporal_api_enums_v1_update_proto_rawDesc = []byte{
	0x0a, 0x22, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2a, 0x8b, 0x02, 0x0a, 0x25,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65,
	0x53, 0x74, 0x61, 0x67, 0x65, 0x12, 0x39, 0x0a, 0x35, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f,
	0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x4c, 0x49, 0x46, 0x45, 0x43, 0x59, 0x43, 0x4c, 0x45, 0x5f, 0x53, 0x54, 0x41,
	0x47, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x36, 0x0a, 0x32, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46,
	0x4c, 0x4f, 0x57, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4c, 0x49,
	0x46, 0x45, 0x43, 0x59, 0x43, 0x4c, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x47, 0x45, 0x5f, 0x41, 0x44,
	0x4d, 0x49, 0x54, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x36, 0x0a, 0x32, 0x55, 0x50, 0x44, 0x41,
	0x54, 0x45, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x45, 0x58, 0x45, 0x43,
	0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4c, 0x49, 0x46, 0x45, 0x43, 0x59, 0x43, 0x4c, 0x45, 0x5f,
	0x53, 0x54, 0x41, 0x47, 0x45, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x45, 0x44, 0x10, 0x02,
	0x12, 0x37, 0x0a, 0x33, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46,
	0x4c, 0x4f, 0x57, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4c, 0x49,
	0x46, 0x45, 0x43, 0x59, 0x43, 0x4c, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x47, 0x45, 0x5f, 0x43, 0x4f,
	0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x03, 0x2a, 0x76, 0x0a, 0x1a, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x2d, 0x0a, 0x29, 0x55, 0x50, 0x44, 0x41, 0x54,
	0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x45, 0x44, 0x5f, 0x45, 0x56, 0x45, 0x4e,
	0x54, 0x5f, 0x4f, 0x52, 0x49, 0x47, 0x49, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x29, 0x0a, 0x25, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45,
	0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x45, 0x44, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54,
	0x5f, 0x4f, 0x52, 0x49, 0x47, 0x49, 0x4e, 0x5f, 0x52, 0x45, 0x41, 0x50, 0x50, 0x4c, 0x59, 0x10,
	0x01, 0x42, 0x83, 0x01, 0x0a, 0x18, 0x69, 0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61,
	0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0b,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x21, 0x67,
	0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0xaa, 0x02, 0x17, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x69, 0x6f, 0x2e, 0x41, 0x70,
	0x69, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x56, 0x31, 0xea, 0x02, 0x1a, 0x54, 0x65, 0x6d,
	0x70, 0x6f, 0x72, 0x61, 0x6c, 0x69, 0x6f, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x45, 0x6e,
	0x75, 0x6d, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_temporal_api_enums_v1_update_proto_rawDescOnce sync.Once
	file_temporal_api_enums_v1_update_proto_rawDescData = file_temporal_api_enums_v1_update_proto_rawDesc
)

func file_temporal_api_enums_v1_update_proto_rawDescGZIP() []byte {
	file_temporal_api_enums_v1_update_proto_rawDescOnce.Do(func() {
		file_temporal_api_enums_v1_update_proto_rawDescData = protoimpl.X.CompressGZIP(file_temporal_api_enums_v1_update_proto_rawDescData)
	})
	return file_temporal_api_enums_v1_update_proto_rawDescData
}

var file_temporal_api_enums_v1_update_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_temporal_api_enums_v1_update_proto_goTypes = []interface{}{
	(UpdateWorkflowExecutionLifecycleStage)(0), // 0: temporal.api.enums.v1.UpdateWorkflowExecutionLifecycleStage
	(UpdateRequestedEventOrigin)(0),            // 1: temporal.api.enums.v1.UpdateRequestedEventOrigin
}
var file_temporal_api_enums_v1_update_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_temporal_api_enums_v1_update_proto_init() }
func file_temporal_api_enums_v1_update_proto_init() {
	if File_temporal_api_enums_v1_update_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_temporal_api_enums_v1_update_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_api_enums_v1_update_proto_goTypes,
		DependencyIndexes: file_temporal_api_enums_v1_update_proto_depIdxs,
		EnumInfos:         file_temporal_api_enums_v1_update_proto_enumTypes,
	}.Build()
	File_temporal_api_enums_v1_update_proto = out.File
	file_temporal_api_enums_v1_update_proto_rawDesc = nil
	file_temporal_api_enums_v1_update_proto_goTypes = nil
	file_temporal_api_enums_v1_update_proto_depIdxs = nil
}