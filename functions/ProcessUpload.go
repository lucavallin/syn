// Package functions contains Google Cloud Functions
package functions

import (
	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
	vision "cloud.google.com/go/vision/apiv1"
	"context"
	"github.com/thoas/go-funk"
	vision3 "google.golang.org/genproto/googleapis/cloud/vision/v1"
	"log"
	"os"
	"strings"
	"time"
)

// GCSEvent is the payload of a GCS event. Please refer to the docs for
// additional information regarding GCS events.
type GCSEvent struct {
	Bucket string `json:"bucket" firestore`
	Name   string `json:"name"`
}

type File struct {
	Bucket string `json:"bucket" firestore:"bucket"`
	Name string `json:"name" firestore:"name"`
}

type Upload struct {
	File `json:"file" firestore:"file"`
	Labels []*vision3.EntityAnnotation `json:"labels" firestore:"labels"`
	Created time.Time `json:"created" firestore:"created"`
}

// ProcessUpload prints a message when a file is changed in a Cloud Storage bucket.
func ProcessUpload(ctx context.Context, e GCSEvent) error {
	log.Printf("Processing upload: %s", e.Name)
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

	gcs, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	defer gcs.Close()

	object := gcs.Bucket(e.Bucket).Object(e.Name)
	objectAttrs, err := object.Attrs(ctx)
	if err != nil {
		return err
	}

	rc, err := object.NewReader(ctx)
	if err != nil {
		return err
	}
	defer rc.Close()

	// Uploads are stored to Firestore only if Vision API returns at least one of these labels (comma-separated)
	acceptedLabels := strings.Split(strings.ToLower(os.Getenv("ACCEPTED_LABELS")), ",")
	if len(acceptedLabels) == 0 {
		log.Printf("Upload rejected: No ACCEPTED_LABELS provided")
		if err := object.Delete(ctx); err != nil {
			log.Printf("Failed to delete upload: %s", e.Name)
			return err
		}
		return nil
	}

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

	// Reject images that don't contain the allowed labels
	allowed := funk.Contains(labels, func(l *vision3.EntityAnnotation) bool {
		return funk.Contains(acceptedLabels, l.Description)
	})
	if !allowed {
		log.Printf("Upload deleted: no allowed labels detected")
		if err := object.Delete(ctx); err != nil {
			log.Printf("Failed to delete upload: %s", e.Name)
			return err
		}
		return nil
	}

	firestore, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return err
	}

	uploads := firestore.Collection("Uploads")
	doc, _, err := uploads.Add(ctx, Upload{
		File: File{
			Bucket: object.BucketName(),
			Name: object.ObjectName(),
		},
		Labels:  labels,
		Created: objectAttrs.Created,
	})
	if err != nil {
		return err
	}

	log.Printf("Created Firestore document: %s", doc.ID)

	return nil
}
