package slackutil

import (
	"net/http"
	"net/url"
)

type Slack struct {
	WebhookURL string
}

func Init(url string) *Slack {
	if url == "" {
		return nil
	}
	s := Slack{WebhookURL: url}
	return &s
}

func (s *Slack) SendWebhook(text string) error {
	_, err := http.PostForm(s.WebhookURL, url.Values{"payload": {string(text)}})
	if err != nil {
		return err
	}
	return nil
}
