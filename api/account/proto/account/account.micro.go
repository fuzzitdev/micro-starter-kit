// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: api/account/proto/account/account.proto

package api_account

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	proto1 "github.com/micro/go-micro/api/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Client API for AccountService service

type AccountService interface {
	List(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error)
}

type accountService struct {
	c    client.Client
	name string
}

func NewAccountService(name string, c client.Client) AccountService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "api.account"
	}
	return &accountService{
		c:    c,
		name: name,
	}
}

func (c *accountService) List(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error) {
	req := c.c.NewRequest(c.name, "AccountService.List", in)
	out := new(proto1.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AccountService service

type AccountServiceHandler interface {
	List(context.Context, *proto1.Request, *proto1.Response) error
}

func RegisterAccountServiceHandler(s server.Server, hdlr AccountServiceHandler, opts ...server.HandlerOption) error {
	type accountService interface {
		List(ctx context.Context, in *proto1.Request, out *proto1.Response) error
	}
	type AccountService struct {
		accountService
	}
	h := &accountServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&AccountService{h}, opts...))
}

type accountServiceHandler struct {
	AccountServiceHandler
}

func (h *accountServiceHandler) List(ctx context.Context, in *proto1.Request, out *proto1.Response) error {
	return h.AccountServiceHandler.List(ctx, in, out)
}
