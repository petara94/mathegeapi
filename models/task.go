package models

type Task struct {
	Model
	MathEgeID     uint   `gorm:"unique;default:null" json:"math_ege_id" entity:"protected"`
	TaskText      string `json:"task_text"`
	PatternTaskID *uint  `gorm:"default:null" json:"pattern_task_id"`

	PatternTask *PatternTask `json:"pattern_task"`
	Images      TaskImages   `json:"images"`
}

type Tasks []Task
