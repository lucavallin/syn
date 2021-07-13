package syn

import (
	"strings"
	"time"
)

type File struct {
	Bucket string `json:"bucket" firestore:"bucket"`
	Name string `json:"name" firestore:"name"`
}

type Label struct {
	Description string `json:"description" firestore:"description"`
	Score float32 `json:"score" firestore:"score"`
}

type Upload struct {
	File `json:"file" firestore:"file"`
	Labels []Label `json:"labels" firestore:"labels"`
	Created time.Time `json:"created" firestore:"created"`
}

func NewLabel(description string, score float32) Label {
	return Label{strings.ToLower(description), score}
}