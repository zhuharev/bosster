package main

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pbserver "github.com/zhuharev/bosster/grpc/server"
)

const (
	port = ":2020"

	// Debug enable std logging
	Debug = true
)

func main() {

	webhook := os.Getenv("BOSSTER_WEBHOOK")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pbserver.RegisterPosterServer(s, newServer(webhook))
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
