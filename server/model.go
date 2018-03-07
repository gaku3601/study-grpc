package main

import (
	"context"

	pb "github.com/gaku3601/study-grpc/server/pb"
)

func DecodeGRPCAdminRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.GetAdminInfoMessage)
	return GetAdminInfoMessage{
		TargetAdminName: req.TargetAdminName,
	}, nil
}

func EncodeGRPCAdminResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(AdminInfoResponse)
	return &pb.AdminInfoResponse{
		Name: resp.Name,
		Age:  resp.Age,
	}, nil
}
