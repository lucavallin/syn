package events

import (
	"github.com/thoas/go-funk"
	"time"
)

//FirestoreEvent is the payload of a Firestore event
type FirestoreEvent struct {
	OldValue   FirestoreValue `json:"oldValue"`
	Value      FirestoreValue `json:"value"`
	UpdateMask struct {
		FieldPaths []string `json:"fieldPaths"`
	} `json:"updateMask"`
}

// FirestoreValue holds Firestore fields
type FirestoreValue struct {
	CreateTime time.Time       `json:"createTime"`
	Fields     FirestoreUpload `json:"fields"`
	Name       string          `json:"name"`
	UpdateTime time.Time       `json:"updateTime"`
}

// FirestoreUpload represents a Firebase event of a new record in the Upload collection
type FirestoreUpload struct {
	Created Created `json:"created"`
	File    File    `json:"file"`
	Labels  Labels  `json:"labels"`
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
	Name   StringValue `json:"name"`
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
	Score       DoubleValue `json:"score"`
}

type StringValue struct {
	StringValue string `json:"stringValue"`
}

type DoubleValue struct {
	DoubleValue float64 `json:"doubleValue"`
}

// GetUploadLabels returns the labels of the image as an array of strings
func (e FirestoreEvent) GetUploadLabels() []string {
	return funk.Map(e.Value.Fields.Labels.ArrayValue.Values, func(l LabelValues) string {
		return l.MapValue.Fields.Description.StringValue
	}).([]string)
}
