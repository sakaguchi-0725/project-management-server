package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/sakaguchi-0725/go-todo/internal/domain"
	"github.com/sakaguchi-0725/go-todo/internal/interface/dto"
	"github.com/sakaguchi-0725/go-todo/internal/interface/handler"
	mocks "github.com/sakaguchi-0725/go-todo/internal/test/mocks/usecase"
)

var (
	mockUsecase *mocks.MockTaskUsecase
	taskHandler handler.TaskHandler
)

func TestMain(m *testing.M) {
	ctrl := gomock.NewController(&testing.T{})
	mockUsecase = mocks.NewMockTaskUsecase(ctrl)
	taskHandler = handler.NewTaskHandler(mockUsecase)

	status := m.Run()
	os.Exit(status)
}

func TestGetAllTasks(t *testing.T) {
	res := []domain.Task{
		{ID: 1, Title: "テスト001"},
		{ID: 2, Title: "テスト002"},
		{ID: 3, Title: "テスト003"},
	}

	t.Run("正常系", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockUsecase.EXPECT().GetAllTasks().Return(res, nil).Times(1)

		err := taskHandler.GetAllTasks(c)
		if err != nil {
			t.Errorf("GetAllTasks failed: expected no error, got %v", err)
		}

		if rec.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, rec.Code)
		}
	})

	t.Run("異常系", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockUsecase.EXPECT().GetAllTasks().Return(nil, errors.New("error")).Times(1)

		err := taskHandler.GetAllTasks(c)
		if err == nil {
			t.Errorf("Expected an error from GetAllTasks, but received none")
		}
	})
}

func TestCreateTask(t *testing.T) {
	reqBody, err := json.Marshal(
		dto.TaskRequest{
			Title: "Createテスト",
			Desc:  "テスト",
		},
	)
	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	t.Run("正常系", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockUsecase.EXPECT().CreateTask(gomock.Any()).Return(nil).Times(1)

		err := taskHandler.CreateTask(c)
		if err != nil {
			t.Errorf("Expect no errors, got %v", err)
		}

		if rec.Code != http.StatusCreated {
			t.Errorf("Expected status code %d, got %d", http.StatusCreated, rec.Code)
		}
	})

	t.Run("異常系", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockUsecase.EXPECT().CreateTask(gomock.Any()).
			Return(errors.New("作成に失敗しました")).Times(1)

		err := taskHandler.CreateTask(c)
		if err == nil {
			t.Errorf("Expected an error from GetAllTasks, but received none")
		}
	})
}
