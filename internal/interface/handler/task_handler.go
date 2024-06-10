//go:generate mockgen -source=$GOFILE -destination=../../test/mocks/$GOPACKAGE/mock_$GOFILE -package=mocks
package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sakaguchi-0725/go-todo/internal/interface/dto"
	"github.com/sakaguchi-0725/go-todo/internal/usecase"
	"github.com/sakaguchi-0725/go-todo/internal/usecase/input"
	"github.com/sakaguchi-0725/go-todo/pkg/apperr"
)

type TaskHandler interface {
	GetAllTasks(c echo.Context) error
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
	var req dto.TaskRequest
	if err := c.Bind(&req); err != nil {
		return apperr.NewAppError(apperr.ErrBadRequest, apperr.ErrCategoryCreateTaskInvalidParameter, err.Error())
	}

	input := input.TaskInput{
		Title: req.Title,
		Desc:  req.Desc,
	}
	if err := t.tu.CreateTask(input); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, nil)
}

// DeleteTask implements TaskHandler.
func (t *taskHandler) DeleteTask(c echo.Context) error {
	id := c.Param("taskId")
	taskId, err := strconv.Atoi(id)
	if err != nil {
		return apperr.NewAppError(apperr.ErrBadRequest, apperr.ErrCategoryDeleteTaskInvalidParameter, err.Error())
	}

	if err := t.tu.DeleteTask(uint(taskId)); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

// GetAllTask implements TaskHandler.
func (t *taskHandler) GetAllTasks(c echo.Context) error {
	tasks, err := t.tu.GetAllTasks()
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
	var req dto.TaskRequest
	if err := c.Bind(&req); err != nil {
		return apperr.NewAppError(apperr.ErrBadRequest, apperr.ErrCategoryUpdateTaskInvalidParameter, err.Error())
	}

	input := input.TaskInput{
		ID:    *req.ID,
		Title: req.Title,
		Desc:  req.Desc,
	}

	err := t.tu.UpdateTask(input)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, nil)
}

func NewTaskHandler(tu usecase.TaskUsecase) TaskHandler {
	return &taskHandler{tu}
}
