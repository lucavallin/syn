package syn

import (
	"time"
)

// Event represents the data stored into Firestore, which includes file information,
// labels from Vision API and a timestamp
type Event struct {
	URI     string    `json:"uri" firestore:"uri"`
	Created time.Time `json:"created" firestore:"created"`
	Labels  []Label   `json:"labels" firestore:"labels"`
}

func NewEvent(uri string, created time.Time, labels []Label) *Event {
	return &Event{
		URI:     uri,
		Created: created,
		Labels:  labels,
	}
}
