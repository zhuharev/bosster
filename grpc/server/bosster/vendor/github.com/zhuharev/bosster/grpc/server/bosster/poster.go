package main

import (
	"fmt"

	"github.com/Unknwon/com"
	"github.com/zhuharev/bosster"
	pbserver "github.com/zhuharev/bosster/grpc/server"
)

func (s *server) post(job *Job) {
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

func postToVk(job *Job) (err error) {
	debug("Post to vk")
	intID, err := bosster.PostToVk(bosster.Credentials{
		VkAccessToken: job.PostJob.GetSocialToken(),
		VkOwnerID:     -com.StrTo(job.PostJob.GetSocialId()).MustInt(),
	},
		bosster.Post{Body: job.PostReq.Post.Message,
			ImageURLs: job.PostReq.Post.ImageUrls})
	if err != nil {
		return err
	}

	job.SocialID = fmt.Sprint(intID)

	return nil
}

func postToFb(job *Job) error {
	debug("Post to fb")
	credentials := bosster.Credentials{
		FacebookToken:   job.PostJob.GetSocialToken(),
		FacebookGroupID: job.PostJob.GetSocialId(),
	}
	post := bosster.Post{
		Body:      job.PostReq.Post.Message,
		ImageURLs: job.PostReq.Post.ImageUrls,
	}

	socID, err := bosster.PostToFacebook(credentials, post)

	job.SocialID = socID
	return err
}
