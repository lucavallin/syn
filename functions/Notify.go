// Package functions contains Google Cloud Functions
package functions

import (
	"bytes"
	"cavall.in/syn/events"
	"context"
	"encoding/json"
	"github.com/thoas/go-funk"
	"log"
	"net/http"
	"os"
	"strings"
)

// IftttNotification represents a notification to IFTTT
type IftttNotification struct {
	Labels string `json:"value1"`
	ImageUrl string `json:"value2"`
}

// Notify is triggered by a change to a Firestore document.
func Notify(ctx context.Context, e events.FirestoreEvent) error {
	log.Printf("Event received: %v", e.Value.Name)
	iftttWebhookUrl := os.Getenv("IFTTT_WEBHOOK_URL")

	labels := funk.Map(e.Value.Fields.Labels.ArrayValue.Values, func(l events.LabelValues) string {
		return l.MapValue.Fields.Description.StringValue
	}).([]string)

	notification, err := json.Marshal(IftttNotification{
		Labels: strings.Join(labels, ", "),
		ImageUrl: "",
	})
	if err != nil {
		return err
	}

	_, err = http.Post(iftttWebhookUrl, "application/json", bytes.NewBuffer(notification))
	if err != nil {
		return err
	}

	log.Printf("Notification sent to IFTTT: %s", notification)

	return nil
}
