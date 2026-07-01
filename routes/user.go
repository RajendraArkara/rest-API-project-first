package routes

import (
	"net/http"

	"github.com/RajendraArkara/first-api-project/models"
	"github.com/RajendraArkara/first-api-project/utils"
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

func GetUser(contect *gin.Context) {
	user, err := models.GetUser()
	if err != nil {
		contect.JSON(http.StatusInternalServerError, gin.H{
			"Message": "User not found",
		})
	}

	contect.JSON(http.StatusOK, user)
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

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"Message": "Could not authenticate user",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"Message": "Login successful",
		"Token":   token,
	})
}
