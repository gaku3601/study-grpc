package main

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	AdminEndpoint endpoint.Endpoint
}

func MakeGetAdmininfoEndpoint(a AdminService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAdminInfoMessage)
		output, err := a.GetAdminInfo(req.TargetAdminName)
		if err != nil {
			return nil, err
		}
		return AdminInfoResponse{Name: output.Name, Age: output.Age}, nil
	}
}

func (e Endpoints) Admin(ctx context.Context, targetAdminName string) (string, string) {
	req := GetAdminInfoMessage{
		TargetAdminName: targetAdminName,
	}
	resp, _ := e.AdminEndpoint(ctx, req)
	//if err != nil {
	//	return "", err
	//}
	adminResp := resp.(AdminInfoResponse)
	//if adminResp.Err != "" {
	//	return "", errors.New(adminResp.Err)
	//}
	return adminResp.Name, adminResp.Age
}

//request
type GetAdminInfoMessage struct {
	TargetAdminName string
}

//response
type AdminInfoResponse struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}
