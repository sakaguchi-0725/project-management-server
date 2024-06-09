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
	"github.com/sakaguchi-0725/go-todo/cmd/server"
	"github.com/sakaguchi-0725/go-todo/internal/domain"
	"github.com/sakaguchi-0725/go-todo/internal/interface/dto"
	"github.com/sakaguchi-0725/go-todo/internal/interface/handler"
	mocks "github.com/sakaguchi-0725/go-todo/internal/test/mocks/usecase"
	testutil "github.com/sakaguchi-0725/go-todo/internal/test/util"
)

var (
	mockUsecase *mocks.MockTaskUsecase
	taskHandler handler.TaskHandler
	e           *echo.Echo
)

func TestMain(m *testing.M) {
	ctrl := gomock.NewController(&testing.T{})
	mockUsecase = mocks.NewMockTaskUsecase(ctrl)
	taskHandler = handler.NewTaskHandler(mockUsecase)
	handlers := server.AppHandlers{
		TaskHandler: taskHandler,
	}
	e = server.NewServer(handlers)

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
		req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
		rec := httptest.NewRecorder()
		mockUsecase.EXPECT().GetAllTasks().Return(res, nil).Times(1)
		e.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, rec.Code)
		}
	})

	t.Run("異常系", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
		rec := httptest.NewRecorder()
		mockUsecase.EXPECT().GetAllTasks().Return(nil, errors.New("error")).Times(1)
		e.ServeHTTP(rec, req)

		if rec.Code != http.StatusInternalServerError {
			t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, rec.Code)
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
		req := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		mockUsecase.EXPECT().CreateTask(gomock.Any()).Return(nil).Times(1)
		e.ServeHTTP(rec, req)

		if rec.Code != http.StatusCreated {
			t.Errorf("Expected status code %d, got %d", http.StatusCreated, rec.Code)
		}
	})

	t.Run("異常系", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		mockUsecase.EXPECT().CreateTask(gomock.Any()).
			Return(errors.New("作成に失敗しました")).Times(1)
		e.ServeHTTP(rec, req)

		if rec.Code != http.StatusInternalServerError {
			t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, rec.Code)
		}
	})
}

func TestUpdateTask(t *testing.T) {
	reqBody, err := json.Marshal(
		dto.TaskRequest{
			ID:    testutil.PointerOf(uint(1)),
			Title: "Updateテスト",
			Desc:  "テスト",
		},
	)
	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	t.Run("正常系", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/tasks/1", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		mockUsecase.EXPECT().UpdateTask(gomock.Any()).Return(nil).Times(1)
		e.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, rec.Code)
		}
	})

	t.Run("400", func(t *testing.T) {
		badReqBody := []byte(`{
			"ID": 1,
			"Title": "Createテスト",
			"Desc": "テスト"
		`)
		req := httptest.NewRequest(http.MethodPut, "/tasks/1", bytes.NewBuffer(badReqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		if rec.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("500", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/tasks/1", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		mockUsecase.EXPECT().UpdateTask(gomock.Any()).
			Return(errors.New("更新失敗")).Times(1)
		e.ServeHTTP(rec, req)

		if rec.Code != http.StatusInternalServerError {
			t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, rec.Code)
		}
	})
}
