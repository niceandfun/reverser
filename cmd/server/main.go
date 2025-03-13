package main

import (
	"context"
	"log"
	"net"
	"strings"

	"reverser/constants/hdrs"
	"reverser/internal/wisdom"
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
	var builder strings.Builder

	wsdm, err := wisdom.New()
	if err != nil {
		log.Print(err)
		return nil, err
	}

	for i := len(msg) - 1; i >= 0; i-- {
		builder.WriteString(string(msg[i]))
	}

	headers := metadata.New(map[string]string{
		hdrs.Wisdom: wsdm,
	})

	grpc.SendHeader(ctx, headers)

	return &pb.ReverseResponse{Msg: builder.String()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50050")

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
