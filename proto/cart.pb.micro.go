// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/cart.proto

package cart

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/asim/go-micro/v3/api"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Cart service

func NewCartEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Cart service

type CartService interface {
	Add(ctx context.Context, in *CartInfo, opts ...client.CallOption) (*ResAdd, error)
	Clear(ctx context.Context, in *Clean, opts ...client.CallOption) (*Res, error)
	Incr(ctx context.Context, in *Item, opts ...client.CallOption) (*Res, error)
	Decr(ctx context.Context, in *Item, opts ...client.CallOption) (*Res, error)
	DelItemById(ctx context.Context, in *CartId, opts ...client.CallOption) (*Res, error)
	GetAll(ctx context.Context, in *CartFindAll, opts ...client.CallOption) (*CartAll, error)
}

type cartService struct {
	c    client.Client
	name string
}

func NewCartService(name string, c client.Client) CartService {
	return &cartService{
		c:    c,
		name: name,
	}
}

func (c *cartService) Add(ctx context.Context, in *CartInfo, opts ...client.CallOption) (*ResAdd, error) {
	req := c.c.NewRequest(c.name, "Cart.Add", in)
	out := new(ResAdd)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartService) Clear(ctx context.Context, in *Clean, opts ...client.CallOption) (*Res, error) {
	req := c.c.NewRequest(c.name, "Cart.Clear", in)
	out := new(Res)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartService) Incr(ctx context.Context, in *Item, opts ...client.CallOption) (*Res, error) {
	req := c.c.NewRequest(c.name, "Cart.Incr", in)
	out := new(Res)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartService) Decr(ctx context.Context, in *Item, opts ...client.CallOption) (*Res, error) {
	req := c.c.NewRequest(c.name, "Cart.Decr", in)
	out := new(Res)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartService) DelItemById(ctx context.Context, in *CartId, opts ...client.CallOption) (*Res, error) {
	req := c.c.NewRequest(c.name, "Cart.DelItemById", in)
	out := new(Res)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartService) GetAll(ctx context.Context, in *CartFindAll, opts ...client.CallOption) (*CartAll, error) {
	req := c.c.NewRequest(c.name, "Cart.GetAll", in)
	out := new(CartAll)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Cart service

type CartHandler interface {
	Add(context.Context, *CartInfo, *ResAdd) error
	Clear(context.Context, *Clean, *Res) error
	Incr(context.Context, *Item, *Res) error
	Decr(context.Context, *Item, *Res) error
	DelItemById(context.Context, *CartId, *Res) error
	GetAll(context.Context, *CartFindAll, *CartAll) error
}

func RegisterCartHandler(s server.Server, hdlr CartHandler, opts ...server.HandlerOption) error {
	type cart interface {
		Add(ctx context.Context, in *CartInfo, out *ResAdd) error
		Clear(ctx context.Context, in *Clean, out *Res) error
		Incr(ctx context.Context, in *Item, out *Res) error
		Decr(ctx context.Context, in *Item, out *Res) error
		DelItemById(ctx context.Context, in *CartId, out *Res) error
		GetAll(ctx context.Context, in *CartFindAll, out *CartAll) error
	}
	type Cart struct {
		cart
	}
	h := &cartHandler{hdlr}
	return s.Handle(s.NewHandler(&Cart{h}, opts...))
}

type cartHandler struct {
	CartHandler
}

func (h *cartHandler) Add(ctx context.Context, in *CartInfo, out *ResAdd) error {
	return h.CartHandler.Add(ctx, in, out)
}

func (h *cartHandler) Clear(ctx context.Context, in *Clean, out *Res) error {
	return h.CartHandler.Clear(ctx, in, out)
}

func (h *cartHandler) Incr(ctx context.Context, in *Item, out *Res) error {
	return h.CartHandler.Incr(ctx, in, out)
}

func (h *cartHandler) Decr(ctx context.Context, in *Item, out *Res) error {
	return h.CartHandler.Decr(ctx, in, out)
}

func (h *cartHandler) DelItemById(ctx context.Context, in *CartId, out *Res) error {
	return h.CartHandler.DelItemById(ctx, in, out)
}

func (h *cartHandler) GetAll(ctx context.Context, in *CartFindAll, out *CartAll) error {
	return h.CartHandler.GetAll(ctx, in, out)
}