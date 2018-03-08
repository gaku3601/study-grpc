package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/net/context"

	pb "github.com/gaku3601/study-grpc/server/pb"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()

	var svc AdminService
	svc = impl{}
	errChan := make(chan error)

	endpoints := Endpoints{
		GetAdminInfoEndpoint: MakeGetAdminInfoEndpoint(svc),
	}

	go func() {
		listener, err := net.Listen("tcp", ":19003")
		if err != nil {
			errChan <- err
			return
		}
		handler := NewGRPCServer(ctx, endpoints)
		gRPCServer := grpc.NewServer()
		pb.RegisterAdminServer(gRPCServer, handler)
		errChan <- gRPCServer.Serve(listener)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()
	fmt.Println(<-errChan)
}
