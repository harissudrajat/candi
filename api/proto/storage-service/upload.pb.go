// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/proto/storage-service/upload.proto

package upload

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UploadStatusCode int32

const (
	UploadStatusCode_Unknown UploadStatusCode = 0
	UploadStatusCode_Ok      UploadStatusCode = 1
	UploadStatusCode_Failed  UploadStatusCode = 2
)

var UploadStatusCode_name = map[int32]string{
	0: "Unknown",
	1: "Ok",
	2: "Failed",
}
var UploadStatusCode_value = map[string]int32{
	"Unknown": 0,
	"Ok":      1,
	"Failed":  2,
}

func (x UploadStatusCode) String() string {
	return proto.EnumName(UploadStatusCode_name, int32(x))
}
func (UploadStatusCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_upload_650827a93b8dd2e7, []int{0}
}

type Chunk struct {
	Content              []byte   `protobuf:"bytes,1,opt,name=Content,proto3" json:"Content,omitempty"`
	TotalSize            int64    `protobuf:"varint,2,opt,name=TotalSize,proto3" json:"TotalSize,omitempty"`
	Received             int64    `protobuf:"varint,3,opt,name=Received,proto3" json:"Received,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Chunk) Reset()         { *m = Chunk{} }
func (m *Chunk) String() string { return proto.CompactTextString(m) }
func (*Chunk) ProtoMessage()    {}
func (*Chunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_upload_650827a93b8dd2e7, []int{0}
}
func (m *Chunk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Chunk.Unmarshal(m, b)
}
func (m *Chunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Chunk.Marshal(b, m, deterministic)
}
func (dst *Chunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chunk.Merge(dst, src)
}
func (m *Chunk) XXX_Size() int {
	return xxx_messageInfo_Chunk.Size(m)
}
func (m *Chunk) XXX_DiscardUnknown() {
	xxx_messageInfo_Chunk.DiscardUnknown(m)
}

var xxx_messageInfo_Chunk proto.InternalMessageInfo

func (m *Chunk) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *Chunk) GetTotalSize() int64 {
	if m != nil {
		return m.TotalSize
	}
	return 0
}

func (m *Chunk) GetReceived() int64 {
	if m != nil {
		return m.Received
	}
	return 0
}

type UploadStatus struct {
	Message              string           `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	File                 string           `protobuf:"bytes,2,opt,name=File,proto3" json:"File,omitempty"`
	Size                 int64            `protobuf:"varint,3,opt,name=Size,proto3" json:"Size,omitempty"`
	Code                 UploadStatusCode `protobuf:"varint,4,opt,name=Code,proto3,enum=upload.UploadStatusCode" json:"Code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *UploadStatus) Reset()         { *m = UploadStatus{} }
func (m *UploadStatus) String() string { return proto.CompactTextString(m) }
func (*UploadStatus) ProtoMessage()    {}
func (*UploadStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_upload_650827a93b8dd2e7, []int{1}
}
func (m *UploadStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadStatus.Unmarshal(m, b)
}
func (m *UploadStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadStatus.Marshal(b, m, deterministic)
}
func (dst *UploadStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadStatus.Merge(dst, src)
}
func (m *UploadStatus) XXX_Size() int {
	return xxx_messageInfo_UploadStatus.Size(m)
}
func (m *UploadStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadStatus.DiscardUnknown(m)
}

var xxx_messageInfo_UploadStatus proto.InternalMessageInfo

func (m *UploadStatus) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *UploadStatus) GetFile() string {
	if m != nil {
		return m.File
	}
	return ""
}

func (m *UploadStatus) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *UploadStatus) GetCode() UploadStatusCode {
	if m != nil {
		return m.Code
	}
	return UploadStatusCode_Unknown
}

func init() {
	proto.RegisterType((*Chunk)(nil), "upload.Chunk")
	proto.RegisterType((*UploadStatus)(nil), "upload.UploadStatus")
	proto.RegisterEnum("upload.UploadStatusCode", UploadStatusCode_name, UploadStatusCode_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UploadServiceClient is the client API for UploadService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UploadServiceClient interface {
	Upload(ctx context.Context, opts ...grpc.CallOption) (UploadService_UploadClient, error)
}

type uploadServiceClient struct {
	cc *grpc.ClientConn
}

func NewUploadServiceClient(cc *grpc.ClientConn) UploadServiceClient {
	return &uploadServiceClient{cc}
}

func (c *uploadServiceClient) Upload(ctx context.Context, opts ...grpc.CallOption) (UploadService_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_UploadService_serviceDesc.Streams[0], "/upload.UploadService/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &uploadServiceUploadClient{stream}
	return x, nil
}

type UploadService_UploadClient interface {
	Send(*Chunk) error
	CloseAndRecv() (*UploadStatus, error)
	grpc.ClientStream
}

type uploadServiceUploadClient struct {
	grpc.ClientStream
}

func (x *uploadServiceUploadClient) Send(m *Chunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *uploadServiceUploadClient) CloseAndRecv() (*UploadStatus, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UploadServiceServer is the server API for UploadService service.
type UploadServiceServer interface {
	Upload(UploadService_UploadServer) error
}

func RegisterUploadServiceServer(s *grpc.Server, srv UploadServiceServer) {
	s.RegisterService(&_UploadService_serviceDesc, srv)
}

func _UploadService_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UploadServiceServer).Upload(&uploadServiceUploadServer{stream})
}

type UploadService_UploadServer interface {
	SendAndClose(*UploadStatus) error
	Recv() (*Chunk, error)
	grpc.ServerStream
}

type uploadServiceUploadServer struct {
	grpc.ServerStream
}

func (x *uploadServiceUploadServer) SendAndClose(m *UploadStatus) error {
	return x.ServerStream.SendMsg(m)
}

func (x *uploadServiceUploadServer) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _UploadService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "upload.UploadService",
	HandlerType: (*UploadServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _UploadService_Upload_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "api/proto/storage-service/upload.proto",
}

func init() {
	proto.RegisterFile("api/proto/storage-service/upload.proto", fileDescriptor_upload_650827a93b8dd2e7)
}

var fileDescriptor_upload_650827a93b8dd2e7 = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4b, 0xfb, 0x40,
	0x10, 0xc5, 0xbb, 0x6d, 0xfe, 0xe9, 0x3f, 0x63, 0x2b, 0x61, 0xf0, 0xb0, 0x14, 0x0f, 0x21, 0x07,
	0x09, 0xa2, 0x0d, 0xb6, 0xdf, 0xc0, 0x40, 0x6f, 0x22, 0x6c, 0xed, 0xc9, 0xd3, 0xda, 0x0c, 0x75,
	0x49, 0xd8, 0x0d, 0xc9, 0xa6, 0x82, 0x07, 0x3f, 0xbb, 0x64, 0xd3, 0xa8, 0x88, 0xb7, 0x79, 0xef,
	0xed, 0xce, 0x8f, 0x79, 0x70, 0x25, 0x2b, 0x95, 0x56, 0xb5, 0xb1, 0x26, 0x6d, 0xac, 0xa9, 0xe5,
	0x81, 0x6e, 0x1b, 0xaa, 0x8f, 0x6a, 0x4f, 0x69, 0x5b, 0x95, 0x46, 0xe6, 0x4b, 0x17, 0xa2, 0xdf,
	0xab, 0xf8, 0x19, 0xfe, 0x65, 0xaf, 0xad, 0x2e, 0x90, 0xc3, 0x34, 0x33, 0xda, 0x92, 0xb6, 0x9c,
	0x45, 0x2c, 0x99, 0x89, 0x41, 0xe2, 0x25, 0x04, 0x4f, 0xc6, 0xca, 0x72, 0xab, 0xde, 0x89, 0x8f,
	0x23, 0x96, 0x4c, 0xc4, 0xb7, 0x81, 0x0b, 0xf8, 0x2f, 0x68, 0x4f, 0xea, 0x48, 0x39, 0x9f, 0xb8,
	0xf0, 0x4b, 0xc7, 0x1f, 0x30, 0xdb, 0x39, 0xcc, 0xd6, 0x4a, 0xdb, 0x36, 0x1d, 0xe3, 0x81, 0x9a,
	0x46, 0x1e, 0xc8, 0x31, 0x02, 0x31, 0x48, 0x44, 0xf0, 0x36, 0xaa, 0xec, 0xd7, 0x07, 0xc2, 0xcd,
	0x9d, 0xe7, 0x90, 0xfd, 0x56, 0x37, 0xe3, 0x0d, 0x78, 0x99, 0xc9, 0x89, 0x7b, 0x11, 0x4b, 0xce,
	0x57, 0x7c, 0x79, 0xba, 0xe9, 0x27, 0xa5, 0xcb, 0x85, 0x7b, 0x75, 0xbd, 0x86, 0xf0, 0x77, 0x82,
	0x67, 0x30, 0xdd, 0xe9, 0x42, 0x9b, 0x37, 0x1d, 0x8e, 0xd0, 0x87, 0xf1, 0x63, 0x11, 0x32, 0x04,
	0xf0, 0x37, 0x52, 0x95, 0x94, 0x87, 0xe3, 0xd5, 0x3d, 0xcc, 0x4f, 0x9f, 0xfa, 0xde, 0xf0, 0x0e,
	0xfc, 0xde, 0xc0, 0xf9, 0xc0, 0x73, 0x95, 0x2d, 0x2e, 0xfe, 0xc2, 0xc7, 0xa3, 0x84, 0xbd, 0xf8,
	0xae, 0xe4, 0xf5, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x10, 0xae, 0xd4, 0x8e, 0x8e, 0x01, 0x00,
	0x00,
}