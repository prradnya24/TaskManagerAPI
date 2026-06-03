package routes

import (
	"task-manager/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.POST("/tasks", controllers.CreateTask)

	r.GET("/tasks", controllers.GetTasks)
}
