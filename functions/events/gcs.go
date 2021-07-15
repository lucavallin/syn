package events

// GCSEvent is the payload of a GCS event
type GCSEvent struct {
	Bucket string `json:"bucket"`
	Name   string `json:"name"`
}
