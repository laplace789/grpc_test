package main

import (
	"context"
	"log"
	"net"

	"github.com/laplace789/grpc_test/pb"
	"google.golang.org/grpc"
)

//GreeterServer impelents grpc's GreeterServer interface
type GreeterServer struct{}

//GetValue will random return a value for test
func (s *GreeterServer) GetValue(c context.Context, p *pb.Person) (*pb.Adress, error) {
	a := &pb.Adress{Address: "aaa"}
	return a, nil
}

const (
	port = ":50051"
)

func main() {
	// Create gRPC Server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	g := &GreeterServer{}
	s := grpc.NewServer()
	log.Println("gRPC server is running.")
	pb.RegisterGreeterServer(s, g)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
