package routers

import (
	"task-manager/Delivery/controllers"
	"task-manager/Infrastructure"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userController *controllers.UserController, taskController *controllers.TaskController, jwtService *infrastructure.JWTService) *gin.Engine {
	r := gin.Default()

	// Auth routes
	r.POST("/register", userController.Register)
	r.POST("/login", userController.Login)

	// Task routes (protected)
	taskRoutes := r.Group("/tasks")
	taskRoutes.Use(infrastructure.AuthMiddleware(jwtService))
	{
		taskRoutes.GET("/", taskController.GetAll)
		taskRoutes.GET("/:id", taskController.GetByID)
		taskRoutes.POST("/", taskController.Create)
		taskRoutes.PUT("/:id", taskController.Update)
		taskRoutes.DELETE("/:id", taskController.Delete)
	}

	return r
}
