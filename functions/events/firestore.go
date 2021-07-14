package events

import "time"

// FirestoreUpload represents a Firebase event of a new record in the Upload collection
type FirestoreUpload struct {
	Created struct {
		TimestampValue time.Time `json:"timestampValue"`
	} `json:"created"`
	File struct {
		MapValue struct {
			Fields struct {
				Bucket struct {
					StringValue string `json:"stringValue"`
				} `json:"bucket"`
				Name struct {
					StringValue string `json:"stringValue"`
				} `json:"name"`
			} `json:"fields"`
		} `json:"mapValue"`
	} `json:"file"`
	Labels []Label `json:"labels"`
}

type Label struct {
	ArrayValue struct {
		Values []struct {
			MapValue struct {
				Fields struct {
					Description struct {
						StringValue string `json:"stringValue"`
					} `json:"description"`
					Score struct {
						DoubleValue float64 `json:"doubleValue"`
					} `json:"score"`
				} `json:"fields"`
			} `json:"mapValue"`
		} `json:"values"`
	} `json:"arrayValue"`
}
