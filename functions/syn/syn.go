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
	Labels  []Label   `json:"labels" firestore:"labels"`
	Created time.Time `json:"created" firestore:"created"`
}

func NewUpload(bucket string, name string, labels []Label, created time.Time) *Upload {
	return &Upload{
		File: File{
			Bucket: bucket,
			Name:   name,
		},
		Labels:  labels,
		Created: created,
	}
}
