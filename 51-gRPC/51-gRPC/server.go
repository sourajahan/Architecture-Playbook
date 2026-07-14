package main

import (
	"context"
	"google.golang.org/grpc"
	"net"
)

// Server implements the generated OrderServiceServer interface
type server struct {
	UnimplementedOrderServiceServer
}

func (s *server) GetOrder(ctx context.Context, req *OrderRequest) (*OrderResponse, error) {
	// Business logic goes here
	return &OrderResponse{Status: "SHIPPED", TotalAmount: 99.99}, nil
}

func main() {
	lis, _ := net.Listen("tcp", ":50051")
	s := grpc.NewServer()
	RegisterOrderServiceServer(s, &server{})
	s.Serve(lis)
}
