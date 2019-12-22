// Code generated by protoc-gen-go. DO NOT EDIT.
// source: attendance-service.proto

package attendance_service

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type AttendanceRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Lat                  float64  `protobuf:"fixed64,2,opt,name=lat,proto3" json:"lat,omitempty"`
	Long                 float64  `protobuf:"fixed64,3,opt,name=long,proto3" json:"long,omitempty"`
	UserId               string   `protobuf:"bytes,4,opt,name=userId,proto3" json:"userId,omitempty"`
	Description          string   `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	MemberName           string   `protobuf:"bytes,6,opt,name=memberName,proto3" json:"memberName,omitempty"`
	ImageUrl             string   `protobuf:"bytes,7,opt,name=imageUrl,proto3" json:"imageUrl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AttendanceRequest) Reset()         { *m = AttendanceRequest{} }
func (m *AttendanceRequest) String() string { return proto.CompactTextString(m) }
func (*AttendanceRequest) ProtoMessage()    {}
func (*AttendanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4d460dc24252017, []int{0}
}

func (m *AttendanceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttendanceRequest.Unmarshal(m, b)
}
func (m *AttendanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttendanceRequest.Marshal(b, m, deterministic)
}
func (m *AttendanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttendanceRequest.Merge(m, src)
}
func (m *AttendanceRequest) XXX_Size() int {
	return xxx_messageInfo_AttendanceRequest.Size(m)
}
func (m *AttendanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AttendanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AttendanceRequest proto.InternalMessageInfo

func (m *AttendanceRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *AttendanceRequest) GetLat() float64 {
	if m != nil {
		return m.Lat
	}
	return 0
}

func (m *AttendanceRequest) GetLong() float64 {
	if m != nil {
		return m.Long
	}
	return 0
}

func (m *AttendanceRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *AttendanceRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *AttendanceRequest) GetMemberName() string {
	if m != nil {
		return m.MemberName
	}
	return ""
}

func (m *AttendanceRequest) GetImageUrl() string {
	if m != nil {
		return m.ImageUrl
	}
	return ""
}

type AttendanceResponse struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Error                bool     `protobuf:"varint,3,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AttendanceResponse) Reset()         { *m = AttendanceResponse{} }
func (m *AttendanceResponse) String() string { return proto.CompactTextString(m) }
func (*AttendanceResponse) ProtoMessage()    {}
func (*AttendanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4d460dc24252017, []int{1}
}

func (m *AttendanceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttendanceResponse.Unmarshal(m, b)
}
func (m *AttendanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttendanceResponse.Marshal(b, m, deterministic)
}
func (m *AttendanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttendanceResponse.Merge(m, src)
}
func (m *AttendanceResponse) XXX_Size() int {
	return xxx_messageInfo_AttendanceResponse.Size(m)
}
func (m *AttendanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AttendanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AttendanceResponse proto.InternalMessageInfo

func (m *AttendanceResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *AttendanceResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *AttendanceResponse) GetError() bool {
	if m != nil {
		return m.Error
	}
	return false
}

func init() {
	proto.RegisterType((*AttendanceRequest)(nil), "AttendanceRequest")
	proto.RegisterType((*AttendanceResponse)(nil), "AttendanceResponse")
}

func init() { proto.RegisterFile("attendance-service.proto", fileDescriptor_e4d460dc24252017) }

var fileDescriptor_e4d460dc24252017 = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0x4b, 0x4e, 0xc3, 0x30,
	0x10, 0x95, 0xfb, 0xef, 0xb0, 0x29, 0xc3, 0x47, 0x56, 0x84, 0x50, 0x94, 0x55, 0x85, 0x44, 0x23,
	0x60, 0xc7, 0x0e, 0xb1, 0x62, 0xc3, 0x22, 0x08, 0x76, 0x2c, 0xdc, 0x74, 0x14, 0x59, 0x4a, 0xec,
	0x30, 0x76, 0x7b, 0x00, 0xae, 0xc0, 0x89, 0x38, 0x03, 0x57, 0xe0, 0x20, 0x28, 0x0e, 0xb4, 0x91,
	0xca, 0xee, 0x7d, 0x3c, 0xb6, 0xdf, 0x1b, 0x90, 0xca, 0x7b, 0x32, 0x2b, 0x65, 0x72, 0xba, 0x74,
	0xc4, 0x1b, 0x9d, 0xd3, 0xa2, 0x66, 0xeb, 0x6d, 0x74, 0x56, 0x58, 0x5b, 0x94, 0x94, 0xaa, 0x5a,
	0xa7, 0xca, 0x18, 0xeb, 0x95, 0xd7, 0xd6, 0xb8, 0xd6, 0x4d, 0x3e, 0x05, 0x1c, 0xde, 0x6d, 0x47,
	0x33, 0x7a, 0x5b, 0x93, 0xf3, 0x38, 0x83, 0xbe, 0xaa, 0xb5, 0x14, 0xb1, 0x98, 0x4f, 0xb3, 0x06,
	0x36, 0x4a, 0xa9, 0xbc, 0xec, 0xc5, 0x62, 0x2e, 0xb2, 0x06, 0x22, 0xc2, 0xa0, 0xb4, 0xa6, 0x90,
	0xfd, 0x20, 0x05, 0x8c, 0xa7, 0x30, 0x5a, 0x3b, 0xe2, 0x87, 0x95, 0x1c, 0x84, 0xd1, 0x5f, 0x86,
	0x31, 0x1c, 0xac, 0xc8, 0xe5, 0xac, 0xeb, 0xe6, 0x6d, 0x39, 0x0c, 0x66, 0x57, 0xc2, 0x73, 0x80,
	0x8a, 0xaa, 0x25, 0xf1, 0xa3, 0xaa, 0x48, 0x8e, 0xc2, 0x81, 0x8e, 0x82, 0x11, 0x4c, 0x74, 0xa5,
	0x0a, 0x7a, 0xe6, 0x52, 0x8e, 0x83, 0xbb, 0xe5, 0xc9, 0x0b, 0x60, 0x37, 0x82, 0xab, 0xad, 0x71,
	0xf4, 0x4f, 0x06, 0x09, 0xe3, 0x8a, 0x9c, 0x53, 0x05, 0x85, 0x1c, 0xd3, 0xec, 0x8f, 0xe2, 0x31,
	0x0c, 0x89, 0xd9, 0x72, 0x08, 0x33, 0xc9, 0x5a, 0x72, 0xcd, 0xdd, 0x6a, 0x9e, 0xda, 0x52, 0xf1,
	0x15, 0x66, 0xf7, 0x4c, 0xca, 0xd3, 0xce, 0x42, 0x5c, 0xec, 0x55, 0x18, 0x1d, 0x2d, 0xf6, 0xff,
	0x94, 0xc4, 0xef, 0x5f, 0xdf, 0x1f, 0xbd, 0x28, 0x39, 0x49, 0x37, 0x57, 0xe9, 0x6e, 0x63, 0x69,
	0x1e, 0x6e, 0xbc, 0x15, 0x17, 0xcb, 0x51, 0x58, 0xcb, 0xcd, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x11, 0x3b, 0xce, 0xa2, 0xd0, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AttendanceServiceClient is the client API for AttendanceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AttendanceServiceClient interface {
	CreateAttendance(ctx context.Context, in *AttendanceRequest, opts ...grpc.CallOption) (*AttendanceResponse, error)
}

type attendanceServiceClient struct {
	cc *grpc.ClientConn
}

func NewAttendanceServiceClient(cc *grpc.ClientConn) AttendanceServiceClient {
	return &attendanceServiceClient{cc}
}

func (c *attendanceServiceClient) CreateAttendance(ctx context.Context, in *AttendanceRequest, opts ...grpc.CallOption) (*AttendanceResponse, error) {
	out := new(AttendanceResponse)
	err := c.cc.Invoke(ctx, "/AttendanceService/CreateAttendance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AttendanceServiceServer is the server API for AttendanceService service.
type AttendanceServiceServer interface {
	CreateAttendance(context.Context, *AttendanceRequest) (*AttendanceResponse, error)
}

// UnimplementedAttendanceServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAttendanceServiceServer struct {
}

func (*UnimplementedAttendanceServiceServer) CreateAttendance(ctx context.Context, req *AttendanceRequest) (*AttendanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAttendance not implemented")
}

func RegisterAttendanceServiceServer(s *grpc.Server, srv AttendanceServiceServer) {
	s.RegisterService(&_AttendanceService_serviceDesc, srv)
}

func _AttendanceService_CreateAttendance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttendanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttendanceServiceServer).CreateAttendance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AttendanceService/CreateAttendance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttendanceServiceServer).CreateAttendance(ctx, req.(*AttendanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AttendanceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "AttendanceService",
	HandlerType: (*AttendanceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAttendance",
			Handler:    _AttendanceService_CreateAttendance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "attendance-service.proto",
}
