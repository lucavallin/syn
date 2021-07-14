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
	Labels Labels `json:"labels"`
}

type Labels struct {
	ArrayValue LabelArrayValue `json:"arrayValue"`
}

type LabelArrayValue struct {
	Values []LabelValues `json:"values"`
}

type LabelValues struct {
	MapValue LabelsMapValue `json:"mapValue"`
}

type LabelsMapValue struct {
	Fields LabelFields `json:"fields"`
}

type LabelFields struct {
	Description LabelDescription `json:"description"`
	Score LabelScore `json:"score"`
}

type LabelDescription struct {
	StringValue string `json:"stringValue"`
}

type LabelScore struct {
	DoubleValue float64 `json:"doubleValue"`
}