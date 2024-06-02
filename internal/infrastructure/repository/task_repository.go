//go:generate mockgen -source=$GOFILE -destination=../../test/mocks/$GOPACKAGE/mock_$GOFILE -package=mocks
package repository

import (
	"github.com/sakaguchi-0725/go-todo/internal/domain"
	"gorm.io/gorm"
)

type TaskRepository interface {
	GetAllTask() ([]domain.Task, error)
	GetTaskById(taksId uint) (*domain.Task, error)
	CreateTask(domain.Task) error
	UpdateTask(domain.Task) (*domain.Task, error)
	DeleteTask(taskId uint) error
}

type taskRepository struct {
	db *gorm.DB
}

// CreateTask implements TaskRepository.
func (t *taskRepository) CreateTask(domain.Task) error {
	panic("unimplemented")
}

// DeleteTask implements TaskRepository.
func (t *taskRepository) DeleteTask(taskId uint) error {
	panic("unimplemented")
}

// GetAllTask implements TaskRepository.
func (t *taskRepository) GetAllTask() ([]domain.Task, error) {
	var tasks []domain.Task
	if err := t.db.Find(&tasks); err.Error != nil {
		return nil, err.Error
	}

	return tasks, nil
}

// GetTaskById implements TaskRepository.
func (t *taskRepository) GetTaskById(taskId uint) (*domain.Task, error) {
	panic("unimplemented")
}

// UpdateTask implements TaskRepository.
func (t *taskRepository) UpdateTask(domain.Task) (*domain.Task, error) {
	panic("unimplemented")
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db}
}
