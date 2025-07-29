package controllers

import (
	"net/http"
	"task-manager/Domain"
	"task-manager/Usecases"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskUsecase *usecases.TaskUsecase
}

func NewTaskController(t *usecases.TaskUsecase) *TaskController {
	return &TaskController{taskUsecase: t}
}

type TaskDTO struct {
	Title   string `json:"title" binding:"required"`
	DueDate string `json:"due_date"`
	Status  string `json:"status"`
	UserID  string `json:"user_id"`
}

func (tc *TaskController) GetAll(c *gin.Context) {
	tasks, err := tc.taskUsecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (tc *TaskController) GetByID(c *gin.Context) {
	id := c.Param("id")
	task, err := tc.taskUsecase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (tc *TaskController) Create(c *gin.Context) {
	var dto TaskDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task := &domain.Task{
		Title:   dto.Title,
		DueDate: dto.DueDate,
		Status:  dto.Status,
		UserID:  dto.UserID,
	}

	if err := tc.taskUsecase.Create(task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Task created successfully"})
}

func (tc *TaskController) Update(c *gin.Context) {
	id := c.Param("id")

	var dto TaskDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task := &domain.Task{
		Title:   dto.Title,
		DueDate: dto.DueDate,
		Status:  dto.Status,
		UserID:  dto.UserID,
	}

	if err := tc.taskUsecase.Update(id, task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}

func (tc *TaskController) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := tc.taskUsecase.Delete(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
