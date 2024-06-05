package usecase_test

import (
	"errors"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sakaguchi-0725/go-todo/internal/domain"
	mocks "github.com/sakaguchi-0725/go-todo/internal/test/mocks/repository"
	"github.com/sakaguchi-0725/go-todo/internal/usecase"
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
