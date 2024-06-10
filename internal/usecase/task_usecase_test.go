package usecase_test

import (
	"errors"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sakaguchi-0725/go-todo/internal/domain"
	mocks "github.com/sakaguchi-0725/go-todo/internal/test/mocks/repository"
	"github.com/sakaguchi-0725/go-todo/internal/usecase"
	"github.com/sakaguchi-0725/go-todo/internal/usecase/input"
)

var (
	mockRepo    *mocks.MockTaskRepository
	taskUsecase usecase.TaskUsecase
)

func TestMain(m *testing.M) {
	ctrl := gomock.NewController(&testing.T{})
	mockRepo = mocks.NewMockTaskRepository(ctrl)
	taskUsecase = usecase.NewTaskUsecase(mockRepo)

	status := m.Run()
	os.Exit(status)
}

func TestGetAllTasks(t *testing.T) {
	res := []domain.Task{
		{ID: 1, Title: "テスト01"},
		{ID: 2, Title: "テスト02"},
	}

	t.Run("正常系", func(t *testing.T) {
		mockRepo.EXPECT().GetAllTasks().Return(res, nil).Times(1)
		result, err := taskUsecase.GetAllTasks()

		if err != nil {
			t.Errorf("GetAllTasks failed: expected no error, got %v", err)
		}

		if len(result) != 2 {
			t.Errorf("GetAllTasks returned %d tasks, expected %d", len(result), len(res))
		}

		for i, task := range result {
			if task.Title != res[i].Title {
				t.Errorf("Task %d title mismatch: expected %s, got %s", i, res[i].Title, task.Title)
			}
		}
	})

	t.Run("異常系", func(t *testing.T) {
		mockRepo.EXPECT().GetAllTasks().
			Return(nil, errors.New("データの取得に失敗しました")).
			Times(1)

		_, err := taskUsecase.GetAllTasks()
		if err == nil {
			t.Errorf("Expected an error from GetAllTasks, but received none")
		}

		if err != nil && err.Error() != "データの取得に失敗しました" {
			t.Errorf("Unexpected error message: got '%s', want 'データの取得に失敗しました'", err.Error())
		}
	})
}

func TestCreateTask(t *testing.T) {
	req := input.TaskInput{
		Title: "Createテスト",
		Desc:  "テスト",
	}

	t.Run("正常系", func(t *testing.T) {
		mockRepo.EXPECT().CreateTask(gomock.Any()).
			Return(nil).Times(1)

		if err := taskUsecase.CreateTask(req); err != nil {
			t.Errorf("Expect no errors, got %d", err)
		}
	})

	t.Run("異常系", func(t *testing.T) {
		mockRepo.EXPECT().CreateTask(gomock.Any()).
			Return(errors.New("データの作成に失敗しました")).Times(1)

		err := taskUsecase.CreateTask(req)
		if err == nil {
			t.Error("Expected an error from CreateTask, but received none")
		}

		if err != nil && err.Error() != "データの作成に失敗しました" {
			t.Errorf("Unexpected error message: got '%s', want 'データの作成に失敗しました'", err.Error())
		}
	})
}

func TestUpdateTask(t *testing.T) {
	req := input.TaskInput{
		ID:    1,
		Title: "テスト",
		Desc:  "",
	}
	t.Run("正常系", func(t *testing.T) {
		mockRepo.EXPECT().UpdateTask(gomock.Any()).Return(nil).Times(1)

		err := taskUsecase.UpdateTask(req)
		if err != nil {
			t.Errorf("Expect no error, got %v", err)
		}
	})

	t.Run("異常系", func(t *testing.T) {
		mockRepo.EXPECT().UpdateTask(gomock.Any()).
			Return(errors.New("更新失敗")).Times(1)

		err := taskUsecase.UpdateTask(req)
		if err == nil {
			t.Errorf("Expected an error from UpdateTask, but received none")
		}

		if err.Error() != "更新失敗" {
			t.Errorf("Unexpected error message: got '%s', want '更新失敗'", err.Error())
		}
	})
}

func TestDeleteTask(t *testing.T) {
	t.Run("正常系", func(t *testing.T) {
		mockRepo.EXPECT().DeleteTask(gomock.Any()).Return(nil).Times(1)

		err := taskUsecase.DeleteTask(1)
		if err != nil {
			t.Errorf("Expect no error, got %v", err)
		}
	})

	t.Run("異常系", func(t *testing.T) {
		mockRepo.EXPECT().DeleteTask(gomock.Any()).
			Return(errors.New("削除失敗")).Times(1)

		err := taskUsecase.DeleteTask(1)
		if err == nil {
			t.Errorf("Expected an error from DeleteTask, but received none")
		}

		if err.Error() != "削除失敗" {
			t.Errorf("Unexpected error message: got '%s', want '削除失敗'", err.Error())
		}
	})
}
