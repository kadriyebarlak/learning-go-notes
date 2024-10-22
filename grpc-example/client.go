package main

import (
	"context"
	"log"
	"time"

	pb "learning/grpc-example/helloworld"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// say hello request
	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "World Hanna"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting: %s", r.GetMessage())

	//say goodbye request
	farewellClient := pb.NewFarewellClient(conn)

	goodbyeResponse, err := farewellClient.SayGoodbye(ctx, &pb.GoodbyeRequest{Name: "Hanna"})
	if err != nil {
		log.Fatalf("could not say goodbye: %v", err)
	}
	log.Printf("Farewell: %s", goodbyeResponse.GetMessage())
}
