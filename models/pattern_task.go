package models

type PatternTask struct {
	Model
	PatternText   string `json:"pattern_text"`
	EgeTaskNumber uint8  `json:"ege_task_number"`
}

type PatternTasks []PatternTask
