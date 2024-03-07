// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: junction/junction/query.proto

package junction

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Query_Params_FullMethodName                      = "/junction.junction.Query/Params"
	Query_GetStation_FullMethodName                  = "/junction.junction.Query/GetStation"
	Query_ListStations_FullMethodName                = "/junction.junction.Query/ListStations"
	Query_GetStationDetailsByAddress_FullMethodName  = "/junction.junction.Query/GetStationDetailsByAddress"
	Query_GetPod_FullMethodName                      = "/junction.junction.Query/GetPod"
	Query_GetLatestSubmittedPodNumber_FullMethodName = "/junction.junction.Query/GetLatestSubmittedPodNumber"
	Query_GetLatestVerifiedPodNumber_FullMethodName  = "/junction.junction.Query/GetLatestVerifiedPodNumber"
	Query_FetchVrn_FullMethodName                    = "/junction.junction.Query/FetchVrn"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of GetStation items.
	GetStation(ctx context.Context, in *QueryGetStationRequest, opts ...grpc.CallOption) (*QueryGetStationResponse, error)
	// Queries a list of ListStations items.
	ListStations(ctx context.Context, in *QueryListStationsRequest, opts ...grpc.CallOption) (*QueryListStationsResponse, error)
	// Queries a list of GetStationDetailsByAddress items.
	GetStationDetailsByAddress(ctx context.Context, in *QueryGetStationDetailsByAddressRequest, opts ...grpc.CallOption) (*QueryGetStationDetailsByAddressResponse, error)
	// Queries a list of GetPod items.
	GetPod(ctx context.Context, in *QueryGetPodRequest, opts ...grpc.CallOption) (*QueryGetPodResponse, error)
	// Queries a list of GetLatestSubmittedPodNumber items.
	GetLatestSubmittedPodNumber(ctx context.Context, in *QueryGetLatestSubmittedPodNumberRequest, opts ...grpc.CallOption) (*QueryGetLatestSubmittedPodNumberResponse, error)
	// Queries a list of GetLatestVerifiedPodNumber items.
	GetLatestVerifiedPodNumber(ctx context.Context, in *QueryGetLatestVerifiedPodNumberRequest, opts ...grpc.CallOption) (*QueryGetLatestVerifiedPodNumberResponse, error)
	// Queries a list of FetchVrn items.
	FetchVrn(ctx context.Context, in *QueryFetchVrnRequest, opts ...grpc.CallOption) (*QueryFetchVrnResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, Query_Params_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetStation(ctx context.Context, in *QueryGetStationRequest, opts ...grpc.CallOption) (*QueryGetStationResponse, error) {
	out := new(QueryGetStationResponse)
	err := c.cc.Invoke(ctx, Query_GetStation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ListStations(ctx context.Context, in *QueryListStationsRequest, opts ...grpc.CallOption) (*QueryListStationsResponse, error) {
	out := new(QueryListStationsResponse)
	err := c.cc.Invoke(ctx, Query_ListStations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetStationDetailsByAddress(ctx context.Context, in *QueryGetStationDetailsByAddressRequest, opts ...grpc.CallOption) (*QueryGetStationDetailsByAddressResponse, error) {
	out := new(QueryGetStationDetailsByAddressResponse)
	err := c.cc.Invoke(ctx, Query_GetStationDetailsByAddress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetPod(ctx context.Context, in *QueryGetPodRequest, opts ...grpc.CallOption) (*QueryGetPodResponse, error) {
	out := new(QueryGetPodResponse)
	err := c.cc.Invoke(ctx, Query_GetPod_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetLatestSubmittedPodNumber(ctx context.Context, in *QueryGetLatestSubmittedPodNumberRequest, opts ...grpc.CallOption) (*QueryGetLatestSubmittedPodNumberResponse, error) {
	out := new(QueryGetLatestSubmittedPodNumberResponse)
	err := c.cc.Invoke(ctx, Query_GetLatestSubmittedPodNumber_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetLatestVerifiedPodNumber(ctx context.Context, in *QueryGetLatestVerifiedPodNumberRequest, opts ...grpc.CallOption) (*QueryGetLatestVerifiedPodNumberResponse, error) {
	out := new(QueryGetLatestVerifiedPodNumberResponse)
	err := c.cc.Invoke(ctx, Query_GetLatestVerifiedPodNumber_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) FetchVrn(ctx context.Context, in *QueryFetchVrnRequest, opts ...grpc.CallOption) (*QueryFetchVrnResponse, error) {
	out := new(QueryFetchVrnResponse)
	err := c.cc.Invoke(ctx, Query_FetchVrn_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of GetStation items.
	GetStation(context.Context, *QueryGetStationRequest) (*QueryGetStationResponse, error)
	// Queries a list of ListStations items.
	ListStations(context.Context, *QueryListStationsRequest) (*QueryListStationsResponse, error)
	// Queries a list of GetStationDetailsByAddress items.
	GetStationDetailsByAddress(context.Context, *QueryGetStationDetailsByAddressRequest) (*QueryGetStationDetailsByAddressResponse, error)
	// Queries a list of GetPod items.
	GetPod(context.Context, *QueryGetPodRequest) (*QueryGetPodResponse, error)
	// Queries a list of GetLatestSubmittedPodNumber items.
	GetLatestSubmittedPodNumber(context.Context, *QueryGetLatestSubmittedPodNumberRequest) (*QueryGetLatestSubmittedPodNumberResponse, error)
	// Queries a list of GetLatestVerifiedPodNumber items.
	GetLatestVerifiedPodNumber(context.Context, *QueryGetLatestVerifiedPodNumberRequest) (*QueryGetLatestVerifiedPodNumberResponse, error)
	// Queries a list of FetchVrn items.
	FetchVrn(context.Context, *QueryFetchVrnRequest) (*QueryFetchVrnResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) GetStation(context.Context, *QueryGetStationRequest) (*QueryGetStationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStation not implemented")
}
func (UnimplementedQueryServer) ListStations(context.Context, *QueryListStationsRequest) (*QueryListStationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListStations not implemented")
}
func (UnimplementedQueryServer) GetStationDetailsByAddress(context.Context, *QueryGetStationDetailsByAddressRequest) (*QueryGetStationDetailsByAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStationDetailsByAddress not implemented")
}
func (UnimplementedQueryServer) GetPod(context.Context, *QueryGetPodRequest) (*QueryGetPodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPod not implemented")
}
func (UnimplementedQueryServer) GetLatestSubmittedPodNumber(context.Context, *QueryGetLatestSubmittedPodNumberRequest) (*QueryGetLatestSubmittedPodNumberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatestSubmittedPodNumber not implemented")
}
func (UnimplementedQueryServer) GetLatestVerifiedPodNumber(context.Context, *QueryGetLatestVerifiedPodNumberRequest) (*QueryGetLatestVerifiedPodNumberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatestVerifiedPodNumber not implemented")
}
func (UnimplementedQueryServer) FetchVrn(context.Context, *QueryFetchVrnRequest) (*QueryFetchVrnResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchVrn not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Params_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetStation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetStationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetStation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetStation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetStation(ctx, req.(*QueryGetStationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ListStations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryListStationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ListStations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ListStations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ListStations(ctx, req.(*QueryListStationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetStationDetailsByAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetStationDetailsByAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetStationDetailsByAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetStationDetailsByAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetStationDetailsByAddress(ctx, req.(*QueryGetStationDetailsByAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetPod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetPodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetPod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetPod_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetPod(ctx, req.(*QueryGetPodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetLatestSubmittedPodNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetLatestSubmittedPodNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetLatestSubmittedPodNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetLatestSubmittedPodNumber_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetLatestSubmittedPodNumber(ctx, req.(*QueryGetLatestSubmittedPodNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetLatestVerifiedPodNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetLatestVerifiedPodNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetLatestVerifiedPodNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetLatestVerifiedPodNumber_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetLatestVerifiedPodNumber(ctx, req.(*QueryGetLatestVerifiedPodNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_FetchVrn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFetchVrnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).FetchVrn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_FetchVrn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).FetchVrn(ctx, req.(*QueryFetchVrnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "junction.junction.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "GetStation",
			Handler:    _Query_GetStation_Handler,
		},
		{
			MethodName: "ListStations",
			Handler:    _Query_ListStations_Handler,
		},
		{
			MethodName: "GetStationDetailsByAddress",
			Handler:    _Query_GetStationDetailsByAddress_Handler,
		},
		{
			MethodName: "GetPod",
			Handler:    _Query_GetPod_Handler,
		},
		{
			MethodName: "GetLatestSubmittedPodNumber",
			Handler:    _Query_GetLatestSubmittedPodNumber_Handler,
		},
		{
			MethodName: "GetLatestVerifiedPodNumber",
			Handler:    _Query_GetLatestVerifiedPodNumber_Handler,
		},
		{
			MethodName: "FetchVrn",
			Handler:    _Query_FetchVrn_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "junction/junction/query.proto",
}
