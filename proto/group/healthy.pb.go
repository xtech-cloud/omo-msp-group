// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/group/healthy.proto

package group

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

// 回显的请求
type EchoRequest struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoRequest) Reset()         { *m = EchoRequest{} }
func (m *EchoRequest) String() string { return proto.CompactTextString(m) }
func (*EchoRequest) ProtoMessage()    {}
func (*EchoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ba45f8dc1fddde1, []int{0}
}

func (m *EchoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoRequest.Unmarshal(m, b)
}
func (m *EchoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoRequest.Marshal(b, m, deterministic)
}
func (m *EchoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoRequest.Merge(m, src)
}
func (m *EchoRequest) XXX_Size() int {
	return xxx_messageInfo_EchoRequest.Size(m)
}
func (m *EchoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EchoRequest proto.InternalMessageInfo

func (m *EchoRequest) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

// 回显的回复
type EchoResponse struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoResponse) Reset()         { *m = EchoResponse{} }
func (m *EchoResponse) String() string { return proto.CompactTextString(m) }
func (*EchoResponse) ProtoMessage()    {}
func (*EchoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ba45f8dc1fddde1, []int{1}
}

func (m *EchoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoResponse.Unmarshal(m, b)
}
func (m *EchoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoResponse.Marshal(b, m, deterministic)
}
func (m *EchoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoResponse.Merge(m, src)
}
func (m *EchoResponse) XXX_Size() int {
	return xxx_messageInfo_EchoResponse.Size(m)
}
func (m *EchoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EchoResponse proto.InternalMessageInfo

func (m *EchoResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*EchoRequest)(nil), "group.EchoRequest")
	proto.RegisterType((*EchoResponse)(nil), "group.EchoResponse")
}

func init() {
	proto.RegisterFile("proto/group/healthy.proto", fileDescriptor_1ba45f8dc1fddde1)
}

var fileDescriptor_1ba45f8dc1fddde1 = []byte{
	// 132 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x2f, 0xca, 0x2f, 0x2d, 0xd0, 0xcf, 0x48, 0x4d, 0xcc, 0x29, 0xc9, 0xa8, 0xd4,
	0x03, 0x8b, 0x09, 0xb1, 0x82, 0x05, 0x95, 0xe4, 0xb9, 0xb8, 0x5d, 0x93, 0x33, 0xf2, 0x83, 0x52,
	0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x04, 0xb8, 0x98, 0x73, 0x8b, 0xd3, 0x25, 0x18, 0x15, 0x18,
	0x35, 0x38, 0x83, 0x40, 0x4c, 0x25, 0x05, 0x2e, 0x1e, 0x88, 0x82, 0xe2, 0x82, 0xfc, 0xbc, 0xe2,
	0x54, 0x4c, 0x15, 0x46, 0x36, 0x5c, 0xec, 0x1e, 0x10, 0xa3, 0x85, 0x0c, 0xb9, 0x58, 0x40, 0x8a,
	0x85, 0x84, 0xf4, 0xc0, 0xa6, 0xeb, 0x21, 0x19, 0x2d, 0x25, 0x8c, 0x22, 0x06, 0x31, 0x4d, 0x89,
	0x21, 0x89, 0x0d, 0xec, 0x1c, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x20, 0xa7, 0x7e, 0xa7,
	0xab, 0x00, 0x00, 0x00,
}
