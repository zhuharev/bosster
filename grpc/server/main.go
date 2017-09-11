package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) Post(ctx context.Context, in *PostRequest) (*PostReply, error) {
	//emulate enqueue
	pj := &PostJob{
		SocialId:       "0",
		SocialProvider: SocialType_FACEBOOK,
		Status:         STATUS_ENQUEUED,
	}
	if !in.Async {
		pj.Status = STATUS_COMPLETED
		pj.SocialId = "1"
	}
	return &PostReply{Jobs: []*PostJob{
		pj,
	}}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterPosterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
