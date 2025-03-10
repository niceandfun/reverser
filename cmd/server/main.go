package main

import (
	"context"
	"log"
	"net"

	"reverser/internal/services/wisdom"
	pb "reverser/proto/reverser"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type server struct {
	pb.UnimplementedReverserServiceServer
}

// Reverse takes a string message and returns its reverse version.
// It also attaches wisdom metadata to the response.
//
// Parameters:
//   - ctx: The context for the RPC call
//   - in: The ReverseRequest containing the message to reverse
//
// Returns:
//   - *ReverseResponse: Contains the reversed message
//   - error: Any error that occurred during processing
func (s *server) Reverse(ctx context.Context, in *pb.ReverseRequest) (*pb.ReverseResponse, error) {
	msg := in.GetMsg()
	new_msg := ""

	wsdm, err := wisdom.New()

	for i := len(msg) - 1; i >= 0; i-- {
		new_msg += string(msg[i])
	}

	header := metadata.New(map[string]string{
		"wisdom": wsdm.String(),
	})
	grpc.SendHeader(ctx, header)

	return &pb.ReverseResponse{Msg: new_msg}, err
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
