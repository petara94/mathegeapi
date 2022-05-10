package models

type PatternTask struct {
	Model
	PatternText   string `json:"pattern_text"`
	EgeTaskNumber uint8  `json:"ege_task_number"`
}

func (p *PatternTask) AllowedToUpdate() map[string]interface{} {
	return map[string]interface{}{}
}

type PatternTasks []PatternTask
