package syn

import (
	"time"
)

// File contains object information for an image uploaded by Raspberry Pi
type File struct {
	Bucket string `json:"bucket" firestore:"bucket"`
	Name   string `json:"name" firestore:"name"`
}

// Upload represents the data stored into Firestore, which includes file information,
// labels from Vision API and a timestamp
type Upload struct {
	File    `json:"file" firestore:"file"`
	Created time.Time `json:"created" firestore:"created"`
	Labels  []Label   `json:"labels" firestore:"labels"`
}

func NewUpload(bucket string, name string, created time.Time, labels []Label) *Upload {
	return &Upload{
		File: File{
			Bucket: bucket,
			Name:   name,
		},
		Created: created,
		Labels:  labels,
	}
}
