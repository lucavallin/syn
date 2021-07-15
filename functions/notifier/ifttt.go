package notifier

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type IftttNotifier struct {
	WebhookUrl string `json:"webhook_url"`
}

// IftttNotification represents a notification to IFTTT
type IftttNotification struct {
	Labels   string `json:"value1"`
	ImageUrl string `json:"value2"`
}

func NewIftttNotifier(webhookUrl string) *IftttNotifier {
	return &IftttNotifier{webhookUrl}
}

func NewIftttNotification(labels string, imageUrl string) *IftttNotification {
	return &IftttNotification{labels, imageUrl}
}

// Notify triggers a webhook-based IFTTT notification
func (n *IftttNotifier) Notify(notification *IftttNotification) error {
	body, err := json.Marshal(notification)
	if err != nil {
		return err
	}

	_, err = http.Post(n.WebhookUrl, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	log.Printf("Notification sent to IFTTT: %s", body)

	return nil
}
