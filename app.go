package main

import (
	"context"
	"log"
	"net"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"

	sq "github.com/ciochetta/go-square/grpc"
)

// Declaring the service that will be exposed over GRPC
type server struct {
	sq.UnimplementedSquareServer
}

// Our exposed method
func (s server) GetSquare(ctx context.Context, in *sq.GetSquareRequest) (*sq.GetSquareResponse, error) {

	log.Println("Received request: ", in.GetNumber())

	return &sq.GetSquareResponse{
		Number: in.Number * in.Number,
	}, nil

}

func run() error {

	lis, err := net.Listen("tcp", ":1234")

	if err != nil {
		return err
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
	)

	sq.RegisterSquareServer(s, &server{})

	log.Println("Listening on port 1234")

	if err := s.Serve(lis); err != nil {
		return err
	}

	return nil

}
