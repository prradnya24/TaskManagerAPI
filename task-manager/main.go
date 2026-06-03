package main

import (
	"task-manager/config"
	"task-manager/models"
	"task-manager/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(&models.Task{})
	r := gin.Default()
	routes.SetupRoutes(r)

	r.Run(":8080")

}
