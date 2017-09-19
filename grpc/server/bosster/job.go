package main

import pbserver "github.com/zhuharev/bosster/grpc/server"

type Job struct {
	PostJob  *pbserver.PostJob
	PostReq  *pbserver.PostRequest
	Type     pbserver.SocialType
	done     chan struct{}
	isDone   bool
	SocialID string
	Error    string
	Status   pbserver.STATUS
}

func NewJob(postReq *pbserver.PostRequest, postJob *pbserver.PostJob, socType pbserver.SocialType) *Job {
	debug("New job %v", postReq)
	return &Job{
		PostJob: postJob,
		PostReq: postReq,
		Type:    socType,

		done: make(chan struct{}, 1),
	}
}

func (j *Job) Wait() {
	<-j.done
}

func (j *Job) doneJob() {
	j.done <- struct{}{}
	j.isDone = true
}

func (j *Job) IsDone() bool {
	return j.isDone
}

type Jobs []*Job

func (js Jobs) ConvertToReply() *pbserver.PostReply {
	jobs := []*pbserver.PostJob{}
	for _, job := range js {
		jobs = append(jobs, job.PostJob)
	}
	return &pbserver.PostReply{Jobs: jobs}
}
