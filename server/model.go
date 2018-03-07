package main

import (
	"context"

	pb "github.com/gaku3601/study-grpc/server/pb"
)

func DecodeGetAdminInfoRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.GetAdminInfoRequest)
	return GetAdminInfoRequest{
		TargetAdminName: req.TargetAdminName,
	}, nil
}

func EncodeGetAdminInfoResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(GetAdminInfoResponse)
	return &pb.GetAdminInfoResponse{
		Name: resp.Name,
		Age:  resp.Age,
	}, nil
}
