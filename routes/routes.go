package routes

import (
	"github.com/RajendraArkara/first-api-project/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/event", getEvents)
	server.GET("/event/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/event", createEvents)
	authenticated.PUT("/event/:id", updateEvent)
	authenticated.DELETE("/event/:id", deleteEvent)

	server.POST("/signup", signUp)
	server.POST("/login", login)
	server.GET("/user", GetUser)
}
