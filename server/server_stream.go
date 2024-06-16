package main

import (
	pb "go-grpc/proto"
	"log"
	"time"
)

func (s *helloServer) SayHelloServerStreaming(in *pb.NamesList, stream pb.GreetService_SayHelloServerStreamingServer) error {
	log.Printf("SayHelloServerStreaming is called with names: %v\n", in.Names)

	for _, name := range in.Names{
		res := &pb.HelloResponse{
			Message: "Hello "+ name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(time.Second * 2)
	}
	return nil
}
