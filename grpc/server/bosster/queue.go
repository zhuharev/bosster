package main

import (
	pbserver "github.com/zhuharev/bosster/grpc/server"
)

var queuePoolSize = 1000

type Queue struct {
	ch chan *Job
}

func NewQueue() *Queue {
	return &Queue{ch: make(chan *Job, queuePoolSize)}
}

func (q *Queue) Start() {
	go func() {
		for job := range q.ch {
			post(job)
		}
	}()
}

func (q *Queue) Enqueue(postJob *pbserver.PostRequest, socType pbserver.SocialType) *Job {
	j := NewJob(postJob, socType)
	q.ch <- j
	return j
}
