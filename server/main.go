package main

import (
	pb "go-grpc/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":9080"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start the server %v", err)
		return
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("server started at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start the grpc server %v", err)
		return
	}
}
