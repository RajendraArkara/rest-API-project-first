package main

import (
	"github.com/RajendraArkara/first-api-project/db"
	"github.com/RajendraArkara/first-api-project/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(":8080") // Localhost:8080
}
