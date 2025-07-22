package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Status string

const(
	StatusNotStarted	Status = "not_started"
	StatusInProgress	Status = "in_progress"
	StatusCompleted		Status = "completed"
)
type Task struct {
	ID          primitive.ObjectID			`json:"id" bson:"_id,omitempty"`
	Title       string 			`json:"title" bson:"title"`
	DueDate     *time.Time 		`json:"due_date,omitempty" bson:"due_date"`
	Status      *Status 		`json:"status,omitempty" bson:"status"`
} 