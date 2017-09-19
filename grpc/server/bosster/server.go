package main

import (
	"log"

	"golang.org/x/net/context"

	pbserver "github.com/zhuharev/bosster/grpc/server"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	*Queue
}

func newServer() *server {
	return &server{Queue: NewQueue()}
}

// Post implements interface of Boster
func (s *server) Post(ctx context.Context, in *pbserver.PostRequest) (*pbserver.PostReply, error) {
	log.Println("post")
	var enqueuedJobs Jobs
	for _, target := range in.Targets {
		job := s.Enqueue(in, target, target.Type)
		enqueuedJobs = append(enqueuedJobs, job)
	}

	log.Println(in.Sync)

	if in.Sync {
		for _, job := range enqueuedJobs {
			job.Wait()
		}
	}

	return enqueuedJobs.ConvertToReply(), nil
}
