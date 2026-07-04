package routes

import (
	"net/http"
	"strconv"

	"github.com/RajendraArkara/first-api-project/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
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
			"Message": "Could not fetch event.",
		})
		return
	}
}

func cancelRegistration(context *gin.Context) {

}
