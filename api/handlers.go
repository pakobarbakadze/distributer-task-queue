package api

import (
	"distributed-task-queue/db"
	"distributed-task-queue/queue"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SubmitTask(c *gin.Context) {
	var task db.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.ID = uuid.New().String()
	task.Status = "pending"
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	if err := db.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save task to database"})
	}

	if err := queue.PublishTask(task.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to publish task to queue"})
		return
	}

	c.JSON(http.StatusAccepted, task)
}
