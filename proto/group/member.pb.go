// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/group/member.proto

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

// 加入成员的请求
type MemberAddRequest struct {
	Collection           string   `protobuf:"bytes,1,opt,name=collection,proto3" json:"collection,omitempty"`
	Element              string   `protobuf:"bytes,2,opt,name=element,proto3" json:"element,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MemberAddRequest) Reset()         { *m = MemberAddRequest{} }
func (m *MemberAddRequest) String() string { return proto.CompactTextString(m) }
func (*MemberAddRequest) ProtoMessage()    {}
func (*MemberAddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_640ebdb6edb0325d, []int{0}
}

func (m *MemberAddRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemberAddRequest.Unmarshal(m, b)
}
func (m *MemberAddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemberAddRequest.Marshal(b, m, deterministic)
}
func (m *MemberAddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberAddRequest.Merge(m, src)
}
func (m *MemberAddRequest) XXX_Size() int {
	return xxx_messageInfo_MemberAddRequest.Size(m)
}
func (m *MemberAddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberAddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MemberAddRequest proto.InternalMessageInfo

func (m *MemberAddRequest) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}

func (m *MemberAddRequest) GetElement() string {
	if m != nil {
		return m.Element
	}
	return ""
}

// 列举成员的请求
type MemberListRequest struct {
	Offset               int64    `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Count                int64    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Collection           string   `protobuf:"bytes,3,opt,name=collection,proto3" json:"collection,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MemberListRequest) Reset()         { *m = MemberListRequest{} }
func (m *MemberListRequest) String() string { return proto.CompactTextString(m) }
func (*MemberListRequest) ProtoMessage()    {}
func (*MemberListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_640ebdb6edb0325d, []int{1}
}

func (m *MemberListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemberListRequest.Unmarshal(m, b)
}
func (m *MemberListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemberListRequest.Marshal(b, m, deterministic)
}
func (m *MemberListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberListRequest.Merge(m, src)
}
func (m *MemberListRequest) XXX_Size() int {
	return xxx_messageInfo_MemberListRequest.Size(m)
}
func (m *MemberListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MemberListRequest proto.InternalMessageInfo

func (m *MemberListRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *MemberListRequest) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *MemberListRequest) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}

// 列举成员的回复
type MemberListResponse struct {
	Status               *Status         `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Total                uint64          `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Entity               []*MemberEntity `protobuf:"bytes,3,rep,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *MemberListResponse) Reset()         { *m = MemberListResponse{} }
func (m *MemberListResponse) String() string { return proto.CompactTextString(m) }
func (*MemberListResponse) ProtoMessage()    {}
func (*MemberListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_640ebdb6edb0325d, []int{2}
}

func (m *MemberListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemberListResponse.Unmarshal(m, b)
}
func (m *MemberListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemberListResponse.Marshal(b, m, deterministic)
}
func (m *MemberListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberListResponse.Merge(m, src)
}
func (m *MemberListResponse) XXX_Size() int {
	return xxx_messageInfo_MemberListResponse.Size(m)
}
func (m *MemberListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MemberListResponse proto.InternalMessageInfo

func (m *MemberListResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *MemberListResponse) GetTotal() uint64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *MemberListResponse) GetEntity() []*MemberEntity {
	if m != nil {
		return m.Entity
	}
	return nil
}

// 移除成员的请求
type MemberRemoveRequest struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MemberRemoveRequest) Reset()         { *m = MemberRemoveRequest{} }
func (m *MemberRemoveRequest) String() string { return proto.CompactTextString(m) }
func (*MemberRemoveRequest) ProtoMessage()    {}
func (*MemberRemoveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_640ebdb6edb0325d, []int{3}
}

func (m *MemberRemoveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemberRemoveRequest.Unmarshal(m, b)
}
func (m *MemberRemoveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemberRemoveRequest.Marshal(b, m, deterministic)
}
func (m *MemberRemoveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberRemoveRequest.Merge(m, src)
}
func (m *MemberRemoveRequest) XXX_Size() int {
	return xxx_messageInfo_MemberRemoveRequest.Size(m)
}
func (m *MemberRemoveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberRemoveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MemberRemoveRequest proto.InternalMessageInfo

func (m *MemberRemoveRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

// 获取成员的请求
type MemberGetRequest struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MemberGetRequest) Reset()         { *m = MemberGetRequest{} }
func (m *MemberGetRequest) String() string { return proto.CompactTextString(m) }
func (*MemberGetRequest) ProtoMessage()    {}
func (*MemberGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_640ebdb6edb0325d, []int{4}
}

func (m *MemberGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemberGetRequest.Unmarshal(m, b)
}
func (m *MemberGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemberGetRequest.Marshal(b, m, deterministic)
}
func (m *MemberGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberGetRequest.Merge(m, src)
}
func (m *MemberGetRequest) XXX_Size() int {
	return xxx_messageInfo_MemberGetRequest.Size(m)
}
func (m *MemberGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MemberGetRequest proto.InternalMessageInfo

func (m *MemberGetRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

// 获取成员信息的回复
type MemberGetResponse struct {
	Status               *Status       `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Entity               *MemberEntity `protobuf:"bytes,2,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *MemberGetResponse) Reset()         { *m = MemberGetResponse{} }
func (m *MemberGetResponse) String() string { return proto.CompactTextString(m) }
func (*MemberGetResponse) ProtoMessage()    {}
func (*MemberGetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_640ebdb6edb0325d, []int{5}
}

func (m *MemberGetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemberGetResponse.Unmarshal(m, b)
}
func (m *MemberGetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemberGetResponse.Marshal(b, m, deterministic)
}
func (m *MemberGetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberGetResponse.Merge(m, src)
}
func (m *MemberGetResponse) XXX_Size() int {
	return xxx_messageInfo_MemberGetResponse.Size(m)
}
func (m *MemberGetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberGetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MemberGetResponse proto.InternalMessageInfo

func (m *MemberGetResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *MemberGetResponse) GetEntity() *MemberEntity {
	if m != nil {
		return m.Entity
	}
	return nil
}

// 获取成员所在的集合的请求
type MemberWhereRequest struct {
	Element              string   `protobuf:"bytes,1,opt,name=element,proto3" json:"element,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MemberWhereRequest) Reset()         { *m = MemberWhereRequest{} }
func (m *MemberWhereRequest) String() string { return proto.CompactTextString(m) }
func (*MemberWhereRequest) ProtoMessage()    {}
func (*MemberWhereRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_640ebdb6edb0325d, []int{6}
}

func (m *MemberWhereRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemberWhereRequest.Unmarshal(m, b)
}
func (m *MemberWhereRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemberWhereRequest.Marshal(b, m, deterministic)
}
func (m *MemberWhereRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberWhereRequest.Merge(m, src)
}
func (m *MemberWhereRequest) XXX_Size() int {
	return xxx_messageInfo_MemberWhereRequest.Size(m)
}
func (m *MemberWhereRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberWhereRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MemberWhereRequest proto.InternalMessageInfo

func (m *MemberWhereRequest) GetElement() string {
	if m != nil {
		return m.Element
	}
	return ""
}

// 获取成员所在的集合的回复
type MemberWhereResponse struct {
	Status               *Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Collection           []string `protobuf:"bytes,2,rep,name=collection,proto3" json:"collection,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MemberWhereResponse) Reset()         { *m = MemberWhereResponse{} }
func (m *MemberWhereResponse) String() string { return proto.CompactTextString(m) }
func (*MemberWhereResponse) ProtoMessage()    {}
func (*MemberWhereResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_640ebdb6edb0325d, []int{7}
}

func (m *MemberWhereResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemberWhereResponse.Unmarshal(m, b)
}
func (m *MemberWhereResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemberWhereResponse.Marshal(b, m, deterministic)
}
func (m *MemberWhereResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberWhereResponse.Merge(m, src)
}
func (m *MemberWhereResponse) XXX_Size() int {
	return xxx_messageInfo_MemberWhereResponse.Size(m)
}
func (m *MemberWhereResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberWhereResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MemberWhereResponse proto.InternalMessageInfo

func (m *MemberWhereResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *MemberWhereResponse) GetCollection() []string {
	if m != nil {
		return m.Collection
	}
	return nil
}

func init() {
	proto.RegisterType((*MemberAddRequest)(nil), "group.MemberAddRequest")
	proto.RegisterType((*MemberListRequest)(nil), "group.MemberListRequest")
	proto.RegisterType((*MemberListResponse)(nil), "group.MemberListResponse")
	proto.RegisterType((*MemberRemoveRequest)(nil), "group.MemberRemoveRequest")
	proto.RegisterType((*MemberGetRequest)(nil), "group.MemberGetRequest")
	proto.RegisterType((*MemberGetResponse)(nil), "group.MemberGetResponse")
	proto.RegisterType((*MemberWhereRequest)(nil), "group.MemberWhereRequest")
	proto.RegisterType((*MemberWhereResponse)(nil), "group.MemberWhereResponse")
}

func init() {
	proto.RegisterFile("proto/group/member.proto", fileDescriptor_640ebdb6edb0325d)
}

var fileDescriptor_640ebdb6edb0325d = []byte{
	// 412 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4d, 0x6b, 0xdb, 0x40,
	0x10, 0xad, 0xb5, 0xb6, 0x4a, 0xc7, 0x14, 0xda, 0xb5, 0x69, 0x55, 0x1d, 0x8a, 0x11, 0xb4, 0xb8,
	0x14, 0x64, 0x70, 0xa1, 0x87, 0xd2, 0x42, 0x5d, 0x28, 0xbe, 0x38, 0x97, 0xcd, 0x21, 0x97, 0x5c,
	0x64, 0x6b, 0x6c, 0x8b, 0x48, 0x5a, 0x47, 0xbb, 0x0a, 0xe4, 0x92, 0x3f, 0x96, 0x3f, 0x17, 0xbc,
	0x1f, 0xb6, 0xd6, 0x71, 0x0c, 0xbe, 0x69, 0x66, 0xde, 0xcc, 0x7b, 0x33, 0x6f, 0x05, 0xc1, 0xa6,
	0xe2, 0x92, 0x8f, 0x56, 0x15, 0xaf, 0x37, 0xa3, 0x02, 0x8b, 0x39, 0x56, 0xb1, 0x4a, 0xd1, 0x8e,
	0xca, 0x85, 0x0e, 0x40, 0xac, 0x93, 0x0a, 0x53, 0x0d, 0x88, 0x66, 0xf0, 0xee, 0x42, 0x35, 0x4c,
	0xd2, 0x94, 0xe1, 0x6d, 0x8d, 0x42, 0xd2, 0xcf, 0x00, 0x0b, 0x9e, 0xe7, 0xb8, 0x90, 0x19, 0x2f,
	0x83, 0xd6, 0xa0, 0x35, 0x7c, 0xc3, 0x1a, 0x19, 0x1a, 0xc0, 0x6b, 0xcc, 0xb1, 0xc0, 0x52, 0x06,
	0x9e, 0x2a, 0xda, 0x30, 0x4a, 0xe0, 0xbd, 0x9e, 0x36, 0xcb, 0x84, 0xb4, 0xe3, 0x3e, 0x80, 0xcf,
	0x97, 0x4b, 0x81, 0x52, 0x8d, 0x22, 0xcc, 0x44, 0xb4, 0x0f, 0x9d, 0x05, 0xaf, 0xcd, 0x10, 0xc2,
	0x74, 0x70, 0x40, 0x4e, 0x0e, 0xc9, 0xa3, 0x07, 0xa0, 0x4d, 0x0a, 0xb1, 0xe1, 0xa5, 0x40, 0xfa,
	0x05, 0x7c, 0x21, 0x13, 0x59, 0x0b, 0xc5, 0xd1, 0x1d, 0xbf, 0x8d, 0xd5, 0xae, 0xf1, 0xa5, 0x4a,
	0x32, 0x53, 0xdc, 0x52, 0x4a, 0x2e, 0x93, 0x5c, 0x51, 0xb6, 0x99, 0x0e, 0xe8, 0x77, 0xf0, 0xb1,
	0x94, 0x99, 0xbc, 0x0f, 0xc8, 0x80, 0x0c, 0xbb, 0xe3, 0x9e, 0x69, 0xd6, 0x3c, 0xff, 0x55, 0x89,
	0x19, 0x48, 0xf4, 0x0d, 0x7a, 0x3a, 0xcf, 0xb0, 0xe0, 0x77, 0x68, 0x97, 0xa4, 0xd0, 0xae, 0xeb,
	0x2c, 0x35, 0xd7, 0x52, 0xdf, 0xd1, 0x57, 0x7b, 0xdb, 0x29, 0xca, 0x53, 0xb8, 0x95, 0xbd, 0x9a,
	0xc2, 0x9d, 0xb7, 0xd1, 0x5e, 0xbb, 0xa7, 0x60, 0x27, 0xb5, 0xc7, 0xf6, 0x76, 0x57, 0x6b, 0xac,
	0x76, 0xd2, 0x1b, 0x76, 0xb6, 0x5c, 0x3b, 0xaf, 0xed, 0xae, 0x06, 0x7f, 0x9e, 0x34, 0xd7, 0x49,
	0x6f, 0x40, 0x5c, 0x27, 0xc7, 0x8f, 0x1e, 0xf8, 0x7a, 0x3c, 0xfd, 0x09, 0x64, 0x92, 0xa6, 0xf4,
	0xa3, 0x23, 0x7e, 0xff, 0x22, 0xc3, 0xbe, 0x29, 0xfc, 0xcb, 0x93, 0xf2, 0xc6, 0xea, 0x88, 0x5e,
	0xd1, 0x3f, 0xd0, 0xde, 0x3e, 0x03, 0x1a, 0x38, 0x8d, 0x8d, 0xc7, 0x17, 0x7e, 0x3a, 0x52, 0xd9,
	0xb5, 0xff, 0x06, 0x5f, 0xbb, 0x48, 0x43, 0x07, 0xe6, 0x58, 0xfb, 0x22, 0xf9, 0x2f, 0x20, 0x53,
	0x94, 0x07, 0xa2, 0xf7, 0x56, 0x87, 0xc1, 0xf3, 0xc2, 0xae, 0xf7, 0x2f, 0x74, 0xd4, 0x4d, 0xa9,
	0xab, 0xaf, 0xe9, 0x4b, 0x18, 0x1e, 0x2b, 0xd9, 0x09, 0x73, 0x5f, 0xfd, 0xbf, 0x3f, 0x9e, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xb0, 0x8f, 0xfc, 0x6a, 0xfc, 0x03, 0x00, 0x00,
}
