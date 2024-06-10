//go:generate mockgen -source=$GOFILE -destination=../../test/mocks/$GOPACKAGE/mock_$GOFILE -package=mocks
package repository

import (
	"github.com/sakaguchi-0725/go-todo/internal/domain"
	"gorm.io/gorm"
)

type TaskRepository interface {
	GetAllTasks() ([]domain.Task, error)
	GetTaskById(taksId uint) (*domain.Task, error)
	CreateTask(task domain.Task) error
	UpdateTask(task domain.Task) error
	DeleteTask(taskId uint) error
}

type taskRepository struct {
	db *gorm.DB
}

// CreateTask implements TaskRepository.
func (t *taskRepository) CreateTask(task domain.Task) error {
	if err := t.db.Create(&task).Error; err != nil {
		return err
	}

	return nil
}

// DeleteTask implements TaskRepository.
func (t *taskRepository) DeleteTask(taskId uint) error {
	if err := t.db.Delete(&domain.Task{}, taskId).Error; err != nil {
		return err
	}
	return nil
}

// GetAllTask implements TaskRepository.
func (t *taskRepository) GetAllTasks() ([]domain.Task, error) {
	var tasks []domain.Task
	if err := t.db.Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

// GetTaskById implements TaskRepository.
func (t *taskRepository) GetTaskById(taskId uint) (*domain.Task, error) {
	panic("unimplemented")
}

// UpdateTask implements TaskRepository.
func (t *taskRepository) UpdateTask(task domain.Task) error {
	if err := t.db.Save(&task).Error; err != nil {
		return err
	}
	return nil
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db}
}
