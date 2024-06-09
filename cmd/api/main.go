package main

import (
	"github.com/sakaguchi-0725/go-todo/cmd/server"
	"github.com/sakaguchi-0725/go-todo/internal/infrastructure/database"
	"github.com/sakaguchi-0725/go-todo/internal/infrastructure/repository"
	"github.com/sakaguchi-0725/go-todo/internal/interface/handler"
	"github.com/sakaguchi-0725/go-todo/internal/usecase"
	"github.com/sakaguchi-0725/go-todo/pkg/config"
)

func main() {
	config.LoadConfig()
	db := database.NewDB()

	taskRepository := repository.NewTaskRepository(db)
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	taskHandler := handler.NewTaskHandler(taskUsecase)

	handlers := server.AppHandlers{
		TaskHandler: taskHandler,
	}

	e := server.NewServer(handlers)
	e.Logger.Fatal(e.Start(":8080"))
}
