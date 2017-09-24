package main

import dry "github.com/ungerik/go-dry"

func (s *server) applyWebhook(job *Job) error {
	debug("Apply webhook %s %v", s.webhook, job)
	return dry.HTTPPostJSON(s.webhook, job.ToWebhookData())
}

type WebhookData struct {
	ExternalID string `json:"external_id"`
	Error      string `json:"error,omitempty"`
	HasError   bool   `json:"has_error"`
	SocialID   string `json:"social_id"`
}

func (job *Job) ToWebhookData() WebhookData {
	return WebhookData{
		ExternalID: job.PostJob.ExternalId,
		Error:      job.Error,
		HasError:   job.Error != "",
		SocialID:   job.SocialID,
	}
}
