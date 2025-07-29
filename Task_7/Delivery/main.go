package main

import (
	"task-manager/Delivery/routers"
	"task-manager/Infrastructure"
	"task-manager/Repositories"
	"task-manager/Usecases"
	"task-manager/Delivery/controllers"
)

func main() {
	// Mongo URI
	dbURI := "mongodb://localhost:27017"

	// Initialize repositories
	userRepo := repositories.NewUserRepositoryMongo(dbURI)
	taskRepo := repositories.NewTaskRepositoryMongo(dbURI)

	// Services
	passwordService := infrastructure.NewPasswordService()
	jwtService := infrastructure.NewJWTService("your_secret_key")

	// Usecases
	userUsecase := usecases.NewUserUsecase(userRepo, passwordService, jwtService)
	taskUsecase := usecases.NewTaskUsecase(taskRepo)

	// Controllers
	userController := controllers.NewUserController(userUsecase)
	taskController := controllers.NewTaskController(taskUsecase)

	// Router
	r := routers.SetupRouter(userController, taskController, jwtService)
	
	// Run server
	r.Run(":8080")
}
