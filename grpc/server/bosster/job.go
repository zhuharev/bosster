package main

import (
	pbserver "github.com/zhuharev/bosster/grpc/server"
)

type Job struct {
	PostReq  *pbserver.PostRequest
	Type     pbserver.SocialType
	done     chan struct{}
	isDone   bool
	SocialID string
	Error    string
	Status   pbserver.STATUS
}

func NewJob(postReq *pbserver.PostRequest, socType pbserver.SocialType) *Job {
	return &Job{
		PostReq: postReq,
		Type:    socType,
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
		jobs = append(jobs, &pbserver.PostJob{
			SocialId:       job.SocialID,
			SocialOwnerId:  job.PostReq.GetSocialId(),
			SocialProvider: job.Type,
			Status:         job.Status,
			Error:          job.Error,
		})
	}
	return &pbserver.PostReply{Jobs: jobs}
}
