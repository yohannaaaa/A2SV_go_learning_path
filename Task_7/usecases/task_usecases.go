package usecases

import (
	domain "task-manager/Domain"
)

type TaskUsecase struct {
	taskRepo domain.TaskRepository
}

func NewTaskUsecase(r domain.TaskRepository) *TaskUsecase {
	return &TaskUsecase{taskRepo: r}
}

func (t *TaskUsecase) GetAll() ([]domain.Task, error) {
	return t.taskRepo.GetAll()
}

func (t *TaskUsecase) GetByID(id string) (*domain.Task, error) {
	return t.taskRepo.GetByID(id)
}

func (t *TaskUsecase) Create(task *domain.Task) error {
	return t.taskRepo.Create(task)
}

func (t *TaskUsecase) Update(id string, task *domain.Task) error {
	return t.taskRepo.Update(id, task)
}

func (t *TaskUsecase) Delete(id string) error {
	return t.taskRepo.Delete(id)
}
