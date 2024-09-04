package main

import (
	"context"
	"log"
	"net"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

// Server is used to implement the NinePService.
type server struct {
	pb.UnimplementedNinePServiceServer
}

// Implement the Walk method.
func (s *server) Walk(ctx context.Context, req *pb.WalkRequest) (*pb.WalkResponse, error) {
	// Example implementation of the Walk method.
	guids := []string{
		uuid.New().String(), // Generate example GUIDs.
		uuid.New().String(),
	}
	return &pb.WalkResponse{Guids: guids}, nil
}

// Implement other methods (Open, Read, Write, Create, Stat, Clunk, Flush) similarly.

func cmain() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterNinePServiceServer(s, &server{})

	log.Printf("Server listening on %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
