package server

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/sakaguchi-0725/go-todo/internal/interface/handler"
	"github.com/sakaguchi-0725/go-todo/internal/interface/middleware"
	"go.uber.org/zap"
)

type AppHandlers struct {
	TaskHandler handler.TaskHandler
}

func NewServer(handlers AppHandlers) *echo.Echo {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Loggerの初期化に失敗しました: %v", err)
	}
	e := echo.New()
	e.Use(middleware.LoggerMiddleware(logger))
	e.Use(middleware.ErrorMiddleware(logger))

	t := e.Group("/tasks")
	t.GET("", handlers.TaskHandler.GetAllTasks)
	t.POST("", handlers.TaskHandler.CreateTask)

	return e
}
