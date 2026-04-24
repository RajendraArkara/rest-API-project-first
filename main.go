package main

import (
	"net/http"

	"github.com/RajendraArkara/first-api-project/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/event", getEvents)
	server.POST("/event", createEvents)

	server.Run(":8080") // Localhost:8080
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvents(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "coul not parse the data"})
		return
	}

	event.ID = 1
	event.UserID = 1
	context.JSON(http.StatusCreated, gin.H{"message": "event created!", "event": event})
}
