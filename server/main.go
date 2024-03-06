package main

import (
	"log"
	"net"

	pb "github.com/ikrammert/golang-demo-grpc/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	//listen on the port
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start the server\nServisi başlatma başarısız%+v", err)
	}
	// create a new gRPC server
	grpcServer := grpc.NewServer()

	// register the greet service
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})

	log.Printf("Server Started at %v", lis.Addr())

	//list is the port, the grpc server needs to start there
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
}
