package router

import (
	"task_manager/controllers"
	"task_manager/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine{
	r := gin.Default()

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	taskRoutes := r.Group("/tasks")
	taskRoutes.Use(middleware.AuthMiddleware())
	{
		taskRoutes.GET("/", controllers.GetTasks)
		taskRoutes.GET("/:id", controllers.GetTask)
		taskRoutes.POST("/", middleware.RoleMiddleware("admin"), controllers.CreateTask)
		taskRoutes.PUT("/:id", middleware.RoleMiddleware("admin"), controllers.UpdateTask)
		taskRoutes.DELETE("/:id", middleware.RoleMiddleware("admin"), controllers.DeleteTask)
	}
	
	return r
} 