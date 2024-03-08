// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: mockgcp/cloud/aiplatform/v1beta1/vizier_service.proto

package aiplatformpb

import (
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// VizierServiceClient is the client API for VizierService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VizierServiceClient interface {
	// Creates a Study. A resource name will be generated after creation of the
	// Study.
	CreateStudy(ctx context.Context, in *CreateStudyRequest, opts ...grpc.CallOption) (*Study, error)
	// Gets a Study by name.
	GetStudy(ctx context.Context, in *GetStudyRequest, opts ...grpc.CallOption) (*Study, error)
	// Lists all the studies in a region for an associated project.
	ListStudies(ctx context.Context, in *ListStudiesRequest, opts ...grpc.CallOption) (*ListStudiesResponse, error)
	// Deletes a Study.
	DeleteStudy(ctx context.Context, in *DeleteStudyRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Looks a study up using the user-defined display_name field instead of the
	// fully qualified resource name.
	LookupStudy(ctx context.Context, in *LookupStudyRequest, opts ...grpc.CallOption) (*Study, error)
	// Adds one or more Trials to a Study, with parameter values
	// suggested by Vertex AI Vizier. Returns a long-running
	// operation associated with the generation of Trial suggestions.
	// When this long-running operation succeeds, it will contain
	// a
	// [SuggestTrialsResponse][mockgcp.cloud.aiplatform.v1beta1.SuggestTrialsResponse].
	SuggestTrials(ctx context.Context, in *SuggestTrialsRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Adds a user provided Trial to a Study.
	CreateTrial(ctx context.Context, in *CreateTrialRequest, opts ...grpc.CallOption) (*Trial, error)
	// Gets a Trial.
	GetTrial(ctx context.Context, in *GetTrialRequest, opts ...grpc.CallOption) (*Trial, error)
	// Lists the Trials associated with a Study.
	ListTrials(ctx context.Context, in *ListTrialsRequest, opts ...grpc.CallOption) (*ListTrialsResponse, error)
	// Adds a measurement of the objective metrics to a Trial. This measurement
	// is assumed to have been taken before the Trial is complete.
	AddTrialMeasurement(ctx context.Context, in *AddTrialMeasurementRequest, opts ...grpc.CallOption) (*Trial, error)
	// Marks a Trial as complete.
	CompleteTrial(ctx context.Context, in *CompleteTrialRequest, opts ...grpc.CallOption) (*Trial, error)
	// Deletes a Trial.
	DeleteTrial(ctx context.Context, in *DeleteTrialRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Checks  whether a Trial should stop or not. Returns a
	// long-running operation. When the operation is successful,
	// it will contain a
	// [CheckTrialEarlyStoppingStateResponse][mockgcp.cloud.aiplatform.v1beta1.CheckTrialEarlyStoppingStateResponse].
	CheckTrialEarlyStoppingState(ctx context.Context, in *CheckTrialEarlyStoppingStateRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Stops a Trial.
	StopTrial(ctx context.Context, in *StopTrialRequest, opts ...grpc.CallOption) (*Trial, error)
	// Lists the pareto-optimal Trials for multi-objective Study or the
	// optimal Trials for single-objective Study. The definition of
	// pareto-optimal can be checked in wiki page.
	// https://en.wikipedia.org/wiki/Pareto_efficiency
	ListOptimalTrials(ctx context.Context, in *ListOptimalTrialsRequest, opts ...grpc.CallOption) (*ListOptimalTrialsResponse, error)
}

type vizierServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVizierServiceClient(cc grpc.ClientConnInterface) VizierServiceClient {
	return &vizierServiceClient{cc}
}

func (c *vizierServiceClient) CreateStudy(ctx context.Context, in *CreateStudyRequest, opts ...grpc.CallOption) (*Study, error) {
	out := new(Study)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.aiplatform.v1beta1.VizierService/CreateStudy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vizierServiceClient) GetStudy(ctx context.Context, in *GetStudyRequest, opts ...grpc.CallOption) (*Study, error) {
	out := new(Study)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.aiplatform.v1beta1.VizierService/GetStudy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vizierServiceClient) ListStudies(ctx context.Context, in *ListStudiesRequest, opts ...grpc.CallOption) (*ListStudiesResponse, error) {
	out := new(ListStudiesResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.aiplatform.v1beta1.VizierService/ListStudies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vizierServiceClient) DeleteStudy(ctx context.Context, in *DeleteStudyRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.aiplatform.v1beta1.VizierService/DeleteStudy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vizierServiceClient) LookupStudy(ctx context.Context, in *LookupStudyRequest, opts ...grpc.CallOption) (*Study, error) {
	out := new(Study)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.aiplatform.v1beta1.VizierService/LookupStudy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vizierServiceClient) SuggestTrials(ctx context.Context, in *SuggestTrialsRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.aiplatform.v1beta1.VizierService/SuggestTrials", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vizierServiceClient) CreateTrial(ctx context.Context, in *CreateTrialRequest, opts ...grpc.CallOption) (*Trial, error) {
	out := new(Trial)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.aiplatform.v1beta1.VizierService/CreateTrial", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vizierServiceClient) GetTrial(ctx context.Context, in *GetTrialRequest, opts ...grpc.CallOption) (*Trial, error) {
	out := new(Trial)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.aiplatform.v1beta1.VizierService/GetTrial", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vizierServiceClient) ListTrials(ctx context.Context, in *ListTrialsRequest, opts ...grpc.CallOption) (*ListTrialsResponse, error) {
	out := new(ListTrialsResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.aiplatform.v1beta1.VizierService/ListTrials", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vizierServiceClient) AddTrialMeasurement(ctx context.Context, in *AddTrialMeasurementRequest, opts ...grpc.CallOption) (*Trial, error) {
	out := new(Trial)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.aiplatform.v1beta1.VizierService/AddTrialMeasurement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vizierServiceClient) CompleteTrial(ctx context.Context, in *CompleteTrialRequest, opts ...grpc.CallOption) (*Trial, error) {
	out := new(Trial)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.aiplatform.v1beta1.VizierService/CompleteTrial", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vizierServiceClient) DeleteTrial(ctx context.Context, in *DeleteTrialRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.aiplatform.v1beta1.VizierService/DeleteTrial", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vizierServiceClient) CheckTrialEarlyStoppingState(ctx context.Context, in *CheckTrialEarlyStoppingStateRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.aiplatform.v1beta1.VizierService/CheckTrialEarlyStoppingState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vizierServiceClient) StopTrial(ctx context.Context, in *StopTrialRequest, opts ...grpc.CallOption) (*Trial, error) {
	out := new(Trial)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.aiplatform.v1beta1.VizierService/StopTrial", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vizierServiceClient) ListOptimalTrials(ctx context.Context, in *ListOptimalTrialsRequest, opts ...grpc.CallOption) (*ListOptimalTrialsResponse, error) {
	out := new(ListOptimalTrialsResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.aiplatform.v1beta1.VizierService/ListOptimalTrials", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VizierServiceServer is the server API for VizierService service.
// All implementations must embed UnimplementedVizierServiceServer
// for forward compatibility
type VizierServiceServer interface {
	// Creates a Study. A resource name will be generated after creation of the
	// Study.
	CreateStudy(context.Context, *CreateStudyRequest) (*Study, error)
	// Gets a Study by name.
	GetStudy(context.Context, *GetStudyRequest) (*Study, error)
	// Lists all the studies in a region for an associated project.
	ListStudies(context.Context, *ListStudiesRequest) (*ListStudiesResponse, error)
	// Deletes a Study.
	DeleteStudy(context.Context, *DeleteStudyRequest) (*empty.Empty, error)
	// Looks a study up using the user-defined display_name field instead of the
	// fully qualified resource name.
	LookupStudy(context.Context, *LookupStudyRequest) (*Study, error)
	// Adds one or more Trials to a Study, with parameter values
	// suggested by Vertex AI Vizier. Returns a long-running
	// operation associated with the generation of Trial suggestions.
	// When this long-running operation succeeds, it will contain
	// a
	// [SuggestTrialsResponse][mockgcp.cloud.aiplatform.v1beta1.SuggestTrialsResponse].
	SuggestTrials(context.Context, *SuggestTrialsRequest) (*longrunningpb.Operation, error)
	// Adds a user provided Trial to a Study.
	CreateTrial(context.Context, *CreateTrialRequest) (*Trial, error)
	// Gets a Trial.
	GetTrial(context.Context, *GetTrialRequest) (*Trial, error)
	// Lists the Trials associated with a Study.
	ListTrials(context.Context, *ListTrialsRequest) (*ListTrialsResponse, error)
	// Adds a measurement of the objective metrics to a Trial. This measurement
	// is assumed to have been taken before the Trial is complete.
	AddTrialMeasurement(context.Context, *AddTrialMeasurementRequest) (*Trial, error)
	// Marks a Trial as complete.
	CompleteTrial(context.Context, *CompleteTrialRequest) (*Trial, error)
	// Deletes a Trial.
	DeleteTrial(context.Context, *DeleteTrialRequest) (*empty.Empty, error)
	// Checks  whether a Trial should stop or not. Returns a
	// long-running operation. When the operation is successful,
	// it will contain a
	// [CheckTrialEarlyStoppingStateResponse][mockgcp.cloud.aiplatform.v1beta1.CheckTrialEarlyStoppingStateResponse].
	CheckTrialEarlyStoppingState(context.Context, *CheckTrialEarlyStoppingStateRequest) (*longrunningpb.Operation, error)
	// Stops a Trial.
	StopTrial(context.Context, *StopTrialRequest) (*Trial, error)
	// Lists the pareto-optimal Trials for multi-objective Study or the
	// optimal Trials for single-objective Study. The definition of
	// pareto-optimal can be checked in wiki page.
	// https://en.wikipedia.org/wiki/Pareto_efficiency
	ListOptimalTrials(context.Context, *ListOptimalTrialsRequest) (*ListOptimalTrialsResponse, error)
	mustEmbedUnimplementedVizierServiceServer()
}

// UnimplementedVizierServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVizierServiceServer struct {
}

func (UnimplementedVizierServiceServer) CreateStudy(context.Context, *CreateStudyRequest) (*Study, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStudy not implemented")
}
func (UnimplementedVizierServiceServer) GetStudy(context.Context, *GetStudyRequest) (*Study, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudy not implemented")
}
func (UnimplementedVizierServiceServer) ListStudies(context.Context, *ListStudiesRequest) (*ListStudiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListStudies not implemented")
}
func (UnimplementedVizierServiceServer) DeleteStudy(context.Context, *DeleteStudyRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStudy not implemented")
}
func (UnimplementedVizierServiceServer) LookupStudy(context.Context, *LookupStudyRequest) (*Study, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LookupStudy not implemented")
}
func (UnimplementedVizierServiceServer) SuggestTrials(context.Context, *SuggestTrialsRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SuggestTrials not implemented")
}
func (UnimplementedVizierServiceServer) CreateTrial(context.Context, *CreateTrialRequest) (*Trial, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTrial not implemented")
}
func (UnimplementedVizierServiceServer) GetTrial(context.Context, *GetTrialRequest) (*Trial, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrial not implemented")
}
func (UnimplementedVizierServiceServer) ListTrials(context.Context, *ListTrialsRequest) (*ListTrialsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTrials not implemented")
}
func (UnimplementedVizierServiceServer) AddTrialMeasurement(context.Context, *AddTrialMeasurementRequest) (*Trial, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTrialMeasurement not implemented")
}
func (UnimplementedVizierServiceServer) CompleteTrial(context.Context, *CompleteTrialRequest) (*Trial, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteTrial not implemented")
}
func (UnimplementedVizierServiceServer) DeleteTrial(context.Context, *DeleteTrialRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTrial not implemented")
}
func (UnimplementedVizierServiceServer) CheckTrialEarlyStoppingState(context.Context, *CheckTrialEarlyStoppingStateRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckTrialEarlyStoppingState not implemented")
}
func (UnimplementedVizierServiceServer) StopTrial(context.Context, *StopTrialRequest) (*Trial, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopTrial not implemented")
}
func (UnimplementedVizierServiceServer) ListOptimalTrials(context.Context, *ListOptimalTrialsRequest) (*ListOptimalTrialsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOptimalTrials not implemented")
}
func (UnimplementedVizierServiceServer) mustEmbedUnimplementedVizierServiceServer() {}

// UnsafeVizierServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VizierServiceServer will
// result in compilation errors.
type UnsafeVizierServiceServer interface {
	mustEmbedUnimplementedVizierServiceServer()
}

func RegisterVizierServiceServer(s grpc.ServiceRegistrar, srv VizierServiceServer) {
	s.RegisterService(&VizierService_ServiceDesc, srv)
}

func _VizierService_CreateStudy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStudyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VizierServiceServer).CreateStudy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.aiplatform.v1beta1.VizierService/CreateStudy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VizierServiceServer).CreateStudy(ctx, req.(*CreateStudyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VizierService_GetStudy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStudyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VizierServiceServer).GetStudy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.aiplatform.v1beta1.VizierService/GetStudy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VizierServiceServer).GetStudy(ctx, req.(*GetStudyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VizierService_ListStudies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStudiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VizierServiceServer).ListStudies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.aiplatform.v1beta1.VizierService/ListStudies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VizierServiceServer).ListStudies(ctx, req.(*ListStudiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VizierService_DeleteStudy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStudyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VizierServiceServer).DeleteStudy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.aiplatform.v1beta1.VizierService/DeleteStudy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VizierServiceServer).DeleteStudy(ctx, req.(*DeleteStudyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VizierService_LookupStudy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupStudyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VizierServiceServer).LookupStudy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.aiplatform.v1beta1.VizierService/LookupStudy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VizierServiceServer).LookupStudy(ctx, req.(*LookupStudyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VizierService_SuggestTrials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuggestTrialsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VizierServiceServer).SuggestTrials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.aiplatform.v1beta1.VizierService/SuggestTrials",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VizierServiceServer).SuggestTrials(ctx, req.(*SuggestTrialsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VizierService_CreateTrial_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTrialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VizierServiceServer).CreateTrial(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.aiplatform.v1beta1.VizierService/CreateTrial",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VizierServiceServer).CreateTrial(ctx, req.(*CreateTrialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VizierService_GetTrial_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTrialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VizierServiceServer).GetTrial(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.aiplatform.v1beta1.VizierService/GetTrial",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VizierServiceServer).GetTrial(ctx, req.(*GetTrialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VizierService_ListTrials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTrialsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VizierServiceServer).ListTrials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.aiplatform.v1beta1.VizierService/ListTrials",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VizierServiceServer).ListTrials(ctx, req.(*ListTrialsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VizierService_AddTrialMeasurement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTrialMeasurementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VizierServiceServer).AddTrialMeasurement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.aiplatform.v1beta1.VizierService/AddTrialMeasurement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VizierServiceServer).AddTrialMeasurement(ctx, req.(*AddTrialMeasurementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VizierService_CompleteTrial_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteTrialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VizierServiceServer).CompleteTrial(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.aiplatform.v1beta1.VizierService/CompleteTrial",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VizierServiceServer).CompleteTrial(ctx, req.(*CompleteTrialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VizierService_DeleteTrial_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTrialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VizierServiceServer).DeleteTrial(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.aiplatform.v1beta1.VizierService/DeleteTrial",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VizierServiceServer).DeleteTrial(ctx, req.(*DeleteTrialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VizierService_CheckTrialEarlyStoppingState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckTrialEarlyStoppingStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VizierServiceServer).CheckTrialEarlyStoppingState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.aiplatform.v1beta1.VizierService/CheckTrialEarlyStoppingState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VizierServiceServer).CheckTrialEarlyStoppingState(ctx, req.(*CheckTrialEarlyStoppingStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VizierService_StopTrial_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopTrialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VizierServiceServer).StopTrial(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.aiplatform.v1beta1.VizierService/StopTrial",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VizierServiceServer).StopTrial(ctx, req.(*StopTrialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VizierService_ListOptimalTrials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOptimalTrialsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VizierServiceServer).ListOptimalTrials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.aiplatform.v1beta1.VizierService/ListOptimalTrials",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VizierServiceServer).ListOptimalTrials(ctx, req.(*ListOptimalTrialsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VizierService_ServiceDesc is the grpc.ServiceDesc for VizierService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VizierService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mockgcp.cloud.aiplatform.v1beta1.VizierService",
	HandlerType: (*VizierServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateStudy",
			Handler:    _VizierService_CreateStudy_Handler,
		},
		{
			MethodName: "GetStudy",
			Handler:    _VizierService_GetStudy_Handler,
		},
		{
			MethodName: "ListStudies",
			Handler:    _VizierService_ListStudies_Handler,
		},
		{
			MethodName: "DeleteStudy",
			Handler:    _VizierService_DeleteStudy_Handler,
		},
		{
			MethodName: "LookupStudy",
			Handler:    _VizierService_LookupStudy_Handler,
		},
		{
			MethodName: "SuggestTrials",
			Handler:    _VizierService_SuggestTrials_Handler,
		},
		{
			MethodName: "CreateTrial",
			Handler:    _VizierService_CreateTrial_Handler,
		},
		{
			MethodName: "GetTrial",
			Handler:    _VizierService_GetTrial_Handler,
		},
		{
			MethodName: "ListTrials",
			Handler:    _VizierService_ListTrials_Handler,
		},
		{
			MethodName: "AddTrialMeasurement",
			Handler:    _VizierService_AddTrialMeasurement_Handler,
		},
		{
			MethodName: "CompleteTrial",
			Handler:    _VizierService_CompleteTrial_Handler,
		},
		{
			MethodName: "DeleteTrial",
			Handler:    _VizierService_DeleteTrial_Handler,
		},
		{
			MethodName: "CheckTrialEarlyStoppingState",
			Handler:    _VizierService_CheckTrialEarlyStoppingState_Handler,
		},
		{
			MethodName: "StopTrial",
			Handler:    _VizierService_StopTrial_Handler,
		},
		{
			MethodName: "ListOptimalTrials",
			Handler:    _VizierService_ListOptimalTrials_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mockgcp/cloud/aiplatform/v1beta1/vizier_service.proto",
}