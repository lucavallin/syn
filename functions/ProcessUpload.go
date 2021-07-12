// Package p contains Google Cloud Functions
package p

import (
	"cloud.google.com/go/storage"
	vision "cloud.google.com/go/vision/apiv1"
	"context"
	"fmt"
	"log"
)

// GCSEvent is the payload of a GCS event. Please refer to the docs for
// additional information regarding GCS events.
type GCSEvent struct {
	Bucket string `json:"bucket"`
	Name   string `json:"name"`
}

// ProcessUpload prints a message when a file is changed in a Cloud Storage bucket.
func ProcessUpload(ctx context.Context, e GCSEvent) error {
	log.Printf("Processing file: %s", e.Name)

	gcs, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	defer gcs.Close()

	rc, err := gcs.Bucket(e.Bucket).Object(e.Name).NewReader(ctx)
	if err != nil {
		return err
	}
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

	fmt.Println("Labels:")
	for _, label := range labels {
		fmt.Println(label.Description)
	}

	return nil
}
