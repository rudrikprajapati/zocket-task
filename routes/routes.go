package routes

import (
	"task-manager-backend/controllers"
	"task-manager-backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	auth := r.Group("/").Use(middleware.Auth())
	auth.POST("/tasks", controllers.CreateTask)
	auth.GET("/tasks", controllers.GetTasks)
	auth.GET("/users", controllers.GetUsers)
	auth.POST("/tasks/suggest", controllers.SuggestTasks)
	auth.GET("/ws", controllers.WebSocketHandler)
}
