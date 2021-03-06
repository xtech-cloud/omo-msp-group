// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/group/member.proto

package group

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Member service

type MemberService interface {
	// 加入一个成员
	Add(ctx context.Context, in *MemberAddRequest, opts ...client.CallOption) (*BlankResponse, error)
	// 列举一个集合中的成员
	List(ctx context.Context, in *MemberListRequest, opts ...client.CallOption) (*MemberListResponse, error)
	// 移除一个成员
	Remove(ctx context.Context, in *MemberRemoveRequest, opts ...client.CallOption) (*BlankResponse, error)
	// 获取一个成员信息
	Get(ctx context.Context, in *MemberGetRequest, opts ...client.CallOption) (*MemberGetResponse, error)
	// 获取所在的集合
	Where(ctx context.Context, in *MemberWhereRequest, opts ...client.CallOption) (*MemberWhereResponse, error)
}

type memberService struct {
	c    client.Client
	name string
}

func NewMemberService(name string, c client.Client) MemberService {
	return &memberService{
		c:    c,
		name: name,
	}
}

func (c *memberService) Add(ctx context.Context, in *MemberAddRequest, opts ...client.CallOption) (*BlankResponse, error) {
	req := c.c.NewRequest(c.name, "Member.Add", in)
	out := new(BlankResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberService) List(ctx context.Context, in *MemberListRequest, opts ...client.CallOption) (*MemberListResponse, error) {
	req := c.c.NewRequest(c.name, "Member.List", in)
	out := new(MemberListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberService) Remove(ctx context.Context, in *MemberRemoveRequest, opts ...client.CallOption) (*BlankResponse, error) {
	req := c.c.NewRequest(c.name, "Member.Remove", in)
	out := new(BlankResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberService) Get(ctx context.Context, in *MemberGetRequest, opts ...client.CallOption) (*MemberGetResponse, error) {
	req := c.c.NewRequest(c.name, "Member.Get", in)
	out := new(MemberGetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberService) Where(ctx context.Context, in *MemberWhereRequest, opts ...client.CallOption) (*MemberWhereResponse, error) {
	req := c.c.NewRequest(c.name, "Member.Where", in)
	out := new(MemberWhereResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Member service

type MemberHandler interface {
	// 加入一个成员
	Add(context.Context, *MemberAddRequest, *BlankResponse) error
	// 列举一个集合中的成员
	List(context.Context, *MemberListRequest, *MemberListResponse) error
	// 移除一个成员
	Remove(context.Context, *MemberRemoveRequest, *BlankResponse) error
	// 获取一个成员信息
	Get(context.Context, *MemberGetRequest, *MemberGetResponse) error
	// 获取所在的集合
	Where(context.Context, *MemberWhereRequest, *MemberWhereResponse) error
}

func RegisterMemberHandler(s server.Server, hdlr MemberHandler, opts ...server.HandlerOption) error {
	type member interface {
		Add(ctx context.Context, in *MemberAddRequest, out *BlankResponse) error
		List(ctx context.Context, in *MemberListRequest, out *MemberListResponse) error
		Remove(ctx context.Context, in *MemberRemoveRequest, out *BlankResponse) error
		Get(ctx context.Context, in *MemberGetRequest, out *MemberGetResponse) error
		Where(ctx context.Context, in *MemberWhereRequest, out *MemberWhereResponse) error
	}
	type Member struct {
		member
	}
	h := &memberHandler{hdlr}
	return s.Handle(s.NewHandler(&Member{h}, opts...))
}

type memberHandler struct {
	MemberHandler
}

func (h *memberHandler) Add(ctx context.Context, in *MemberAddRequest, out *BlankResponse) error {
	return h.MemberHandler.Add(ctx, in, out)
}

func (h *memberHandler) List(ctx context.Context, in *MemberListRequest, out *MemberListResponse) error {
	return h.MemberHandler.List(ctx, in, out)
}

func (h *memberHandler) Remove(ctx context.Context, in *MemberRemoveRequest, out *BlankResponse) error {
	return h.MemberHandler.Remove(ctx, in, out)
}

func (h *memberHandler) Get(ctx context.Context, in *MemberGetRequest, out *MemberGetResponse) error {
	return h.MemberHandler.Get(ctx, in, out)
}

func (h *memberHandler) Where(ctx context.Context, in *MemberWhereRequest, out *MemberWhereResponse) error {
	return h.MemberHandler.Where(ctx, in, out)
}
