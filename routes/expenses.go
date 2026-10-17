package routes

import (
	"fmt"
	"net/http"

	"exampl.com/expense-tracker/models"
	"github.com/gin-gonic/gin"
)

func getExpenses(context *gin.Context) {
	expenses, err := models.GetAllExpenses()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch expenses. Try again leter."})
		return
	}

	context.JSON(http.StatusOK, expenses)
}

func createExpense(context *gin.Context) {
	var expense models.Expense
	err := context.ShouldBindJSON(&expense)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	// Handle user assigment...

	err = expense.Save()
	fmt.Println("ERRRROOOOOOOR: ", err)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create expense. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Expense created!"})
}
