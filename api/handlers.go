package api

import (
	"distributed-task-queue/db"
	"log"
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

	log.Println(task)

	c.JSON(http.StatusAccepted, task)
}
