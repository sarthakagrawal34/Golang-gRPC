package main

import (
	pb "go-grpc/proto"
	"io"
	"log"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	log.Println("SayHelloClientStreaming server is started")

	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("SayHelloClientStreaming server is finished")
			return stream.SendAndClose(&pb.MessagesList{Messages: messages})
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with name: %v", req.Name)
		messages = append(messages, "Hello " + req.Name)
	}
}
