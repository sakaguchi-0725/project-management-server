package repository_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/sakaguchi-0725/go-todo/internal/domain"
	"github.com/sakaguchi-0725/go-todo/internal/infrastructure/repository"
	testutil "github.com/sakaguchi-0725/go-todo/internal/test/util"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	err  error
	repo repository.TaskRepository
)

func TestMain(m *testing.M) {
	db, err = testutil.NewTestDB()
	if err != nil {
		fmt.Printf("database connection error: %d", err)
		return
	}
	repo = repository.NewTaskRepository(db)

	status := m.Run()

	os.Exit(status)
}

func TestGetAllTasks(t *testing.T) {
	tasks := []*domain.Task{
		{ID: 1, Title: "テスト01"},
		{ID: 2, Title: "テスト02"},
	}
	createTestData(tasks)

	t.Run("正常系", func(t *testing.T) {
		defer testutil.FlushRecords(db, &domain.Task{})

		result, err := repo.GetAllTasks()
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(result) != 2 {
			t.Errorf("Expected result length to be 2, got %d", len(result))
		}

		for i, task := range result {
			if task.Title != tasks[i].Title {
				t.Errorf("Expected task title to be '%s', got '%s'", tasks[i].Title, task.Title)
			}
		}
	})
}

func createTestData(data []*domain.Task) {
	for _, task := range data {
		if err := db.Create(task).Error; err != nil {
			log.Fatalf("Failed to insert task: %v", err)
		}
	}
}
