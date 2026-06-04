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


func UpdateTask(c *gin.Context) {
	id := c.Param("id")

	var task models.Task

	// check if task exists
	result := config.DB.First(&task, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	// bind new data from request body
	var updatedTask models.Task
	if err := c.BindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// update only the fields sent
	config.DB.Model(&task).Updates(updatedTask)

	c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")

	var task models.Task

	// check if task exists
	result := config.DB.First(&task, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	config.DB.Delete(&task, id)

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}