package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	sq "github.com/ciochetta/go-square/grpc"
)

// Declaring the service that will be exposed over GRPC
type server struct {
	sq.UnimplementedSquareServer
}

// Our exposed method
func (s server) GetSquare(ctx context.Context, in *sq.GetSquareRequest) (*sq.GetSquareResponse, error) {

	return &sq.GetSquareResponse{
		Number: in.Number * in.Number,
	}, nil

}

func run() {

	lis, err := net.Listen("tcp", ":1234")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	sq.RegisterSquareServer(s, &server{})

	log.Println("Listening on port 1234")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
