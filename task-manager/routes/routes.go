package routes

import (
	"task-manager/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.POST("/tasks", controllers.CreateTask)

	r.GET("/tasks", controllers.GetTasks)

	r.PUT("/tasks/:id", controllers.UpdateTask)

	r.DELETE("/tasks/:id", controllers.DeleteTask)

	r.GET("/tasks/:id", controllers.GetTaskByID)
}
