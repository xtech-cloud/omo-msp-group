// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/group/collection.proto

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

// Client API for Collection service

type CollectionService interface {
	// 创建一个集合
	Make(ctx context.Context, in *CollectionMakeRequest, opts ...client.CallOption) (*BlankResponse, error)
	// 列举集合
	List(ctx context.Context, in *CollectionListRequest, opts ...client.CallOption) (*CollectionListResponse, error)
	// 删除一个集合
	Remove(ctx context.Context, in *CollectionRemoveRequest, opts ...client.CallOption) (*BlankResponse, error)
	// 获取一个集合信息
	Get(ctx context.Context, in *CollectionGetRequest, opts ...client.CallOption) (*CollectionGetResponse, error)
}

type collectionService struct {
	c    client.Client
	name string
}

func NewCollectionService(name string, c client.Client) CollectionService {
	return &collectionService{
		c:    c,
		name: name,
	}
}

func (c *collectionService) Make(ctx context.Context, in *CollectionMakeRequest, opts ...client.CallOption) (*BlankResponse, error) {
	req := c.c.NewRequest(c.name, "Collection.Make", in)
	out := new(BlankResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionService) List(ctx context.Context, in *CollectionListRequest, opts ...client.CallOption) (*CollectionListResponse, error) {
	req := c.c.NewRequest(c.name, "Collection.List", in)
	out := new(CollectionListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionService) Remove(ctx context.Context, in *CollectionRemoveRequest, opts ...client.CallOption) (*BlankResponse, error) {
	req := c.c.NewRequest(c.name, "Collection.Remove", in)
	out := new(BlankResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionService) Get(ctx context.Context, in *CollectionGetRequest, opts ...client.CallOption) (*CollectionGetResponse, error) {
	req := c.c.NewRequest(c.name, "Collection.Get", in)
	out := new(CollectionGetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Collection service

type CollectionHandler interface {
	// 创建一个集合
	Make(context.Context, *CollectionMakeRequest, *BlankResponse) error
	// 列举集合
	List(context.Context, *CollectionListRequest, *CollectionListResponse) error
	// 删除一个集合
	Remove(context.Context, *CollectionRemoveRequest, *BlankResponse) error
	// 获取一个集合信息
	Get(context.Context, *CollectionGetRequest, *CollectionGetResponse) error
}

func RegisterCollectionHandler(s server.Server, hdlr CollectionHandler, opts ...server.HandlerOption) error {
	type collection interface {
		Make(ctx context.Context, in *CollectionMakeRequest, out *BlankResponse) error
		List(ctx context.Context, in *CollectionListRequest, out *CollectionListResponse) error
		Remove(ctx context.Context, in *CollectionRemoveRequest, out *BlankResponse) error
		Get(ctx context.Context, in *CollectionGetRequest, out *CollectionGetResponse) error
	}
	type Collection struct {
		collection
	}
	h := &collectionHandler{hdlr}
	return s.Handle(s.NewHandler(&Collection{h}, opts...))
}

type collectionHandler struct {
	CollectionHandler
}

func (h *collectionHandler) Make(ctx context.Context, in *CollectionMakeRequest, out *BlankResponse) error {
	return h.CollectionHandler.Make(ctx, in, out)
}

func (h *collectionHandler) List(ctx context.Context, in *CollectionListRequest, out *CollectionListResponse) error {
	return h.CollectionHandler.List(ctx, in, out)
}

func (h *collectionHandler) Remove(ctx context.Context, in *CollectionRemoveRequest, out *BlankResponse) error {
	return h.CollectionHandler.Remove(ctx, in, out)
}

func (h *collectionHandler) Get(ctx context.Context, in *CollectionGetRequest, out *CollectionGetResponse) error {
	return h.CollectionHandler.Get(ctx, in, out)
}
