package main

import (
	pb "github.com/gaku3601/study-grpc/server/pb"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"golang.org/x/net/context"
)

type grpcServer struct {
	admin grpctransport.Handler
}

// implement LoremServer Interface in lorem.pb.go
func (s *grpcServer) GetAdminInfo(ctx context.Context, r *pb.GetAdminInfoRequest) (*pb.GetAdminInfoResponse, error) {
	_, resp, err := s.admin.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.GetAdminInfoResponse), nil
}

// create new grpc server
func NewGRPCServer(ctx context.Context, endpoint Endpoints) pb.AdminServer {
	return &grpcServer{
		admin: grpctransport.NewServer(
			endpoint.AdminEndpoint,
			DecodeGetAdminInfoRequest,
			EncodeGetAdminInfoResponse,
		),
	}
}
