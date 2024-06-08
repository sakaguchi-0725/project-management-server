//go:generate mockgen -source=$GOFILE -destination=../test/mocks/$GOPACKAGE/mock_$GOFILE -package=mocks
package usecase

import (
	"github.com/sakaguchi-0725/go-todo/internal/domain"
	"github.com/sakaguchi-0725/go-todo/internal/infrastructure/repository"
	"github.com/sakaguchi-0725/go-todo/internal/usecase/input"
)

type TaskUsecase interface {
	GetAllTasks() ([]domain.Task, error)
	GetTaskById(taskId uint) (*domain.Task, error)
	CreateTask(input input.TaskInput) error
	UpdateTask(domain.Task) (*domain.Task, error)
	DeleteTask(taskId uint) error
}

type taskUsecase struct {
	tr repository.TaskRepository
}

// CreateTask implements TaskUsecase.
func (t *taskUsecase) CreateTask(input input.TaskInput) error {
	task := domain.Task{
		Title: input.Title,
		Desc:  input.Desc,
	}

	if err := t.tr.CreateTask(task); err != nil {
		return err
	}

	return nil
}

// DeleteTask implements TaskUsecase.
func (t *taskUsecase) DeleteTask(taskId uint) error {
	panic("unimplemented")
}

// GetAllTask implements TaskUsecase.
func (t *taskUsecase) GetAllTasks() ([]domain.Task, error) {
	tasks, err := t.tr.GetAllTasks()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

// GetTaskById implements TaskUsecase.
func (t *taskUsecase) GetTaskById(taskId uint) (*domain.Task, error) {
	panic("unimplemented")
}

// UpdateTask implements TaskUsecase.
func (t *taskUsecase) UpdateTask(domain.Task) (*domain.Task, error) {
	panic("unimplemented")
}

func NewTaskUsecase(tr repository.TaskRepository) TaskUsecase {
	return &taskUsecase{tr}
}
