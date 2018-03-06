package main

import (
	"log"
	"net"

	pb "github.com/gaku3601/study-grpc/server/pb"
	"google.golang.org/grpc"
)

func main() {
	listenPort, err := net.Listen("tcp", ":19003")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	adminService := &AdminService{}

	//サービスを登録する
	pb.RegisterAdminServer(server, adminService)
	server.Serve(listenPort)
}
