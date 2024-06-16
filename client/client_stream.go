package main

import (
	"context"
	pb "go-grpc/proto"
	"log"
	"time"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Client side streaming started")

	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send names due to error: %v\n", err)
	}

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("error while sending request due to error: %v\n", err)
		}
		log.Printf("Sent the request with name: %s\n", name)
		time.Sleep(1 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response due to error: %v\n", err)
	}
	for _, message := range res.Messages {
		log.Printf("%v\n", message)
	}
}
