package routes

import (
	"exampl.com/expense-tracker/middlewares"
	"github.com/gin-gonic/gin"
)

// Create routes
func RegisterRoutes(server *gin.Engine) {

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.GET("/expenses", getExpenses)
	authenticated.POST("/expenses", createExpense)
	authenticated.DELETE("/expenses/:id", deleteExpense)
	authenticated.PUT("expenses/:id", updateExpense)

	server.POST("/login", login)
	server.POST("/signup", signup)
}
