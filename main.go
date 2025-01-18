package main

import (
	"distributed-task-queue/api"
	"distributed-task-queue/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	db.AutoMigrate()

	r := gin.Default()
	r.POST("/tasks", api.SubmitTask)

	r.Run(":8080")
}
