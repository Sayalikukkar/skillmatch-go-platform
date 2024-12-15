package models

type TaskProgress struct {
	ID          int    `json:"id"`
	TaskID      int    `json:"task_id"`
	Description string `json:"description"`
	Timestamp   string `json:"timestamp"` // RFC3339 format
}
