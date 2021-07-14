package syn

import (
	"time"
)

type File struct {
	Bucket string `json:"bucket" firestore:"bucket"`
	Name string `json:"name" firestore:"name"`
}

type Upload struct {
	File `json:"file" firestore:"file"`
	Labels []Label `json:"labels" firestore:"labels"`
	Created time.Time `json:"created" firestore:"created"`
}
