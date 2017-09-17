package main

import (
	"fmt"

	"github.com/Unknwon/com"
	"github.com/zhuharev/bosster"
	pbserver "github.com/zhuharev/bosster/grpc/server"
)

func post(job *Job) {
	defer job.doneJob()

	switch job.Type {
	case pbserver.SocialType_VK:
		postToVk(job)
	case pbserver.SocialType_FACEBOOK:
		postToFb(job)
	default:
		job.Error = "Unsupported social type"
	}
}

func postToVk(job *Job) (socialID string, err error) {
	intID, err := bosster.PostToVk(bosster.Credentials{
		VkAccessToken: job.PostReq.GetSocialToken(),
		VkOwnerID:     com.StrTo(job.PostReq.GetSocialId()).MustInt(),
	}, bosster.Post{Body: job.PostReq.Post.Message, ImageURLs: job.PostReq.Post.ImageUrls})
	if err != nil {
		return "", err
	}
	return fmt.Sprint(intID), nil
}

func postToFb(job *Job) (string, error) {
	credentials := bosster.Credentials{
		FacebookToken:   job.PostReq.GetSocialToken(),
		FacebookGroupID: job.PostReq.GetSocialId(),
	}
	post := bosster.Post{
		Body:      job.PostReq.Post.Message,
		ImageURLs: job.PostReq.Post.ImageUrls,
	}
	return bosster.PostToFacebook(credentials, post)
}
