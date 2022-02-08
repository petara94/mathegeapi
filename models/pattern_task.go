package models

import (
	"gorm.io/gorm"
)

type PatternTask struct {
	Model
	PatternText   string `json:"pattern_text"`
	EgeTaskNumber int    `json:"ege_task_number"`
}

func (p *PatternTask) BeforeUpdate(*gorm.DB) error {
	p.Model = Model{}
	return nil
}

func (p *PatternTask) Allowed() map[string]interface{} {
	return map[string]interface{}{
		"pattern_text":    p.PatternText,
		"ege_task_number": p.EgeTaskNumber,
	}
}

type PatternTasks []PatternTask
