// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: regen/divvy/v1/query.proto

package divvy

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	grpc "google.golang.org/grpc"
	math "math"
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

func init() { proto.RegisterFile("regen/divvy/v1/query.proto", fileDescriptor_bde75e31310101ac) }

var fileDescriptor_bde75e31310101ac = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8e, 0x31, 0x4e, 0xc4, 0x30,
	0x10, 0x45, 0x93, 0x02, 0x90, 0xb6, 0xa0, 0x40, 0x54, 0x29, 0x86, 0x13, 0x6c, 0x46, 0x81, 0x1b,
	0xd0, 0x50, 0xd3, 0xd2, 0xd9, 0xc9, 0x60, 0x2c, 0xd6, 0x99, 0x60, 0x8f, 0x0d, 0xb9, 0x05, 0xc7,
	0xa2, 0xdc, 0x92, 0x12, 0x25, 0x17, 0x41, 0xeb, 0x64, 0x9b, 0x74, 0xff, 0xfb, 0x7d, 0xff, 0xf9,
	0xbb, 0xca, 0x93, 0xa1, 0x1e, 0x3b, 0x9b, 0xd2, 0x88, 0xa9, 0xc1, 0x8f, 0x48, 0x7e, 0xac, 0x07,
	0xcf, 0xc2, 0x37, 0xd7, 0x99, 0xd5, 0x99, 0xd5, 0xa9, 0xa9, 0xa0, 0xe5, 0xe0, 0x38, 0xa0, 0x56,
	0x81, 0x30, 0x35, 0x9a, 0x44, 0x35, 0xd8, 0xb2, 0xed, 0x97, 0x7c, 0x75, 0x6b, 0xd8, 0x70, 0x96,
	0x78, 0x52, 0xeb, 0x2b, 0x18, 0x66, 0x73, 0x20, 0xcc, 0x4e, 0xc7, 0x57, 0xec, 0xa2, 0x57, 0x62,
	0xf9, 0xfc, 0xeb, 0x6e, 0xcb, 0xc5, 0x3a, 0x0a, 0xa2, 0xdc, 0xb0, 0x06, 0xb6, 0x13, 0x65, 0x1c,
	0x28, 0x2c, 0xec, 0xfe, 0x6a, 0x77, 0xf1, 0x7c, 0x5a, 0xfc, 0xf8, 0xf4, 0x33, 0x41, 0x79, 0x9c,
	0xa0, 0xfc, 0x9b, 0xa0, 0xfc, 0x9e, 0xa1, 0x38, 0xce, 0x50, 0xfc, 0xce, 0x50, 0xbc, 0xec, 0x8d,
	0x95, 0xb7, 0xa8, 0xeb, 0x96, 0x1d, 0xe6, 0xa6, 0x7d, 0x4f, 0xf2, 0xc9, 0xfe, 0x7d, 0x75, 0x07,
	0xea, 0x0c, 0x79, 0xfc, 0x5a, 0x0e, 0xe8, 0xcb, 0x5c, 0xfc, 0xf0, 0x1f, 0x00, 0x00, 0xff, 0xff,
	0x41, 0xf1, 0x23, 0x13, 0x19, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

// QueryServer is the server API for Query service.
type QueryServer interface {
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "regen.divvy.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "regen/divvy/v1/query.proto",
}
