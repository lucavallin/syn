package functions

import (
	"cavall.in/syn/events"
	"cavall.in/syn/syn"
	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
	vision "cloud.google.com/go/vision/apiv1"
	"context"
	"github.com/h2non/filetype"
	"github.com/thoas/go-funk"
	vision3 "google.golang.org/genproto/googleapis/cloud/vision/v1"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// ProcessUpload labels images that have been uploaded to Cloud Storage
// and saves the data into Firestore when the labels we want are found in the image
func ProcessUpload(ctx context.Context, e events.GCSEvent) error {
	log.Printf("Processing upload: %s", e.Name)
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT_ID")
	projectNumber := os.Getenv("GOOGLE_CLOUD_PROJECT_NUMBER")

	//
	// Get object from Cloud Storage
	//
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

	upload, err := ioutil.ReadAll(rc)
	if err != nil {
		return err
	}

	if !filetype.IsImage(upload) {
		log.Printf("Deleting upload: not an image")
		return nil
	}

	//
	// Clean allowed values and fail if none are provided
	// Uploads are stored to Firestore only if Vision API finds at least one of the labels
	//
	acceptedLabels := cleanLabels(os.Getenv("ACCEPTED_LABELS"))
	if len(acceptedLabels) == 0 {
		log.Printf("Deleting upload: No ACCEPTED_LABELS provided")
		if err := object.Delete(ctx); err != nil {
			log.Printf("Failed to delete upload: %s", e.Name)
			return err
		}
		return nil
	}

	//
	// Query Vision API
	//
	labeler, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		return err
	}
	defer labeler.Close()

	image, err := vision.NewImageFromReader(rc)
	if err != nil {
		return err
	}

	detectedLabels, err := labeler.DetectLabels(ctx, image, nil, 5)
	if err != nil {
		return err
	}

	//
	// Transform resulting labels and check if they contain at least one of the allowed labels
	//
	labels := funk.Map(detectedLabels, func(l *vision3.EntityAnnotation) syn.Label {
		return syn.NewLabel(l.Description, l.Score)
	}).([]syn.Label)

	allowed := funk.Contains(labels, func(l syn.Label) bool {
		return -1 != funk.IndexOf(acceptedLabels, l.Description)
	})
	if !allowed {
		log.Printf("Allowed: %v", acceptedLabels)
		log.Printf("Found: %v", labels)
		log.Printf("Deleting upload: no allowed labels detected")
		if err := object.Delete(ctx); err != nil {
			log.Printf("Failed to delete upload: %s", e.Name)
			return err
		}
		return nil
	}

	//
	// Store file (path) and labels to Firestore
	//
	firestore, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return err
	}

	uploads := firestore.Collection("Uploads")
	doc, _, err := uploads.Add(ctx, syn.Upload{
		File: syn.File{
			Bucket: object.BucketName(),
			Name:   object.ObjectName(),
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

// cleanLabels trims the comma-separated allowed labels, makes them lowercase and splits them into a slice
func cleanLabels(labels string) []string {
	lowerLabels := strings.ToLower(labels)
	labelsWithoutWhitespaces := strings.ReplaceAll(lowerLabels, " ", "")

	return strings.Split(labelsWithoutWhitespaces, ",")
}
