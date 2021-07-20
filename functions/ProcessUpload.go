package functions

import (
	"cavall.in/syn/database"
	"cavall.in/syn/events"
	"cavall.in/syn/gcs"
	"cavall.in/syn/syn"
	"cavall.in/syn/visionapi"
	"context"
	"github.com/thoas/go-funk"
	"log"
	"os"
	"strings"
)

// ProcessUpload labels images that have been uploaded to Cloud Storage
// and saves the data into Firestore when the labels we want are found in the image
func ProcessUpload(ctx context.Context, e events.GCSEvent) error {
	log.Printf("Processing upload: %s", e.Name)
	projectId := os.Getenv("GOOGLE_CLOUD_PROJECT_ID")

	//
	// Get object from Cloud Storage
	//
	gcsClient, err := gcs.NewClient(ctx)
	if err != nil {
		return err
	}
	defer gcsClient.Connection.Close()

	object, err := gcsClient.GetObject(e.Bucket, e.Name)
	if err != nil {
		return err
	}

	//
	// Sanitize allowed labels and fail if none are provided
	// Uploads are stored to Firestore only if Vision API finds at least one of the labels
	//
	acceptedLabelsEnv := strings.Split(os.Getenv("ACCEPTED_LABELS"), ",")
	acceptedLabels := funk.Map(acceptedLabelsEnv, func(l string) string {
		return strings.ToLower(strings.TrimSpace(l))
	}).([]string)

	if len(acceptedLabels) == 0 {
		log.Printf("Deleting upload: No ACCEPTED_LABELS provided")
		return gcsClient.Delete(object)
	}

	//
	// Query Vision API
	//
	vision, err := visionapi.NewClient(ctx)
	if err != nil {
		return err
	}
	defer vision.Connection.Close()

	labels, err := vision.DetectImageLabels(object.URI)
	if err != nil {
		return err
	}

	var isAllowedLabel = func(l syn.Label) bool { return -1 != funk.IndexOf(acceptedLabels, l.Description) }
	if !funk.Contains(labels, isAllowedLabel) {
		log.Printf("Deleting upload: no allowed labels detected. Allowed: %v, Found: %v", acceptedLabels, labels)
		return gcsClient.Delete(object)
	}

	//
	// Store labeled uploads to Firestore
	//
	db, err := database.NewClient(ctx, projectId)
	if err != nil {
		return err
	}
	defer db.Connection.Close()

	data := syn.NewEvent(object.URI, object.Created, labels)
	docId, err := db.AddEvent(data)
	if err != nil {
		return err
	}

	log.Printf("Created Firestore document: %s", docId)

	return nil
}
