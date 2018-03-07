package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/gaku3601/study-grpc/server/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:19003", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()
	client := pb.NewAdminClient(conn)
	message := &pb.GetAdminInfoRequest{"こんちゃす!"}
	res, err := client.GetAdminInfo(context.TODO(), message)
	fmt.Printf("%#v\n", res.Name)
	fmt.Printf("%#v\n", res.Age)
}
