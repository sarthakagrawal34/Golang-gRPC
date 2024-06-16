package main

import (
	pb "go-grpc/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const(
	port = ":9080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect the client to the server %v", err)
		return
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Names: []string{"Sarthak", "Jack", "Jessy", "Raj", "Walt"},
	}

	// Unary rpc
	// callSayHello(client)

	// Server streaming rpc
	// callSayHelloServerStream(client, names)

	// Client streaming rpc
	callSayHelloClientStream(client, names)
}