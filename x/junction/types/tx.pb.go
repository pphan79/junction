// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: junction/junction/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgUpdateParams is the Msg/UpdateParams request type.
type MsgUpdateParams struct {
	// authority is the address that controls the module (defaults to x/gov unless overwritten).
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// NOTE: All parameters must be supplied.
	Params Params `protobuf:"bytes,2,opt,name=params,proto3" json:"params"`
}

func (m *MsgUpdateParams) Reset()         { *m = MsgUpdateParams{} }
func (m *MsgUpdateParams) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateParams) ProtoMessage()    {}
func (*MsgUpdateParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_04349ac28bbdc1dc, []int{0}
}
func (m *MsgUpdateParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateParams.Merge(m, src)
}
func (m *MsgUpdateParams) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateParams) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateParams.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateParams proto.InternalMessageInfo

func (m *MsgUpdateParams) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *MsgUpdateParams) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// MsgUpdateParamsResponse defines the response structure for executing a
// MsgUpdateParams message.
type MsgUpdateParamsResponse struct {
}

func (m *MsgUpdateParamsResponse) Reset()         { *m = MsgUpdateParamsResponse{} }
func (m *MsgUpdateParamsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateParamsResponse) ProtoMessage()    {}
func (*MsgUpdateParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_04349ac28bbdc1dc, []int{1}
}
func (m *MsgUpdateParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateParamsResponse.Merge(m, src)
}
func (m *MsgUpdateParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateParamsResponse proto.InternalMessageInfo

type MsgInitStation struct {
	Creator         string   `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Tracks          []string `protobuf:"bytes,2,rep,name=tracks,proto3" json:"tracks,omitempty"`
	VerificationKey []byte   `protobuf:"bytes,3,opt,name=verificationKey,proto3" json:"verificationKey,omitempty"`
	StationId       string   `protobuf:"bytes,4,opt,name=stationId,proto3" json:"stationId,omitempty"`
	StationInfo     string   `protobuf:"bytes,5,opt,name=stationInfo,proto3" json:"stationInfo,omitempty"`
}

func (m *MsgInitStation) Reset()         { *m = MsgInitStation{} }
func (m *MsgInitStation) String() string { return proto.CompactTextString(m) }
func (*MsgInitStation) ProtoMessage()    {}
func (*MsgInitStation) Descriptor() ([]byte, []int) {
	return fileDescriptor_04349ac28bbdc1dc, []int{2}
}
func (m *MsgInitStation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgInitStation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgInitStation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgInitStation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgInitStation.Merge(m, src)
}
func (m *MsgInitStation) XXX_Size() int {
	return m.Size()
}
func (m *MsgInitStation) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgInitStation.DiscardUnknown(m)
}

var xxx_messageInfo_MsgInitStation proto.InternalMessageInfo

func (m *MsgInitStation) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgInitStation) GetTracks() []string {
	if m != nil {
		return m.Tracks
	}
	return nil
}

func (m *MsgInitStation) GetVerificationKey() []byte {
	if m != nil {
		return m.VerificationKey
	}
	return nil
}

func (m *MsgInitStation) GetStationId() string {
	if m != nil {
		return m.StationId
	}
	return ""
}

func (m *MsgInitStation) GetStationInfo() string {
	if m != nil {
		return m.StationInfo
	}
	return ""
}

type MsgInitStationResponse struct {
	Status    bool   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	StationId string `protobuf:"bytes,2,opt,name=stationId,proto3" json:"stationId,omitempty"`
}

func (m *MsgInitStationResponse) Reset()         { *m = MsgInitStationResponse{} }
func (m *MsgInitStationResponse) String() string { return proto.CompactTextString(m) }
func (*MsgInitStationResponse) ProtoMessage()    {}
func (*MsgInitStationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_04349ac28bbdc1dc, []int{3}
}
func (m *MsgInitStationResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgInitStationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgInitStationResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgInitStationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgInitStationResponse.Merge(m, src)
}
func (m *MsgInitStationResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgInitStationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgInitStationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgInitStationResponse proto.InternalMessageInfo

func (m *MsgInitStationResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *MsgInitStationResponse) GetStationId() string {
	if m != nil {
		return m.StationId
	}
	return ""
}

func init() {
	proto.RegisterType((*MsgUpdateParams)(nil), "junction.junction.MsgUpdateParams")
	proto.RegisterType((*MsgUpdateParamsResponse)(nil), "junction.junction.MsgUpdateParamsResponse")
	proto.RegisterType((*MsgInitStation)(nil), "junction.junction.MsgInitStation")
	proto.RegisterType((*MsgInitStationResponse)(nil), "junction.junction.MsgInitStationResponse")
}

func init() { proto.RegisterFile("junction/junction/tx.proto", fileDescriptor_04349ac28bbdc1dc) }

var fileDescriptor_04349ac28bbdc1dc = []byte{
	// 494 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xbf, 0x6f, 0xd3, 0x40,
	0x14, 0xce, 0x35, 0x34, 0xe0, 0x4b, 0x44, 0xd5, 0x53, 0xd5, 0x3a, 0x16, 0x32, 0xc1, 0x2c, 0x26,
	0x12, 0xb6, 0x08, 0x52, 0x87, 0x8a, 0x85, 0x30, 0x55, 0x25, 0x08, 0xb9, 0x62, 0x01, 0x09, 0x74,
	0xb5, 0xaf, 0xae, 0x41, 0xf6, 0x59, 0x77, 0xe7, 0xaa, 0xd9, 0x10, 0x23, 0x13, 0x7f, 0x06, 0x63,
	0x06, 0x26, 0x16, 0xd6, 0xb2, 0x55, 0x4c, 0x4c, 0x08, 0x25, 0x43, 0xfe, 0x8d, 0xca, 0xbe, 0x73,
	0x7e, 0xb8, 0x91, 0xb2, 0xd8, 0xef, 0xfb, 0xde, 0xaf, 0xef, 0xbd, 0xbb, 0x83, 0xc6, 0xc7, 0x2c,
	0xf1, 0x45, 0x44, 0x13, 0x77, 0x66, 0x88, 0x0b, 0x27, 0x65, 0x54, 0x50, 0xb4, 0x5d, 0x52, 0x4e,
	0x69, 0x18, 0xdb, 0x38, 0x8e, 0x12, 0xea, 0x16, 0x5f, 0x19, 0x65, 0xec, 0xf9, 0x94, 0xc7, 0x94,
	0xbb, 0x31, 0x0f, 0xdd, 0xf3, 0x27, 0xf9, 0x4f, 0x39, 0xda, 0xd2, 0xf1, 0xa1, 0x40, 0xae, 0x04,
	0xca, 0xb5, 0x13, 0xd2, 0x90, 0x4a, 0x3e, 0xb7, 0x14, 0x6b, 0xde, 0xd4, 0x92, 0x62, 0x86, 0x63,
	0x95, 0x65, 0xfd, 0x02, 0x70, 0x6b, 0xc0, 0xc3, 0x37, 0x69, 0x80, 0x05, 0x79, 0x5d, 0x78, 0xd0,
	0x3e, 0xd4, 0x70, 0x26, 0xce, 0x28, 0x8b, 0xc4, 0x50, 0x07, 0x1d, 0x60, 0x6b, 0x7d, 0xfd, 0xcf,
	0x8f, 0xc7, 0x3b, 0xaa, 0xdd, 0xf3, 0x20, 0x60, 0x84, 0xf3, 0x63, 0xc1, 0xa2, 0x24, 0xf4, 0xe6,
	0xa1, 0xe8, 0x19, 0x6c, 0xc8, 0xda, 0xfa, 0x46, 0x07, 0xd8, 0xcd, 0x5e, 0xdb, 0xb9, 0x31, 0xac,
	0x23, 0x5b, 0xf4, 0xb5, 0xcb, 0x7f, 0xf7, 0x6b, 0xdf, 0xa7, 0xa3, 0x2e, 0xf0, 0x54, 0xce, 0xc1,
	0xfe, 0x97, 0xe9, 0xa8, 0x3b, 0xaf, 0xf6, 0x75, 0x3a, 0xea, 0x3e, 0x9c, 0x69, 0xbe, 0x98, 0xcb,
	0xaf, 0xa8, 0xb5, 0xda, 0x70, 0xaf, 0x42, 0x79, 0x84, 0xa7, 0x34, 0xe1, 0xc4, 0xfa, 0x09, 0xe0,
	0xdd, 0x01, 0x0f, 0x0f, 0x93, 0x48, 0x1c, 0x0b, 0x9c, 0x67, 0x23, 0x1d, 0xde, 0xf6, 0x19, 0xc1,
	0x82, 0x32, 0x39, 0x99, 0x57, 0x42, 0xb4, 0x0b, 0x1b, 0x82, 0x61, 0xff, 0x53, 0xae, 0xbe, 0x6e,
	0x6b, 0x9e, 0x42, 0xc8, 0x86, 0x5b, 0xe7, 0x84, 0x45, 0xa7, 0x91, 0x5f, 0x54, 0x38, 0x22, 0x43,
	0xbd, 0xde, 0x01, 0x76, 0xcb, 0xab, 0xd2, 0xe8, 0x1e, 0xd4, 0xb8, 0x6c, 0x73, 0x18, 0xe8, 0xb7,
	0x8a, 0xea, 0x73, 0x02, 0x75, 0x60, 0xb3, 0x04, 0xc9, 0x29, 0xd5, 0x37, 0x0b, 0xff, 0x22, 0x75,
	0xd0, 0xca, 0x37, 0x50, 0xea, 0xb1, 0x5e, 0xc1, 0xdd, 0x65, 0xed, 0xe5, 0x58, 0xb9, 0xd2, 0x3c,
	0x2d, 0xe3, 0xc5, 0x08, 0x77, 0x3c, 0x85, 0x96, 0xfb, 0x6f, 0x54, 0xfa, 0xf7, 0x7e, 0x03, 0x58,
	0x1f, 0xf0, 0x10, 0xbd, 0x87, 0xad, 0xa5, 0xd3, 0xb6, 0x56, 0x9c, 0x52, 0x65, 0xa1, 0x46, 0x77,
	0x7d, 0xcc, 0x4c, 0xdd, 0x3b, 0xd8, 0x5c, 0x5c, 0xf8, 0x83, 0xd5, 0xa9, 0x0b, 0x21, 0xc6, 0xa3,
	0xb5, 0x21, 0x65, 0x71, 0x63, 0xf3, 0x73, 0x7e, 0x67, 0xfa, 0x2f, 0x2f, 0xc7, 0x26, 0xb8, 0x1a,
	0x9b, 0xe0, 0xff, 0xd8, 0x04, 0xdf, 0x26, 0x66, 0xed, 0x6a, 0x62, 0xd6, 0xfe, 0x4e, 0xcc, 0xda,
	0xdb, 0x5e, 0x18, 0x89, 0xb3, 0xec, 0xc4, 0xf1, 0x69, 0xec, 0xbe, 0xa0, 0x71, 0x9a, 0x09, 0xc2,
	0x8e, 0x08, 0x09, 0xb0, 0xbb, 0xea, 0x2e, 0x89, 0x61, 0x4a, 0xf8, 0x49, 0xa3, 0x78, 0x0a, 0x4f,
	0xaf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x3f, 0x85, 0x84, 0xec, 0xb8, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	InitStation(ctx context.Context, in *MsgInitStation, opts ...grpc.CallOption) (*MsgInitStationResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, "/junction.junction.Msg/UpdateParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) InitStation(ctx context.Context, in *MsgInitStation, opts ...grpc.CallOption) (*MsgInitStationResponse, error) {
	out := new(MsgInitStationResponse)
	err := c.cc.Invoke(ctx, "/junction.junction.Msg/InitStation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	InitStation(context.Context, *MsgInitStation) (*MsgInitStationResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) UpdateParams(ctx context.Context, req *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (*UnimplementedMsgServer) InitStation(ctx context.Context, req *MsgInitStation) (*MsgInitStationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitStation not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/junction.junction.Msg/UpdateParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_InitStation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgInitStation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).InitStation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/junction.junction.Msg/InitStation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).InitStation(ctx, req.(*MsgInitStation))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "junction.junction.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "InitStation",
			Handler:    _Msg_InitStation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "junction/junction/tx.proto",
}

func (m *MsgUpdateParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgInitStation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgInitStation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgInitStation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.StationInfo) > 0 {
		i -= len(m.StationInfo)
		copy(dAtA[i:], m.StationInfo)
		i = encodeVarintTx(dAtA, i, uint64(len(m.StationInfo)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.StationId) > 0 {
		i -= len(m.StationId)
		copy(dAtA[i:], m.StationId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.StationId)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.VerificationKey) > 0 {
		i -= len(m.VerificationKey)
		copy(dAtA[i:], m.VerificationKey)
		i = encodeVarintTx(dAtA, i, uint64(len(m.VerificationKey)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Tracks) > 0 {
		for iNdEx := len(m.Tracks) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Tracks[iNdEx])
			copy(dAtA[i:], m.Tracks[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.Tracks[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgInitStationResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgInitStationResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgInitStationResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.StationId) > 0 {
		i -= len(m.StationId)
		copy(dAtA[i:], m.StationId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.StationId)))
		i--
		dAtA[i] = 0x12
	}
	if m.Status {
		i--
		if m.Status {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgUpdateParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Params.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgUpdateParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgInitStation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Tracks) > 0 {
		for _, s := range m.Tracks {
			l = len(s)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	l = len(m.VerificationKey)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.StationId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.StationInfo)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgInitStationResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Status {
		n += 2
	}
	l = len(m.StationId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgUpdateParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUpdateParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgInitStation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgInitStation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgInitStation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tracks", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tracks = append(m.Tracks, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VerificationKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VerificationKey = append(m.VerificationKey[:0], dAtA[iNdEx:postIndex]...)
			if m.VerificationKey == nil {
				m.VerificationKey = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StationId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StationId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StationInfo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StationInfo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgInitStationResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgInitStationResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgInitStationResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Status = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StationId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StationId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
