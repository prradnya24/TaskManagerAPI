package controllers

import (
	"net/http"

	"task-manager/config"
	"task-manager/models"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {

	var task models.Task

	c.BindJSON(&task)

	config.DB.Create(&task)

	c.JSON(http.StatusCreated, task)
}

func GetTasks(c *gin.Context) {

	var tasks []models.Task

	config.DB.Find(&tasks)

	c.JSON(http.StatusOK, tasks)
}