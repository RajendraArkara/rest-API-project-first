package routes

import (
	"net/http"

	"github.com/RajendraArkara/first-api-project/models"
	"github.com/gin-gonic/gin"
)

func signUp(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"Message": "could not parse the data",
			"error":   err.Error(),
		})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"Message": "could not save user",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
	})
}
