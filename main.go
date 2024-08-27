package main

import (
	"go-rest-api/controller"
	"go-rest-api/db"
	"go-rest-api/repository"
	"go-rest-api/router"
	"go-rest-api/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	taskRepository := repository.NewTaskRepository(db)
	taskUseCase := usecase.NewTaskUsecase(taskRepository)
	taskController := controller.NewTaskController(taskUseCase)
	e := router.NewRouter(userController, taskController)
	e.Logger.Fatal(e.Start(":8080"))

}
