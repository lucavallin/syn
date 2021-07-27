package functions

import (
	"cavall.in/syn/events"
	"cavall.in/syn/notifier"
	"context"
	"log"
	"os"
	"strings"
)

// Notify is triggered by a change to a Firestore document and it sends a notification to IFTTT
func Notify(ctx context.Context, e events.FirestoreEvent) error {
	log.Printf("Event received: %+v", e)
	iftttWebhookUrl := os.Getenv("IFTTT_WEBHOOK_URL")

	labels := e.GetUploadLabels()

	ifttt := notifier.NewIftttNotifier(iftttWebhookUrl)
	notification := notifier.NewIftttNotification(strings.Join(labels, ", "), "")
	if err := ifttt.Notify(notification); err != nil {
		return err
	}

	return nil
}
