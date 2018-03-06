package main

import (
	"context"
	"fmt"

	pb "github.com/gaku3601/study-grpc/server/pb"
)

type AdminService struct{}

func (a *AdminService) GetAdminInfo(ctx context.Context, message *pb.GetAdminInfoMessage) (*pb.AdminInfoResponse, error) {
	fmt.Println(message.TargetAdminName)
	return &pb.AdminInfoResponse{
		Name: "gaku",
		Age:  "26",
	}, nil
}
