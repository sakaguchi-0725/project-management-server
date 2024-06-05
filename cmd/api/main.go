package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/sakaguchi-0725/go-todo/internal/infrastructure/database"
	"github.com/sakaguchi-0725/go-todo/internal/infrastructure/repository"
	"github.com/sakaguchi-0725/go-todo/internal/interface/handler"
	"github.com/sakaguchi-0725/go-todo/internal/interface/middleware"
	"github.com/sakaguchi-0725/go-todo/internal/usecase"
	"github.com/sakaguchi-0725/go-todo/pkg"
	"go.uber.org/zap"
)

func main() {
	pkg.LoadEnv()
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Loggerの初期化に失敗しました: %v", err)
	}

	e := echo.New()
	e.Use(middleware.LoggerMiddleware(logger))
	e.Use(middleware.ErrorMiddleware(logger))
	db := database.NewDB()

	taskRepository := repository.NewTaskRepository(db)
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	taskHandler := handler.NewTaskHandler(taskUsecase)

	t := e.Group("/tasks")
	t.GET("", taskHandler.GetAllTasks)
	e.Logger.Fatal(e.Start(":8080"))
}
