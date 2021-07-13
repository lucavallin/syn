// Package functions contains Google Cloud Functions
package functions

import (
	"cavall.in/syn/syn"
	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
	vision "cloud.google.com/go/vision/apiv1"
	"context"
	"github.com/thoas/go-funk"
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
	acceptedLabels := syn.CleanLabels(os.Getenv("ACCEPTED_LABELS"))
	if len(acceptedLabels) == 0 {
		log.Printf("Deleting upload: No ACCEPTED_LABELS provided")
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

	detectedLabels, err := labeler.DetectLabels(ctx, image, nil, 10)
	if err != nil {
		return err
	}

	labels := funk.Map(detectedLabels, func(l *vision3.EntityAnnotation) syn.Label {
		return syn.Label{Description: l.Description, Score: l.Score}
	})

	// Reject images that don't contain the allowed labels
	allowed := funk.Contains(labels, func(l syn.Label) bool {
		return funk.Contains(acceptedLabels, l.Description)
	})
	if !allowed {
		log.Printf("Deleting upload: no allowed labels detected")
		log.Printf("Wanted: %v", acceptedLabels)
		log.Printf("Found: %v", labels)
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
	doc, _, err := uploads.Add(ctx, syn.Upload{
		File: syn.File{
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
