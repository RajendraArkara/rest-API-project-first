package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/event", getEvents)
	server.GET("/event/:id", getEvent)
	server.POST("/event", createEvents)
	server.PUT("/event/:id")
}
