package rpcdemo

import (
	"context"
	"google.golang.org/grpc"
	"log"
)

type Server struct {
}

// MakegRPCRequest make outbound gRPC request
func MakegRPCRequest() {
	conn, err := grpc.Dial(":8120", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	defer conn.Close()

	c := NewOutServiceClient(conn)

	response, err := c.MakeRequest(context.Background(), &OutRequest{Content: "testMessage"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Content)
}

