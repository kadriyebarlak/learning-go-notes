package main

import (
	"context"
	"log"
	"net"

	pb "learning/grpc-example/helloworld"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

type farewellServer struct {
	pb.UnimplementedFarewellServer
}

func (s *farewellServer) SayGoodbye(ctx context.Context, in *pb.GoodbyeRequest) (*pb.GoodbyeReply, error) {
	log.Printf("Saying Goodbye to: %v", in.GetName())
	return &pb.GoodbyeReply{Message: "Goodbye " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	pb.RegisterFarewellServer(s, &farewellServer{})

	log.Println("Server is running at :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
