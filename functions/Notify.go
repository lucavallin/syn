// Package functions contains Google Cloud Functions
package functions

import (
	"context"
	"log"
)

// PubSubMessage is the payload of a Pub/Sub event.
type PubSubMessage struct {
	Data []byte `json:"data"`
}

// Notify consumes a Pub/Sub message.
func Notify(ctx context.Context, m PubSubMessage) error {
	log.Printf("%v", m)
	return nil
}