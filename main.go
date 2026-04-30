package main

import (
	"net/http"

	"github.com/RajendraArkara/first-api-project/db"
	"github.com/RajendraArkara/first-api-project/models"
	"github.com/gin-gonic/gin"
)

// cuman buat ijoin github

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/event", getEvents)
	server.POST("/event", createEvents)

	server.Run(":8080") // Localhost:8080
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"Message": "Could not fetch events data. Try again layter",
			"Error":   err.Error(),
		})
	}
	context.JSON(http.StatusOK, events)
}

func createEvents(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"Message": "could not parse the data",
			"error":   err.Error()})
		return
	}

	event.ID = 1
	event.UserID = 1

	event.Save()

	context.JSON(http.StatusCreated, gin.H{"message": "event created!", "event": event})
}
