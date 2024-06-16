package main

import (
	"context"
	pb "go-grpc/proto"
)

func (s *helloServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello " + in.Name,
	}, nil
}
