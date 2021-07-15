package functions

import (
	"cavall.in/syn/database"
	"cavall.in/syn/events"
	"cavall.in/syn/syn"
	"cavall.in/syn/visionapi"
	"cloud.google.com/go/storage"
	"context"
	"github.com/h2non/filetype"
	"github.com/thoas/go-funk"
	"io/ioutil"
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
	// Sanitize allowed labels and fail if none are provided
	// Uploads are stored to Firestore only if Vision API finds at least one of the labels
	//
	acceptedLabelsEnv := strings.Split(os.Getenv("ACCEPTED_LABELS"), ",")
	acceptedLabels := funk.Map(acceptedLabelsEnv, func(l string) string {
		return strings.ToLower(strings.TrimSpace(l))
	}).([]string)
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
	visionApi, err := visionapi.NewClient(ctx)
	if err != nil {
		return err
	}

	labels, err := visionApi.DetectImageLabels(rc)
	if err != nil {
		return err
	}

	var isAllowedLabel = func(l syn.Label) bool { return -1 != funk.IndexOf(acceptedLabels, l.Description) }
	if !funk.Contains(labels, isAllowedLabel) {
		log.Printf("Deleting upload: no allowed labels detected. Allowed: %v, Found: %v", acceptedLabels, labels)
		if err := object.Delete(ctx); err != nil {
			log.Printf("Failed to delete upload: %s", e.Name)
			return err
		}
		return nil
	}

	//
	// Store labeled uploads to Firestore
	//
	db, err := database.NewClient(ctx, projectId)
	if err != nil {
		return err
	}

	data := syn.NewUpload(object.BucketName(), object.ObjectName(), labels, objectAttrs.Created)
	docId, err := db.AddUpload(data)
	if err != nil {
		return err
	}

	log.Printf("Created Firestore document: %s", docId)

	return nil
}
