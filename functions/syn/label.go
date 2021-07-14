package syn

import "strings"

type Label struct {
	Description string `json:"description" firestore:"description"`
	Score float32 `json:"score" firestore:"score"`
}

func NewLabel(description string, score float32) Label {
	return Label{strings.ToLower(description), score}
}
