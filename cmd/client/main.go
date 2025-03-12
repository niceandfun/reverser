package main

import (
	"context"
	"log"
	"time"

	"reverser/constants/hdrs"
	pb "reverser/proto/reverser"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {
	conn, err := grpc.NewClient("localhost:50050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("не удалось подключиться: %v", err)
	}
	defer conn.Close()

	client := pb.NewReverserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var md metadata.MD

	response, err := client.Reverse(
		ctx,
		&pb.ReverseRequest{Msg: "abcdef"},
		grpc.Header(&md),
	)
	if err != nil {
		log.Fatalf("ошибка при вызове метода: %v", err)
	}

	log.Printf("Response: %v", response)
	log.Printf("Headers: %v", md)
	if wisdom := md.Get(hdrs.Wisdom); len(wisdom) > 0 {
		log.Printf("Random wisdom: %s", wisdom[0])
	}
}
