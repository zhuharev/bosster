package main

import (
	"log"

	pbserver "github.com/zhuharev/bosster/grpc/server"
)

var queuePoolSize = 1000

type Queue struct {
	ch chan *Job

	s *server
}

func NewQueue(s *server) *Queue {
	return &Queue{ch: make(chan *Job, queuePoolSize), s: s}
}

// Start return self for chainability
func (q *Queue) Start() *Queue {
	go func() {
		for job := range q.ch {
			q.s.post(job)
			if job.Error != "" && Debug {
				log.Println(job.Error)
			}
			if q.s.hasWebhook {
				q.s.applyWebhook(job)
			}
		}
	}()
	return q
}

func (q *Queue) Enqueue(postReq *pbserver.PostRequest, postJob *pbserver.PostJob, socType pbserver.SocialType) *Job {
	j := NewJob(postReq, postJob, socType)
	q.ch <- j
	return j
}
