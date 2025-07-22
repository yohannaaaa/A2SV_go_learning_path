package models

import "time"

type Status string

const(
	StatusNotStarted	Status = "not_started"
	StatusInProgress	Status = "in_progress"
	StatusCompleted		Status = "completed"
)
type Task struct {
	ID          string 			`json:"id" bson:"_id,omitempty"`
	Title       string 			`json:"title" bson:"title"`
	DueDate     *time.Time 		`json:"due_date,omitempty" bson:"due_date"`
	Status      *Status 		`json:"status,omitempty" bson:"status"`
} 