package main

import (
	"golang.org/x/net/context"

	pbserver "github.com/zhuharev/bosster/grpc/server"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	*Queue
}

func newServer() *server {
	return &server{Queue: &Queue{}}
}

// Post implements interface of Boster
func (s *server) Post(ctx context.Context, in *pbserver.PostRequest) (*pbserver.PostReply, error) {

	var enqueuedJobs Jobs
	for _, target := range in.Targets {
		job := s.Enqueue(in, target)
		enqueuedJobs = append(enqueuedJobs, job)
	}

	if !in.Async {
		for _, job := range enqueuedJobs {
			job.Wait()
		}
	}

	return enqueuedJobs.ConvertToReply(), nil
}
