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
	"reflect"
	"strings"
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
	Fields     events.FirestoreUpload  `json:"fields"`
	Name       string      `json:"name"`
	UpdateTime time.Time   `json:"updateTime"`
}

// IftttNotification represents a notification to IFTTT
type IftttNotification struct {
	Labels string `json:"value1"`
	ImageUrl string `json:"value2"`
}

// Notify is triggered by a change to a Firestore document.
func Notify(ctx context.Context, e FirestoreEvent) error {
	log.Printf("Event received: %v", e.Value.Name)
	iftttWebhookUrl := os.Getenv("IFTTT_WEBHOOK_URL")

	type X = reflect.Type(e.Value.Fields.Labels[0])
	labels := funk.Map(e.Value.Fields.Labels, func(l X) string {
		return l.Description
	}).([]string)

	notification := IftttNotification{
		Labels: strings.Join(labels, ", "),
		ImageUrl: "",
	}
	jsonNotification, _ := json.Marshal(IftttNotification{
		Labels: strings.Join(labels, ", "),
		ImageUrl: "",
	})

	_, err := http.Post(iftttWebhookUrl, "application/json", bytes.NewBuffer(jsonNotification))
	if err != nil {
		return err
	}

	log.Printf("Notification sent to IFTTT: %v", notification)

	return nil
}
