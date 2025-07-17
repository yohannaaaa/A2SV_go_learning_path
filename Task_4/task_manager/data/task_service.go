package data

import (
	"errors"
	"task_manager/models"
	"time"
)

var tasks = []models.Task{
	{ID: "1", Title: "Task 1", Description: "First task", DueDate: timePtr(time.Now().AddDate(0, 0, 7)), Status: strPtr(models.StatusCompleted)},
	{ID: "2", Title: "Task 2", Description: "Second task", DueDate: timePtr(time.Now().AddDate(0, 0, 1)), Status: strPtr(models.StatusInProgress)},
	{ID: "3", Title: "Task 3", Description: "Third task", DueDate: timePtr(time.Now().AddDate(0, 0, 7)), Status: strPtr(models.StatusNotStarted)},
}

func timePtr(t time.Time) *time.Time {
	return &t
}

func strPtr(s models.Status) * models.Status {
	return &s
}

func GetAllTasks() []models.Task {
	return tasks
}

func GetTaskByID (id string) (*models.Task, error) {
	for _, task := range tasks {
		if task.ID == id{
			return &task, nil
		}
	}
	return nil, errors.New("task not found")
}

func CreateTask(task models.Task) {	
	tasks = append(tasks, task)
}

func DeleteTask(id string) error {
	for i, task := range tasks{
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i + 1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}

func UpdateTask(id string, updatedTask models.Task) error {
	for i, task := range tasks {
		if task.ID == id{
			if updatedTask.Title != "" {
				tasks[i].Title = updatedTask.Title 
			}
			if updatedTask.DueDate != nil {
				tasks[i].DueDate = updatedTask.DueDate
			}
			if updatedTask.Status != nil {
				tasks[i].Status = updatedTask.Status 
			}
			return nil
		}
	}
	return errors.New("task not found")
}