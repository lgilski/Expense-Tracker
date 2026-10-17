package main

import (
	"exampl.com/expense-tracker/db"
	"exampl.com/expense-tracker/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
