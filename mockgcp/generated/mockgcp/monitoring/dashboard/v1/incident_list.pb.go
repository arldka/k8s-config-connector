// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: mockgcp/monitoring/dashboard/v1/incident_list.proto

package dashboardpb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	monitoredres "google.golang.org/genproto/googleapis/api/monitoredres"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A widget that displays a list of incidents
type IncidentList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. The monitored resource for which incidents are listed.
	// The resource doesn't need to be fully specified. That is, you can specify
	// the resource type but not the values of the resource labels.
	// The resource type and labels are used for filtering.
	MonitoredResources []*monitoredres.MonitoredResource `protobuf:"bytes,1,rep,name=monitored_resources,json=monitoredResources,proto3" json:"monitored_resources,omitempty"`
	// Optional. A list of alert policy names to filter the incident list by.
	// Don't include the project ID prefix in the policy name. For
	// example, use `alertPolicies/utilization`.
	PolicyNames []string `protobuf:"bytes,2,rep,name=policy_names,json=policyNames,proto3" json:"policy_names,omitempty"`
}

func (x *IncidentList) Reset() {
	*x = IncidentList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_monitoring_dashboard_v1_incident_list_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncidentList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncidentList) ProtoMessage() {}

func (x *IncidentList) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_monitoring_dashboard_v1_incident_list_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncidentList.ProtoReflect.Descriptor instead.
func (*IncidentList) Descriptor() ([]byte, []int) {
	return file_mockgcp_monitoring_dashboard_v1_incident_list_proto_rawDescGZIP(), []int{0}
}

func (x *IncidentList) GetMonitoredResources() []*monitoredres.MonitoredResource {
	if x != nil {
		return x.MonitoredResources
	}
	return nil
}

func (x *IncidentList) GetPolicyNames() []string {
	if x != nil {
		return x.PolicyNames
	}
	return nil
}

var File_mockgcp_monitoring_dashboard_v1_incident_list_proto protoreflect.FileDescriptor

var file_mockgcp_monitoring_dashboard_v1_incident_list_proto_rawDesc = []byte{
	0x0a, 0x33, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x6e, 0x67, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x76,
	0x31, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x5f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x01, 0x0a,
	0x0c, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x53, 0x0a,
	0x13, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x65,
	0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x12,
	0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x12, 0x26, 0x0a, 0x0c, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0b, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x42, 0xfb, 0x01, 0x0a, 0x23, 0x63,
	0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x76, 0x31, 0x42, 0x11, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x46, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x70, 0x62, 0x3b, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x70, 0x62, 0xaa,
	0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x4d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67,
	0x5c, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x28,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x4d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x44, 0x61, 0x73, 0x68, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mockgcp_monitoring_dashboard_v1_incident_list_proto_rawDescOnce sync.Once
	file_mockgcp_monitoring_dashboard_v1_incident_list_proto_rawDescData = file_mockgcp_monitoring_dashboard_v1_incident_list_proto_rawDesc
)

func file_mockgcp_monitoring_dashboard_v1_incident_list_proto_rawDescGZIP() []byte {
	file_mockgcp_monitoring_dashboard_v1_incident_list_proto_rawDescOnce.Do(func() {
		file_mockgcp_monitoring_dashboard_v1_incident_list_proto_rawDescData = protoimpl.X.CompressGZIP(file_mockgcp_monitoring_dashboard_v1_incident_list_proto_rawDescData)
	})
	return file_mockgcp_monitoring_dashboard_v1_incident_list_proto_rawDescData
}

var file_mockgcp_monitoring_dashboard_v1_incident_list_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_mockgcp_monitoring_dashboard_v1_incident_list_proto_goTypes = []interface{}{
	(*IncidentList)(nil),                   // 0: mockgcp.monitoring.dashboard.v1.IncidentList
	(*monitoredres.MonitoredResource)(nil), // 1: google.api.MonitoredResource
}
var file_mockgcp_monitoring_dashboard_v1_incident_list_proto_depIdxs = []int32{
	1, // 0: mockgcp.monitoring.dashboard.v1.IncidentList.monitored_resources:type_name -> google.api.MonitoredResource
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_mockgcp_monitoring_dashboard_v1_incident_list_proto_init() }
func file_mockgcp_monitoring_dashboard_v1_incident_list_proto_init() {
	if File_mockgcp_monitoring_dashboard_v1_incident_list_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mockgcp_monitoring_dashboard_v1_incident_list_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncidentList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mockgcp_monitoring_dashboard_v1_incident_list_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mockgcp_monitoring_dashboard_v1_incident_list_proto_goTypes,
		DependencyIndexes: file_mockgcp_monitoring_dashboard_v1_incident_list_proto_depIdxs,
		MessageInfos:      file_mockgcp_monitoring_dashboard_v1_incident_list_proto_msgTypes,
	}.Build()
	File_mockgcp_monitoring_dashboard_v1_incident_list_proto = out.File
	file_mockgcp_monitoring_dashboard_v1_incident_list_proto_rawDesc = nil
	file_mockgcp_monitoring_dashboard_v1_incident_list_proto_goTypes = nil
	file_mockgcp_monitoring_dashboard_v1_incident_list_proto_depIdxs = nil
}
