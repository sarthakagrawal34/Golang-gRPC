package main

import (
	"context"
	pb "go-grpc/proto"
	"io"
	"log"
	"time"
)

func callSayHelloBiDirectionalStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Bi directional streaming started")

	stream, err := client.SayHelloBiDirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send names due to error: %v\n", err)
	}

	waitc := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("error while streaming due to error: %v\n", err)
			}
			log.Printf("%s", message.Message)
		}
		close(waitc)
	}()

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
	stream.CloseSend()
	<-waitc
	log.Printf("Bi directional streaming finished")
	
}
