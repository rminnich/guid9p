package main

import (
	"context"
	"log"
	"time"

	pb "github.com/rminnich/guid9p/generated/proto/ninep"

	"google.golang.org/grpc"
)

func smain() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewNinePServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Walk(ctx, &pb.WalkRequest{Guid: "some-guid", NewGuid: "new-guid", Names: []string{"path", "to", "file"}})
	if err != nil {
		log.Fatalf("could not walk: %v", err)
	}
	log.Printf("Walk response: %v", r.GetGuids())
}
