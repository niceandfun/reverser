package main

import (
	"context"
	"log"
	"net"

	pb "reverser/proto/reverser"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedReverserServiceServer
}

func (s *server) Reverse(ctx context.Context, in *pb.ReverseRequest) (*pb.ReverseResponse, error) {
	return &pb.ReverseResponse{Msg: ""}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterReverserServiceServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
