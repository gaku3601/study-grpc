package main

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	AdminEndpoint endpoint.Endpoint
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

func (e Endpoints) Admin(ctx context.Context, targetAdminName string) (string, string) {
	req := GetAdminInfoRequest{
		TargetAdminName: targetAdminName,
	}
	resp, _ := e.AdminEndpoint(ctx, req)
	//if err != nil {
	//	return "", err
	//}
	adminResp := resp.(GetAdminInfoResponse)
	//if adminResp.Err != "" {
	//	return "", errors.New(adminResp.Err)
	//}
	return adminResp.Name, adminResp.Age
}

//request
type GetAdminInfoRequest struct {
	TargetAdminName string
}

//response
type GetAdminInfoResponse struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}
