package main

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetAdminInfoEndpoint endpoint.Endpoint
}

func New(svc AdminService) Endpoints {
	return Endpoints{
		GetAdminInfoEndpoint: MakeGetAdminInfoEndpoint(svc),
	}
}

func MakeGetAdminInfoEndpoint(a AdminService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAdminInfoRequest)
		output, err := a.GetAdminInfo(req.TargetAdminName)
		if err != nil {
			return nil, err
		}
		return GetAdminInfoResponse{Name: output.Name, Age: output.Age}, nil
	}
}

//request
type GetAdminInfoRequest struct {
	TargetAdminName string
}

//response
type GetAdminInfoResponse struct {
	Name string
	Age  string
}
