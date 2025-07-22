package controllers

import (
	"net/http"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context){
	tasks, err := data.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func GetTask(c * gin.Context) {
	id := c.Param("id")
	task, err := data.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func CreateTask(c *gin.Context) {
	var newTask models.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	if err := data.CreateTask(&newTask); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}
	c.JSON(http.StatusCreated, newTask)
} 

func DeleteTask(c *gin.Context){
	id := c.Param("id")
	err := data.DeleteTask(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var updatedTask models.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	} 
	err := data.UpdateTask(id, &updatedTask)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, updatedTask)
}
