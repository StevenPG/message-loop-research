package main

import (
	"fmt"
	"log"
	"message-loop/pkg/rpcdemo"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// https://tutorialedge.net/golang/go-grpc-beginners-tutorial/

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8130))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := rpcdemo.OutServiceServer{}

	grpcServer := grpc.NewServer()

	rpcdemo.RegisterOutServiceServer(grpcServer, s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
