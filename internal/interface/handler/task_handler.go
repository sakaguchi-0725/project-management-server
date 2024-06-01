package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sakaguchi-0725/go-todo/internal/interface/dto"
	"github.com/sakaguchi-0725/go-todo/internal/usecase"
)

type TaskHandler interface {
	GetAllTask(c echo.Context) error
	GetTaskById(c echo.Context) error
	CreateTask(c echo.Context) error
	UpdateTask(c echo.Context) error
	DeleteTask(c echo.Context) error
}

type taskHandler struct {
	tu usecase.TaskUsecase
}

// CreateTask implements TaskHandler.
func (t *taskHandler) CreateTask(c echo.Context) error {
	panic("unimplemented")
}

// DeleteTask implements TaskHandler.
func (t *taskHandler) DeleteTask(c echo.Context) error {
	panic("unimplemented")
}

// GetAllTask implements TaskHandler.
func (t *taskHandler) GetAllTask(c echo.Context) error {
	tasks, err := t.tu.GetAllTask()
	if err != nil {
		return err
	}

	res := make([]dto.TaskResponse, len(tasks))
	for i, task := range tasks {
		res[i] = dto.TaskResponse{
			ID:        task.ID,
			Title:     task.Title,
			Desc:      task.Desc,
			CreatedAt: task.CreatedAt.Format(time.RFC3339),
			UpdatedAt: task.UpdatedAt.Format(time.RFC3339),
		}
	}

	return c.JSON(http.StatusOK, res)
}

// GetTaskById implements TaskHandler.
func (t *taskHandler) GetTaskById(c echo.Context) error {
	panic("unimplemented")
}

// UpdateTask implements TaskHandler.
func (t *taskHandler) UpdateTask(c echo.Context) error {
	panic("unimplemented")
}

func NewTaskHandler(tu usecase.TaskUsecase) TaskHandler {
	return &taskHandler{tu}
}
