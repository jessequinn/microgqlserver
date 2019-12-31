// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/proto/hello/hello.proto

package go_micro_srv_microgqlserver

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Request struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_dae8d3334d62a95f, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Response struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_dae8d3334d62a95f, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.srv.microgqlserver.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.microgqlserver.Response")
}

func init() { proto.RegisterFile("internal/proto/hello/hello.proto", fileDescriptor_dae8d3334d62a95f) }

var fileDescriptor_dae8d3334d62a95f = []byte{
	// 161 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xc8, 0xcc, 0x2b, 0x49,
	0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0xcf, 0x48, 0xcd, 0xc9, 0x81,
	0x92, 0x7a, 0x60, 0x11, 0x21, 0xe9, 0xf4, 0x7c, 0xbd, 0xdc, 0xcc, 0xe4, 0xa2, 0x7c, 0xbd, 0xe2,
	0xa2, 0x32, 0x08, 0x2b, 0xbd, 0x30, 0xa7, 0x38, 0xb5, 0xa8, 0x2c, 0xb5, 0x48, 0x49, 0x96, 0x8b,
	0x3d, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x48, 0x88, 0x8b, 0x25, 0x2f, 0x31, 0x37, 0x55,
	0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xcc, 0x56, 0x92, 0xe1, 0xe2, 0x08, 0x4a, 0x2d, 0x2e,
	0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x12, 0xe0, 0x62, 0xce, 0x2d, 0x4e, 0x87, 0x4a, 0x83, 0x98, 0x46,
	0xb1, 0x5c, 0xcc, 0xc1, 0x89, 0x95, 0x42, 0x61, 0x5c, 0xac, 0x1e, 0x20, 0xfb, 0x84, 0x54, 0xf4,
	0xf0, 0x58, 0xa5, 0x07, 0xb5, 0x47, 0x4a, 0x95, 0x80, 0x2a, 0x88, 0x75, 0x4a, 0x0c, 0x49, 0x6c,
	0x60, 0xf7, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x1b, 0xa4, 0x22, 0x2c, 0xe3, 0x00, 0x00,
	0x00,
}