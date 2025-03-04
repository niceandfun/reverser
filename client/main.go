package main

import (
	"context"
	"log"
	"time"

	pb "reverser/proto/reverser" // Путь к вашему сгенерированному protobuf пакету

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("не удалось подключиться: %v", err)
	}
	defer conn.Close()

	client := pb.NewReverserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.Reverse(ctx, &pb.ReverseRequest{})
	if err != nil {
		log.Fatalf("ошибка при вызове метода: %v", err)
	}

	log.Printf("Ответ от сервера: %v", response)
}
