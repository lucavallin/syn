package events

import "time"

// FirestoreUpload represents a Firebase event of a new record in the Upload collection
type FirestoreUpload struct {
	Created Created `json:"created"`
	File File `json:"file"`
	Labels Labels `json:"labels"`
}

type Created struct {
	TimestampValue time.Time `json:"timestampValue"`
}

type File struct {
	MapValue FileMapValue `json:"mapValue"`
}

type FileMapValue struct {
	Fields FileFields `json:"fields"`
}

type FileFields struct {
	Bucket StringValue `json:"bucket"`
	Name StringValue `json:"name"`
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
	Description StringValue `json:"description"`
	Score DoubleValue `json:"score"`
}

type StringValue struct {
	StringValue string `json:"stringValue"`
}

type DoubleValue struct {
	DoubleValue float64 `json:"doubleValue"`
}