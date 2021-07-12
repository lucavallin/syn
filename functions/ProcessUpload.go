// Package p contains Google Cloud Functions
package p

import (
	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
	vision "cloud.google.com/go/vision/apiv1"
	"context"
	vision3 "google.golang.org/genproto/googleapis/cloud/vision/v1"
	"log"
	"os"
)

// GCSEvent is the payload of a GCS event. Please refer to the docs for
// additional information regarding GCS events.
type GCSEvent struct {
	Bucket string `json:"bucket"`
	Name   string `json:"name"`
}

type Upload struct {
	Event GCSEvent `json:"event"`
	Labels []*vision3.EntityAnnotation `json:"labels"`
}

// ProcessUpload prints a message when a file is changed in a Cloud Storage bucket.
func ProcessUpload(ctx context.Context, e GCSEvent) error {
	log.Printf("Processing file: %s", e.Name)
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

	gcs, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	defer gcs.Close()

	object := gcs.Bucket(e.Bucket).Object(e.Name)

	rc, err := object.NewReader(ctx)
	defer rc.Close()

	// Creates a client.
	labeler, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		return err
	}
	defer labeler.Close()

	image, err := vision.NewImageFromReader(rc)
	if err != nil {
		return err
	}

	labels, err := labeler.DetectLabels(ctx, image, nil, 10)
	if err != nil {
		return err
	}

	firestore, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return err
	}

	uploads := firestore.Collection("Uploads")
	doc, _, err := uploads.Add(ctx, Upload{
		Event: e,
		Labels:  labels,
	})
	if err != nil {
		return err
	}

	log.Printf("Created Firestore document: %s", doc.ID)

	return nil
}
