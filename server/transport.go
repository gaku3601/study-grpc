package main

import (
	pb "github.com/gaku3601/study-grpc/server/pb"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"golang.org/x/net/context"
)

type grpcServer struct {
	getAdminInfo grpctransport.Handler
}

func (s *grpcServer) GetAdminInfo(ctx context.Context, r *pb.GetAdminInfoRequest) (*pb.GetAdminInfoResponse, error) {
	_, resp, err := s.getAdminInfo.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.GetAdminInfoResponse), nil
}

func NewGRPCServer(ctx context.Context, endpoint Endpoints) pb.AdminServer {
	return &grpcServer{
		getAdminInfo: grpctransport.NewServer(
			endpoint.GetAdminInfoEndpoint,
			DecodeGetAdminInfoRequest,
			EncodeGetAdminInfoResponse,
		),
	}
}
