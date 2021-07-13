// Package functions contains Google Cloud Functions
package functions

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

//FirestoreEvent is the payload of a Firestore event.
type FirestoreEvent struct {
	OldValue   FirestoreValue `json:"oldValue"`
	Value      FirestoreValue `json:"value"`
	UpdateMask struct {
		FieldPaths []string `json:"fieldPaths"`
	} `json:"updateMask"`
}

// FirestoreValue holds Firestore fields.
type FirestoreValue struct {
	CreateTime time.Time `json:"createTime"`
	// Fields is the data for this value. The type depends on the format of your
	// database. Log the interface{} value and inspect the result to see a JSON
	// representation of your database fields.
	Fields     interface{} `json:"fields"`
	Name       string      `json:"name"`
	UpdateTime time.Time   `json:"updateTime"`
}

// IftttNotification represents a notification to IFTTT
type IftttNotification struct {
	Labels string `json:"value1,omitempty"`
	ImageUrl string `json:"value2,omitempty"`
}

// Notify is triggered by a change to a Firestore document.
func Notify(ctx context.Context, e FirestoreEvent) error {
	log.Printf("Event received: %v", e)
	iftttWebhookUrl := os.Getenv("IFTTT_WEBHOOK_URL")

	notification, _ := json.Marshal(IftttNotification{
		Labels: "Hello",
		ImageUrl: "",
	})

	_, err := http.Post(iftttWebhookUrl, "application/json", bytes.NewBuffer(notification))
	if err != nil {
		return err
	}

	log.Printf("Notification sent to IFTTT: %v", notification)

	return nil
}