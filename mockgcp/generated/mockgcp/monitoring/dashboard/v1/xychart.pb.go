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
// source: mockgcp/monitoring/dashboard/v1/xychart.proto

package dashboardpb

import (
	duration "github.com/golang/protobuf/ptypes/duration"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// The types of plotting strategies for data sets.
type XyChart_DataSet_PlotType int32

const (
	// Plot type is unspecified. The view will default to `LINE`.
	XyChart_DataSet_PLOT_TYPE_UNSPECIFIED XyChart_DataSet_PlotType = 0
	// The data is plotted as a set of lines (one line per series).
	XyChart_DataSet_LINE XyChart_DataSet_PlotType = 1
	// The data is plotted as a set of filled areas (one area per series),
	// with the areas stacked vertically (the base of each area is the top of
	// its predecessor, and the base of the first area is the x-axis). Since
	// the areas do not overlap, each is filled with a different opaque color.
	XyChart_DataSet_STACKED_AREA XyChart_DataSet_PlotType = 2
	// The data is plotted as a set of rectangular boxes (one box per series),
	// with the boxes stacked vertically (the base of each box is the top of
	// its predecessor, and the base of the first box is the x-axis). Since
	// the boxes do not overlap, each is filled with a different opaque color.
	XyChart_DataSet_STACKED_BAR XyChart_DataSet_PlotType = 3
	// The data is plotted as a heatmap. The series being plotted must have a
	// `DISTRIBUTION` value type. The value of each bucket in the distribution
	// is displayed as a color. This type is not currently available in the
	// Stackdriver Monitoring application.
	XyChart_DataSet_HEATMAP XyChart_DataSet_PlotType = 4
)

// Enum value maps for XyChart_DataSet_PlotType.
var (
	XyChart_DataSet_PlotType_name = map[int32]string{
		0: "PLOT_TYPE_UNSPECIFIED",
		1: "LINE",
		2: "STACKED_AREA",
		3: "STACKED_BAR",
		4: "HEATMAP",
	}
	XyChart_DataSet_PlotType_value = map[string]int32{
		"PLOT_TYPE_UNSPECIFIED": 0,
		"LINE":                  1,
		"STACKED_AREA":          2,
		"STACKED_BAR":           3,
		"HEATMAP":               4,
	}
)

func (x XyChart_DataSet_PlotType) Enum() *XyChart_DataSet_PlotType {
	p := new(XyChart_DataSet_PlotType)
	*p = x
	return p
}

func (x XyChart_DataSet_PlotType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (XyChart_DataSet_PlotType) Descriptor() protoreflect.EnumDescriptor {
	return file_mockgcp_monitoring_dashboard_v1_xychart_proto_enumTypes[0].Descriptor()
}

func (XyChart_DataSet_PlotType) Type() protoreflect.EnumType {
	return &file_mockgcp_monitoring_dashboard_v1_xychart_proto_enumTypes[0]
}

func (x XyChart_DataSet_PlotType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use XyChart_DataSet_PlotType.Descriptor instead.
func (XyChart_DataSet_PlotType) EnumDescriptor() ([]byte, []int) {
	return file_mockgcp_monitoring_dashboard_v1_xychart_proto_rawDescGZIP(), []int{0, 0, 0}
}

// An axis identifier.
type XyChart_DataSet_TargetAxis int32

const (
	// The target axis was not specified. Defaults to Y1.
	XyChart_DataSet_TARGET_AXIS_UNSPECIFIED XyChart_DataSet_TargetAxis = 0
	// The y_axis (the right axis of chart).
	XyChart_DataSet_Y1 XyChart_DataSet_TargetAxis = 1
	// The y2_axis (the left axis of chart).
	XyChart_DataSet_Y2 XyChart_DataSet_TargetAxis = 2
)

// Enum value maps for XyChart_DataSet_TargetAxis.
var (
	XyChart_DataSet_TargetAxis_name = map[int32]string{
		0: "TARGET_AXIS_UNSPECIFIED",
		1: "Y1",
		2: "Y2",
	}
	XyChart_DataSet_TargetAxis_value = map[string]int32{
		"TARGET_AXIS_UNSPECIFIED": 0,
		"Y1":                      1,
		"Y2":                      2,
	}
)

func (x XyChart_DataSet_TargetAxis) Enum() *XyChart_DataSet_TargetAxis {
	p := new(XyChart_DataSet_TargetAxis)
	*p = x
	return p
}

func (x XyChart_DataSet_TargetAxis) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (XyChart_DataSet_TargetAxis) Descriptor() protoreflect.EnumDescriptor {
	return file_mockgcp_monitoring_dashboard_v1_xychart_proto_enumTypes[1].Descriptor()
}

func (XyChart_DataSet_TargetAxis) Type() protoreflect.EnumType {
	return &file_mockgcp_monitoring_dashboard_v1_xychart_proto_enumTypes[1]
}

func (x XyChart_DataSet_TargetAxis) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use XyChart_DataSet_TargetAxis.Descriptor instead.
func (XyChart_DataSet_TargetAxis) EnumDescriptor() ([]byte, []int) {
	return file_mockgcp_monitoring_dashboard_v1_xychart_proto_rawDescGZIP(), []int{0, 0, 1}
}

// Types of scales used in axes.
type XyChart_Axis_Scale int32

const (
	// Scale is unspecified. The view will default to `LINEAR`.
	XyChart_Axis_SCALE_UNSPECIFIED XyChart_Axis_Scale = 0
	// Linear scale.
	XyChart_Axis_LINEAR XyChart_Axis_Scale = 1
	// Logarithmic scale (base 10).
	XyChart_Axis_LOG10 XyChart_Axis_Scale = 2
)

// Enum value maps for XyChart_Axis_Scale.
var (
	XyChart_Axis_Scale_name = map[int32]string{
		0: "SCALE_UNSPECIFIED",
		1: "LINEAR",
		2: "LOG10",
	}
	XyChart_Axis_Scale_value = map[string]int32{
		"SCALE_UNSPECIFIED": 0,
		"LINEAR":            1,
		"LOG10":             2,
	}
)

func (x XyChart_Axis_Scale) Enum() *XyChart_Axis_Scale {
	p := new(XyChart_Axis_Scale)
	*p = x
	return p
}

func (x XyChart_Axis_Scale) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (XyChart_Axis_Scale) Descriptor() protoreflect.EnumDescriptor {
	return file_mockgcp_monitoring_dashboard_v1_xychart_proto_enumTypes[2].Descriptor()
}

func (XyChart_Axis_Scale) Type() protoreflect.EnumType {
	return &file_mockgcp_monitoring_dashboard_v1_xychart_proto_enumTypes[2]
}

func (x XyChart_Axis_Scale) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use XyChart_Axis_Scale.Descriptor instead.
func (XyChart_Axis_Scale) EnumDescriptor() ([]byte, []int) {
	return file_mockgcp_monitoring_dashboard_v1_xychart_proto_rawDescGZIP(), []int{0, 1, 0}
}

// Chart mode options.
type ChartOptions_Mode int32

const (
	// Mode is unspecified. The view will default to `COLOR`.
	ChartOptions_MODE_UNSPECIFIED ChartOptions_Mode = 0
	// The chart distinguishes data series using different color. Line
	// colors may get reused when there are many lines in the chart.
	ChartOptions_COLOR ChartOptions_Mode = 1
	// The chart uses the Stackdriver x-ray mode, in which each
	// data set is plotted using the same semi-transparent color.
	ChartOptions_X_RAY ChartOptions_Mode = 2
	// The chart displays statistics such as average, median, 95th percentile,
	// and more.
	ChartOptions_STATS ChartOptions_Mode = 3
)

// Enum value maps for ChartOptions_Mode.
var (
	ChartOptions_Mode_name = map[int32]string{
		0: "MODE_UNSPECIFIED",
		1: "COLOR",
		2: "X_RAY",
		3: "STATS",
	}
	ChartOptions_Mode_value = map[string]int32{
		"MODE_UNSPECIFIED": 0,
		"COLOR":            1,
		"X_RAY":            2,
		"STATS":            3,
	}
)

func (x ChartOptions_Mode) Enum() *ChartOptions_Mode {
	p := new(ChartOptions_Mode)
	*p = x
	return p
}

func (x ChartOptions_Mode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChartOptions_Mode) Descriptor() protoreflect.EnumDescriptor {
	return file_mockgcp_monitoring_dashboard_v1_xychart_proto_enumTypes[3].Descriptor()
}

func (ChartOptions_Mode) Type() protoreflect.EnumType {
	return &file_mockgcp_monitoring_dashboard_v1_xychart_proto_enumTypes[3]
}

func (x ChartOptions_Mode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChartOptions_Mode.Descriptor instead.
func (ChartOptions_Mode) EnumDescriptor() ([]byte, []int) {
	return file_mockgcp_monitoring_dashboard_v1_xychart_proto_rawDescGZIP(), []int{1, 0}
}

// A chart that displays data on a 2D (X and Y axes) plane.
type XyChart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The data displayed in this chart.
	DataSets []*XyChart_DataSet `protobuf:"bytes,1,rep,name=data_sets,json=dataSets,proto3" json:"data_sets,omitempty"`
	// The duration used to display a comparison chart. A comparison chart
	// simultaneously shows values from two similar-length time periods
	// (e.g., week-over-week metrics).
	// The duration must be positive, and it can only be applied to charts with
	// data sets of LINE plot type.
	TimeshiftDuration *duration.Duration `protobuf:"bytes,4,opt,name=timeshift_duration,json=timeshiftDuration,proto3" json:"timeshift_duration,omitempty"`
	// Threshold lines drawn horizontally across the chart.
	Thresholds []*Threshold `protobuf:"bytes,5,rep,name=thresholds,proto3" json:"thresholds,omitempty"`
	// The properties applied to the x-axis.
	XAxis *XyChart_Axis `protobuf:"bytes,6,opt,name=x_axis,json=xAxis,proto3" json:"x_axis,omitempty"`
	// The properties applied to the y-axis.
	YAxis *XyChart_Axis `protobuf:"bytes,7,opt,name=y_axis,json=yAxis,proto3" json:"y_axis,omitempty"`
	// The properties applied to the y2-axis.
	Y2Axis *XyChart_Axis `protobuf:"bytes,9,opt,name=y2_axis,json=y2Axis,proto3" json:"y2_axis,omitempty"`
	// Display options for the chart.
	ChartOptions *ChartOptions `protobuf:"bytes,8,opt,name=chart_options,json=chartOptions,proto3" json:"chart_options,omitempty"`
}

func (x *XyChart) Reset() {
	*x = XyChart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_monitoring_dashboard_v1_xychart_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XyChart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XyChart) ProtoMessage() {}

func (x *XyChart) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_monitoring_dashboard_v1_xychart_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XyChart.ProtoReflect.Descriptor instead.
func (*XyChart) Descriptor() ([]byte, []int) {
	return file_mockgcp_monitoring_dashboard_v1_xychart_proto_rawDescGZIP(), []int{0}
}

func (x *XyChart) GetDataSets() []*XyChart_DataSet {
	if x != nil {
		return x.DataSets
	}
	return nil
}

func (x *XyChart) GetTimeshiftDuration() *duration.Duration {
	if x != nil {
		return x.TimeshiftDuration
	}
	return nil
}

func (x *XyChart) GetThresholds() []*Threshold {
	if x != nil {
		return x.Thresholds
	}
	return nil
}

func (x *XyChart) GetXAxis() *XyChart_Axis {
	if x != nil {
		return x.XAxis
	}
	return nil
}

func (x *XyChart) GetYAxis() *XyChart_Axis {
	if x != nil {
		return x.YAxis
	}
	return nil
}

func (x *XyChart) GetY2Axis() *XyChart_Axis {
	if x != nil {
		return x.Y2Axis
	}
	return nil
}

func (x *XyChart) GetChartOptions() *ChartOptions {
	if x != nil {
		return x.ChartOptions
	}
	return nil
}

// Options to control visual rendering of a chart.
type ChartOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The chart mode.
	Mode ChartOptions_Mode `protobuf:"varint,1,opt,name=mode,proto3,enum=mockgcp.monitoring.dashboard.v1.ChartOptions_Mode" json:"mode,omitempty"`
}

func (x *ChartOptions) Reset() {
	*x = ChartOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_monitoring_dashboard_v1_xychart_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChartOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChartOptions) ProtoMessage() {}

func (x *ChartOptions) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_monitoring_dashboard_v1_xychart_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChartOptions.ProtoReflect.Descriptor instead.
func (*ChartOptions) Descriptor() ([]byte, []int) {
	return file_mockgcp_monitoring_dashboard_v1_xychart_proto_rawDescGZIP(), []int{1}
}

func (x *ChartOptions) GetMode() ChartOptions_Mode {
	if x != nil {
		return x.Mode
	}
	return ChartOptions_MODE_UNSPECIFIED
}

// Groups a time series query definition with charting options.
type XyChart_DataSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Fields for querying time series data from the
	// Stackdriver metrics API.
	TimeSeriesQuery *TimeSeriesQuery `protobuf:"bytes,1,opt,name=time_series_query,json=timeSeriesQuery,proto3" json:"time_series_query,omitempty"`
	// How this data should be plotted on the chart.
	PlotType XyChart_DataSet_PlotType `protobuf:"varint,2,opt,name=plot_type,json=plotType,proto3,enum=mockgcp.monitoring.dashboard.v1.XyChart_DataSet_PlotType" json:"plot_type,omitempty"`
	// A template string for naming `TimeSeries` in the resulting data set.
	// This should be a string with interpolations of the form `${label_name}`,
	// which will resolve to the label's value.
	LegendTemplate string `protobuf:"bytes,3,opt,name=legend_template,json=legendTemplate,proto3" json:"legend_template,omitempty"`
	// Optional. The lower bound on data point frequency for this data set,
	// implemented by specifying the minimum alignment period to use in a time
	// series query For example, if the data is published once every 10 minutes,
	// the `min_alignment_period` should be at least 10 minutes. It would not
	// make sense to fetch and align data at one minute intervals.
	MinAlignmentPeriod *duration.Duration `protobuf:"bytes,4,opt,name=min_alignment_period,json=minAlignmentPeriod,proto3" json:"min_alignment_period,omitempty"`
	// Optional. The target axis to use for plotting the metric.
	TargetAxis XyChart_DataSet_TargetAxis `protobuf:"varint,5,opt,name=target_axis,json=targetAxis,proto3,enum=mockgcp.monitoring.dashboard.v1.XyChart_DataSet_TargetAxis" json:"target_axis,omitempty"`
}

func (x *XyChart_DataSet) Reset() {
	*x = XyChart_DataSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_monitoring_dashboard_v1_xychart_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XyChart_DataSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XyChart_DataSet) ProtoMessage() {}

func (x *XyChart_DataSet) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_monitoring_dashboard_v1_xychart_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XyChart_DataSet.ProtoReflect.Descriptor instead.
func (*XyChart_DataSet) Descriptor() ([]byte, []int) {
	return file_mockgcp_monitoring_dashboard_v1_xychart_proto_rawDescGZIP(), []int{0, 0}
}

func (x *XyChart_DataSet) GetTimeSeriesQuery() *TimeSeriesQuery {
	if x != nil {
		return x.TimeSeriesQuery
	}
	return nil
}

func (x *XyChart_DataSet) GetPlotType() XyChart_DataSet_PlotType {
	if x != nil {
		return x.PlotType
	}
	return XyChart_DataSet_PLOT_TYPE_UNSPECIFIED
}

func (x *XyChart_DataSet) GetLegendTemplate() string {
	if x != nil {
		return x.LegendTemplate
	}
	return ""
}

func (x *XyChart_DataSet) GetMinAlignmentPeriod() *duration.Duration {
	if x != nil {
		return x.MinAlignmentPeriod
	}
	return nil
}

func (x *XyChart_DataSet) GetTargetAxis() XyChart_DataSet_TargetAxis {
	if x != nil {
		return x.TargetAxis
	}
	return XyChart_DataSet_TARGET_AXIS_UNSPECIFIED
}

// A chart axis.
type XyChart_Axis struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The label of the axis.
	Label string `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	// The axis scale. By default, a linear scale is used.
	Scale XyChart_Axis_Scale `protobuf:"varint,2,opt,name=scale,proto3,enum=mockgcp.monitoring.dashboard.v1.XyChart_Axis_Scale" json:"scale,omitempty"`
}

func (x *XyChart_Axis) Reset() {
	*x = XyChart_Axis{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_monitoring_dashboard_v1_xychart_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XyChart_Axis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XyChart_Axis) ProtoMessage() {}

func (x *XyChart_Axis) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_monitoring_dashboard_v1_xychart_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XyChart_Axis.ProtoReflect.Descriptor instead.
func (*XyChart_Axis) Descriptor() ([]byte, []int) {
	return file_mockgcp_monitoring_dashboard_v1_xychart_proto_rawDescGZIP(), []int{0, 1}
}

func (x *XyChart_Axis) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *XyChart_Axis) GetScale() XyChart_Axis_Scale {
	if x != nil {
		return x.Scale
	}
	return XyChart_Axis_SCALE_UNSPECIFIED
}

var File_mockgcp_monitoring_dashboard_v1_xychart_proto protoreflect.FileDescriptor

var file_mockgcp_monitoring_dashboard_v1_xychart_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x6e, 0x67, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x76,
	0x31, 0x2f, 0x78, 0x79, 0x63, 0x68, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1f, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2d, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xfd, 0x09, 0x0a, 0x07, 0x58, 0x79, 0x43, 0x68, 0x61, 0x72, 0x74, 0x12, 0x52, 0x0a, 0x09,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x30, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x58, 0x79, 0x43, 0x68, 0x61, 0x72, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65,
	0x74, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x53, 0x65, 0x74, 0x73,
	0x12, 0x48, 0x0a, 0x12, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x69, 0x66, 0x74, 0x5f, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x69,
	0x66, 0x74, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4a, 0x0a, 0x0a, 0x74, 0x68,
	0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a,
	0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x52, 0x0a, 0x74, 0x68, 0x72, 0x65,
	0x73, 0x68, 0x6f, 0x6c, 0x64, 0x73, 0x12, 0x44, 0x0a, 0x06, 0x78, 0x5f, 0x61, 0x78, 0x69, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70,
	0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73, 0x68,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x58, 0x79, 0x43, 0x68, 0x61, 0x72, 0x74,
	0x2e, 0x41, 0x78, 0x69, 0x73, 0x52, 0x05, 0x78, 0x41, 0x78, 0x69, 0x73, 0x12, 0x44, 0x0a, 0x06,
	0x79, 0x5f, 0x61, 0x78, 0x69, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x6d,
	0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e,
	0x67, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x58,
	0x79, 0x43, 0x68, 0x61, 0x72, 0x74, 0x2e, 0x41, 0x78, 0x69, 0x73, 0x52, 0x05, 0x79, 0x41, 0x78,
	0x69, 0x73, 0x12, 0x46, 0x0a, 0x07, 0x79, 0x32, 0x5f, 0x61, 0x78, 0x69, 0x73, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x6d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x58, 0x79, 0x43, 0x68, 0x61, 0x72, 0x74, 0x2e, 0x41, 0x78,
	0x69, 0x73, 0x52, 0x06, 0x79, 0x32, 0x41, 0x78, 0x69, 0x73, 0x12, 0x52, 0x0a, 0x0d, 0x63, 0x68,
	0x61, 0x72, 0x74, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2d, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x6d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x0c, 0x63, 0x68, 0x61, 0x72, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0xbe,
	0x04, 0x0a, 0x07, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x74, 0x12, 0x61, 0x0a, 0x11, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e,
	0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69,
	0x65, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0f, 0x74, 0x69,
	0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x56, 0x0a,
	0x09, 0x70, 0x6c, 0x6f, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x39, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x58, 0x79, 0x43, 0x68, 0x61, 0x72, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53,
	0x65, 0x74, 0x2e, 0x50, 0x6c, 0x6f, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x70, 0x6c, 0x6f,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x6c, 0x65, 0x67, 0x65, 0x6e, 0x64, 0x5f,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x6c, 0x65, 0x67, 0x65, 0x6e, 0x64, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x50,
	0x0a, 0x14, 0x6d, 0x69, 0x6e, 0x5f, 0x61, 0x6c, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x12, 0x6d, 0x69,
	0x6e, 0x41, 0x6c, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64,
	0x12, 0x61, 0x0a, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x78, 0x69, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3b, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e,
	0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x58, 0x79, 0x43, 0x68, 0x61, 0x72, 0x74, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x74, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x78,
	0x69, 0x73, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41,
	0x78, 0x69, 0x73, 0x22, 0x5f, 0x0a, 0x08, 0x50, 0x6c, 0x6f, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x19, 0x0a, 0x15, 0x50, 0x4c, 0x4f, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x49,
	0x4e, 0x45, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41, 0x43, 0x4b, 0x45, 0x44, 0x5f,
	0x41, 0x52, 0x45, 0x41, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x54, 0x41, 0x43, 0x4b, 0x45,
	0x44, 0x5f, 0x42, 0x41, 0x52, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x48, 0x45, 0x41, 0x54, 0x4d,
	0x41, 0x50, 0x10, 0x04, 0x22, 0x39, 0x0a, 0x0a, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x78,
	0x69, 0x73, 0x12, 0x1b, 0x0a, 0x17, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x5f, 0x41, 0x58, 0x49,
	0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x06, 0x0a, 0x02, 0x59, 0x31, 0x10, 0x01, 0x12, 0x06, 0x0a, 0x02, 0x59, 0x32, 0x10, 0x02, 0x1a,
	0x9e, 0x01, 0x0a, 0x04, 0x41, 0x78, 0x69, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x49,
	0x0a, 0x05, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x33, 0x2e,
	0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x58, 0x79, 0x43, 0x68, 0x61, 0x72, 0x74, 0x2e, 0x41, 0x78, 0x69, 0x73, 0x2e, 0x53, 0x63, 0x61,
	0x6c, 0x65, 0x52, 0x05, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x22, 0x35, 0x0a, 0x05, 0x53, 0x63, 0x61,
	0x6c, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x43, 0x41, 0x4c, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4c, 0x49, 0x4e,
	0x45, 0x41, 0x52, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x4f, 0x47, 0x31, 0x30, 0x10, 0x02,
	0x22, 0x95, 0x01, 0x0a, 0x0c, 0x43, 0x68, 0x61, 0x72, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x46, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x32, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4d,
	0x6f, 0x64, 0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x22, 0x3d, 0x0a, 0x04, 0x4d, 0x6f, 0x64,
	0x65, 0x12, 0x14, 0x0a, 0x10, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x4f, 0x4c, 0x4f, 0x52,
	0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x58, 0x5f, 0x52, 0x41, 0x59, 0x10, 0x02, 0x12, 0x09, 0x0a,
	0x05, 0x53, 0x54, 0x41, 0x54, 0x53, 0x10, 0x03, 0x42, 0xf6, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d,
	0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31,
	0x42, 0x0c, 0x58, 0x79, 0x43, 0x68, 0x61, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x46, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67,
	0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31,
	0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x70, 0x62, 0x3b, 0x64, 0x61, 0x73,
	0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x70, 0x62, 0xaa, 0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0x2e, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x4d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x5c, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x28, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a,
	0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0x3a, 0x3a, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mockgcp_monitoring_dashboard_v1_xychart_proto_rawDescOnce sync.Once
	file_mockgcp_monitoring_dashboard_v1_xychart_proto_rawDescData = file_mockgcp_monitoring_dashboard_v1_xychart_proto_rawDesc
)

func file_mockgcp_monitoring_dashboard_v1_xychart_proto_rawDescGZIP() []byte {
	file_mockgcp_monitoring_dashboard_v1_xychart_proto_rawDescOnce.Do(func() {
		file_mockgcp_monitoring_dashboard_v1_xychart_proto_rawDescData = protoimpl.X.CompressGZIP(file_mockgcp_monitoring_dashboard_v1_xychart_proto_rawDescData)
	})
	return file_mockgcp_monitoring_dashboard_v1_xychart_proto_rawDescData
}

var file_mockgcp_monitoring_dashboard_v1_xychart_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_mockgcp_monitoring_dashboard_v1_xychart_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_mockgcp_monitoring_dashboard_v1_xychart_proto_goTypes = []interface{}{
	(XyChart_DataSet_PlotType)(0),   // 0: mockgcp.monitoring.dashboard.v1.XyChart.DataSet.PlotType
	(XyChart_DataSet_TargetAxis)(0), // 1: mockgcp.monitoring.dashboard.v1.XyChart.DataSet.TargetAxis
	(XyChart_Axis_Scale)(0),         // 2: mockgcp.monitoring.dashboard.v1.XyChart.Axis.Scale
	(ChartOptions_Mode)(0),          // 3: mockgcp.monitoring.dashboard.v1.ChartOptions.Mode
	(*XyChart)(nil),                 // 4: mockgcp.monitoring.dashboard.v1.XyChart
	(*ChartOptions)(nil),            // 5: mockgcp.monitoring.dashboard.v1.ChartOptions
	(*XyChart_DataSet)(nil),         // 6: mockgcp.monitoring.dashboard.v1.XyChart.DataSet
	(*XyChart_Axis)(nil),            // 7: mockgcp.monitoring.dashboard.v1.XyChart.Axis
	(*duration.Duration)(nil),       // 8: google.protobuf.Duration
	(*Threshold)(nil),               // 9: mockgcp.monitoring.dashboard.v1.Threshold
	(*TimeSeriesQuery)(nil),         // 10: mockgcp.monitoring.dashboard.v1.TimeSeriesQuery
}
var file_mockgcp_monitoring_dashboard_v1_xychart_proto_depIdxs = []int32{
	6,  // 0: mockgcp.monitoring.dashboard.v1.XyChart.data_sets:type_name -> mockgcp.monitoring.dashboard.v1.XyChart.DataSet
	8,  // 1: mockgcp.monitoring.dashboard.v1.XyChart.timeshift_duration:type_name -> google.protobuf.Duration
	9,  // 2: mockgcp.monitoring.dashboard.v1.XyChart.thresholds:type_name -> mockgcp.monitoring.dashboard.v1.Threshold
	7,  // 3: mockgcp.monitoring.dashboard.v1.XyChart.x_axis:type_name -> mockgcp.monitoring.dashboard.v1.XyChart.Axis
	7,  // 4: mockgcp.monitoring.dashboard.v1.XyChart.y_axis:type_name -> mockgcp.monitoring.dashboard.v1.XyChart.Axis
	7,  // 5: mockgcp.monitoring.dashboard.v1.XyChart.y2_axis:type_name -> mockgcp.monitoring.dashboard.v1.XyChart.Axis
	5,  // 6: mockgcp.monitoring.dashboard.v1.XyChart.chart_options:type_name -> mockgcp.monitoring.dashboard.v1.ChartOptions
	3,  // 7: mockgcp.monitoring.dashboard.v1.ChartOptions.mode:type_name -> mockgcp.monitoring.dashboard.v1.ChartOptions.Mode
	10, // 8: mockgcp.monitoring.dashboard.v1.XyChart.DataSet.time_series_query:type_name -> mockgcp.monitoring.dashboard.v1.TimeSeriesQuery
	0,  // 9: mockgcp.monitoring.dashboard.v1.XyChart.DataSet.plot_type:type_name -> mockgcp.monitoring.dashboard.v1.XyChart.DataSet.PlotType
	8,  // 10: mockgcp.monitoring.dashboard.v1.XyChart.DataSet.min_alignment_period:type_name -> google.protobuf.Duration
	1,  // 11: mockgcp.monitoring.dashboard.v1.XyChart.DataSet.target_axis:type_name -> mockgcp.monitoring.dashboard.v1.XyChart.DataSet.TargetAxis
	2,  // 12: mockgcp.monitoring.dashboard.v1.XyChart.Axis.scale:type_name -> mockgcp.monitoring.dashboard.v1.XyChart.Axis.Scale
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_mockgcp_monitoring_dashboard_v1_xychart_proto_init() }
func file_mockgcp_monitoring_dashboard_v1_xychart_proto_init() {
	if File_mockgcp_monitoring_dashboard_v1_xychart_proto != nil {
		return
	}
	file_mockgcp_monitoring_dashboard_v1_metrics_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mockgcp_monitoring_dashboard_v1_xychart_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XyChart); i {
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
		file_mockgcp_monitoring_dashboard_v1_xychart_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChartOptions); i {
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
		file_mockgcp_monitoring_dashboard_v1_xychart_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XyChart_DataSet); i {
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
		file_mockgcp_monitoring_dashboard_v1_xychart_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XyChart_Axis); i {
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
			RawDescriptor: file_mockgcp_monitoring_dashboard_v1_xychart_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mockgcp_monitoring_dashboard_v1_xychart_proto_goTypes,
		DependencyIndexes: file_mockgcp_monitoring_dashboard_v1_xychart_proto_depIdxs,
		EnumInfos:         file_mockgcp_monitoring_dashboard_v1_xychart_proto_enumTypes,
		MessageInfos:      file_mockgcp_monitoring_dashboard_v1_xychart_proto_msgTypes,
	}.Build()
	File_mockgcp_monitoring_dashboard_v1_xychart_proto = out.File
	file_mockgcp_monitoring_dashboard_v1_xychart_proto_rawDesc = nil
	file_mockgcp_monitoring_dashboard_v1_xychart_proto_goTypes = nil
	file_mockgcp_monitoring_dashboard_v1_xychart_proto_depIdxs = nil
}
