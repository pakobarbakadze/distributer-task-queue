package main

import (
	"distributed-task-queue/api"
	"distributed-task-queue/db"
	"distributed-task-queue/queue"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	db.AutoMigrate()

	queue.InitQueue()

	r := gin.Default()
	r.POST("/tasks", api.SubmitTask)

	r.Run(":8080")
}
