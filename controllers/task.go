package controllers

import (
	"task-manager-backend/db"
	"task-manager-backend/models"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	userID := c.MustGet("user_id").(uint)
	task.CreatedBy = userID
	db.DB.Create(&task)
	Broadcast("Task created: " + task.Title) // Simple broadcast
	c.JSON(200, task)
}

func GetTasks(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	var tasks []models.Task
	db.DB.Where("created_by = ? OR assigned_to = ?", userID, userID).Find(&tasks)
	c.JSON(200, tasks)
}

func GetUsers(c *gin.Context) {
	var users []models.User
	db.DB.Find(&users)
	c.JSON(200, users)
}
