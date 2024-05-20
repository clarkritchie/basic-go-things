package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/clarkritchie/basic-go-things/grpc/helloworld"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedHelloWorldServiceServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloWorldRequest) (*pb.HelloWorldResponse, error) {
	// Get the current time
	currentTime := time.Now()
	currentTimeString := currentTime.Format("2006-01-02 15:04:05")
	message := fmt.Sprintf("Hello world, the time is %v", currentTimeString)
	return &pb.HelloWorldResponse{Message: message}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen on port 50051: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterHelloWorldServiceServer(s, &server{})
	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
