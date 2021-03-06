// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: server/proto/regulator/regulator.proto

package regulator

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for PointsRegulator service

func NewPointsRegulatorEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for PointsRegulator service

type PointsRegulatorService interface {
	RegulateSetsOfFourPoints(ctx context.Context, in *PointsRequest, opts ...client.CallOption) (*PointsResponse, error)
}

type pointsRegulatorService struct {
	c    client.Client
	name string
}

func NewPointsRegulatorService(name string, c client.Client) PointsRegulatorService {
	return &pointsRegulatorService{
		c:    c,
		name: name,
	}
}

func (c *pointsRegulatorService) RegulateSetsOfFourPoints(ctx context.Context, in *PointsRequest, opts ...client.CallOption) (*PointsResponse, error) {
	req := c.c.NewRequest(c.name, "PointsRegulator.RegulateSetsOfFourPoints", in)
	out := new(PointsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PointsRegulator service

type PointsRegulatorHandler interface {
	RegulateSetsOfFourPoints(context.Context, *PointsRequest, *PointsResponse) error
}

func RegisterPointsRegulatorHandler(s server.Server, hdlr PointsRegulatorHandler, opts ...server.HandlerOption) error {
	type pointsRegulator interface {
		RegulateSetsOfFourPoints(ctx context.Context, in *PointsRequest, out *PointsResponse) error
	}
	type PointsRegulator struct {
		pointsRegulator
	}
	h := &pointsRegulatorHandler{hdlr}
	return s.Handle(s.NewHandler(&PointsRegulator{h}, opts...))
}

type pointsRegulatorHandler struct {
	PointsRegulatorHandler
}

func (h *pointsRegulatorHandler) RegulateSetsOfFourPoints(ctx context.Context, in *PointsRequest, out *PointsResponse) error {
	return h.PointsRegulatorHandler.RegulateSetsOfFourPoints(ctx, in, out)
}
