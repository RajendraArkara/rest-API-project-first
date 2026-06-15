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

// semua request handler paraneter nya dari gin.Context
func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"Message": "could not parse the data",
			"error":   err.Error(),
		})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"Message": "Could not authenticed user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Message": "Login successful"})
}
