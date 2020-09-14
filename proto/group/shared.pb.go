// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/group/shared.proto

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

// 状态
type Status struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e698bc990f59333, []int{0}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Status) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// 空白请求
type BlankRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlankRequest) Reset()         { *m = BlankRequest{} }
func (m *BlankRequest) String() string { return proto.CompactTextString(m) }
func (*BlankRequest) ProtoMessage()    {}
func (*BlankRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e698bc990f59333, []int{1}
}

func (m *BlankRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlankRequest.Unmarshal(m, b)
}
func (m *BlankRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlankRequest.Marshal(b, m, deterministic)
}
func (m *BlankRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlankRequest.Merge(m, src)
}
func (m *BlankRequest) XXX_Size() int {
	return xxx_messageInfo_BlankRequest.Size(m)
}
func (m *BlankRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BlankRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BlankRequest proto.InternalMessageInfo

// 空白回复
type BlankResponse struct {
	Status               *Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlankResponse) Reset()         { *m = BlankResponse{} }
func (m *BlankResponse) String() string { return proto.CompactTextString(m) }
func (*BlankResponse) ProtoMessage()    {}
func (*BlankResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e698bc990f59333, []int{2}
}

func (m *BlankResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlankResponse.Unmarshal(m, b)
}
func (m *BlankResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlankResponse.Marshal(b, m, deterministic)
}
func (m *BlankResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlankResponse.Merge(m, src)
}
func (m *BlankResponse) XXX_Size() int {
	return xxx_messageInfo_BlankResponse.Size(m)
}
func (m *BlankResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BlankResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BlankResponse proto.InternalMessageInfo

func (m *BlankResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

// 集合实体
type CollectionEntity struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Capacity             uint64   `protobuf:"varint,3,opt,name=capacity,proto3" json:"capacity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CollectionEntity) Reset()         { *m = CollectionEntity{} }
func (m *CollectionEntity) String() string { return proto.CompactTextString(m) }
func (*CollectionEntity) ProtoMessage()    {}
func (*CollectionEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e698bc990f59333, []int{3}
}

func (m *CollectionEntity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectionEntity.Unmarshal(m, b)
}
func (m *CollectionEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectionEntity.Marshal(b, m, deterministic)
}
func (m *CollectionEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectionEntity.Merge(m, src)
}
func (m *CollectionEntity) XXX_Size() int {
	return xxx_messageInfo_CollectionEntity.Size(m)
}
func (m *CollectionEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectionEntity.DiscardUnknown(m)
}

var xxx_messageInfo_CollectionEntity proto.InternalMessageInfo

func (m *CollectionEntity) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *CollectionEntity) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CollectionEntity) GetCapacity() uint64 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

// 成员实体
type MemberEntity struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MemberEntity) Reset()         { *m = MemberEntity{} }
func (m *MemberEntity) String() string { return proto.CompactTextString(m) }
func (*MemberEntity) ProtoMessage()    {}
func (*MemberEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e698bc990f59333, []int{4}
}

func (m *MemberEntity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemberEntity.Unmarshal(m, b)
}
func (m *MemberEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemberEntity.Marshal(b, m, deterministic)
}
func (m *MemberEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberEntity.Merge(m, src)
}
func (m *MemberEntity) XXX_Size() int {
	return xxx_messageInfo_MemberEntity.Size(m)
}
func (m *MemberEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberEntity.DiscardUnknown(m)
}

var xxx_messageInfo_MemberEntity proto.InternalMessageInfo

func (m *MemberEntity) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func init() {
	proto.RegisterType((*Status)(nil), "group.Status")
	proto.RegisterType((*BlankRequest)(nil), "group.BlankRequest")
	proto.RegisterType((*BlankResponse)(nil), "group.BlankResponse")
	proto.RegisterType((*CollectionEntity)(nil), "group.CollectionEntity")
	proto.RegisterType((*MemberEntity)(nil), "group.MemberEntity")
}

func init() {
	proto.RegisterFile("proto/group/shared.proto", fileDescriptor_7e698bc990f59333)
}

var fileDescriptor_7e698bc990f59333 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0x59, 0x6d, 0x57, 0x3b, 0xb6, 0x22, 0x39, 0x05, 0x4f, 0x4b, 0x40, 0xd8, 0xd3, 0x16,
	0x14, 0xfa, 0x03, 0x14, 0x8f, 0x5e, 0x22, 0x78, 0x4f, 0x77, 0x87, 0x1a, 0xdc, 0x4d, 0x62, 0x66,
	0x72, 0xe8, 0xbf, 0x97, 0x8e, 0x5d, 0x6f, 0xde, 0xde, 0x7b, 0x21, 0xdf, 0x7b, 0x0c, 0xe8, 0x94,
	0x23, 0xc7, 0xed, 0x21, 0xc7, 0x92, 0xb6, 0xf4, 0xe9, 0x32, 0x0e, 0x9d, 0x44, 0x6a, 0x29, 0x99,
	0xd9, 0x41, 0xfd, 0xce, 0x8e, 0x0b, 0x29, 0x05, 0x8b, 0x3e, 0x0e, 0xa8, 0xab, 0xa6, 0x6a, 0x97,
	0x56, 0xb4, 0xd2, 0x70, 0x35, 0x21, 0x91, 0x3b, 0xa0, 0xbe, 0x68, 0xaa, 0x76, 0x65, 0x67, 0x6b,
	0x6e, 0x61, 0xfd, 0x3c, 0xba, 0xf0, 0x65, 0xf1, 0xbb, 0x20, 0xb1, 0xd9, 0xc1, 0xe6, 0xec, 0x29,
	0xc5, 0x40, 0xa8, 0x1e, 0xa0, 0x26, 0x01, 0x0b, 0xf0, 0xe6, 0x71, 0xd3, 0x49, 0x61, 0xf7, 0xdb,
	0x66, 0xcf, 0x8f, 0xe6, 0x03, 0xee, 0x5e, 0xe2, 0x38, 0x62, 0xcf, 0x3e, 0x86, 0xd7, 0xc0, 0x9e,
	0x8f, 0xa7, 0x25, 0xa5, 0xf8, 0x41, 0x3e, 0xae, 0xac, 0xe8, 0x53, 0x16, 0xdc, 0x34, 0xcf, 0x10,
	0xad, 0xee, 0xe1, 0xba, 0x77, 0xc9, 0xf5, 0x9e, 0x8f, 0xfa, 0xb2, 0xa9, 0xda, 0x85, 0xfd, 0xf3,
	0xc6, 0xc0, 0xfa, 0x0d, 0xa7, 0x3d, 0xe6, 0xff, 0x99, 0xfb, 0x5a, 0x2e, 0xf1, 0xf4, 0x13, 0x00,
	0x00, 0xff, 0xff, 0xd6, 0x67, 0x23, 0xb0, 0x25, 0x01, 0x00, 0x00,
}