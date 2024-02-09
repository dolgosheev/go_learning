// Package test_service gRPC сервер с рефлексией
package test_service

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"math/rand"
	"net"

	pb "TestProject/gen/go/test_service/v1"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server определение структуры
type server struct {
	pb.UnimplementedTestServiceServer
}

// GetRandomNumber имплементация метода из контракта
func (s *server) GetRandomNumber(ctx context.Context, in *pb.GetRandomNumberRequest) (*pb.GetRandomNumberResponse, error) {

	log.Printf("Received: %v", "Test")
	return &pb.GetRandomNumberResponse{Value: rand.Int63n(10) + 1}, nil
}

// Register регистрация gRPC сервера
func Register() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTestServiceServer(s, &server{})
	// подключение рефлексии
	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
