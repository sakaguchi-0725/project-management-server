package handler_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/sakaguchi-0725/go-todo/internal/domain"
	"github.com/sakaguchi-0725/go-todo/internal/interface/handler"
	mocks "github.com/sakaguchi-0725/go-todo/internal/test/mocks/usecase"
)

func TestGetAllTasks(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUsecase := mocks.NewMockTaskUsecase(ctrl)
	taskHandler := handler.NewTaskHandler(mockUsecase)

	res := []domain.Task{
		{ID: 1, Title: "テスト001"},
		{ID: 2, Title: "テスト002"},
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
