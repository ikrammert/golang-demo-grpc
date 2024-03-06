package main

import (
	"context"
	"log"
	"time"

	pb "github.com/ikrammert/golang-demo-grpc/proto"
)

func callSayHello(client pb.GreetServiceServer) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("%s", res.Message)
}
