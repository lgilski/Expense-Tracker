package routes

import "github.com/gin-gonic/gin"

// Create routes
func RegisterRoutes(server *gin.Engine) {
	server.GET("/expenses", getExpenses)
	server.POST("/expenses", createExpense)
}
