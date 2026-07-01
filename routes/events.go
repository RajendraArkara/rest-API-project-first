package routes

import (
	"net/http"
	"strconv"

	"github.com/RajendraArkara/first-api-project/models"
	"github.com/RajendraArkara/first-api-project/utils"
	"github.com/gin-gonic/gin"
)

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

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"Message": "Could not parse event ID",
			"Error":   err.Error(),
		})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"Message": "Could not parse event ID.",
			"Error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, event)
}

func createEvents(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{
			"Message": "Not Authorize",
		})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"Message": "Not Authorize",
		})
		return
	}

	var event models.Event
	err = context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"Message": "could not parse the data",
			"error":   err.Error(),
		})
		return
	}

	event.UserID = userId

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"Message": "Could not fetch events. Try again later",
			"Error":   err.Error(),
		})
	}

	context.JSON(http.StatusCreated, gin.H{"message": "event created!", "event": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"Message": "Could not parse event ID",
			"Error":   err.Error(),
		})
		return
	}

	_, err = models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"Message": "Could not fetch events. Try again later (id nya gak ada!)",
			"Error":   err.Error(),
		})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"Message": "Invalid request data",
			"error":   err.Error(),
		})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"Message": "Could not update event. Try again later.",
			"error":   err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "event update!"})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"Message": "Could not parse event ID",
			"Error":   err.Error(),
		})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"Message": "Could not fetch the id.",
			"Error":   err.Error(),
		})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"Message": "Could not delete data.",
			"Error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "event deleted!"})
}
