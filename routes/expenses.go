package routes

import (
	"fmt"
	"net/http"
	"strconv"

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

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create expense. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Expense created!"})
}

func deleteExpense(context *gin.Context) {
	expenseId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse expense id."})
		return
	}

	expense, err := models.GetExpenseByID(expenseId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the expense."})
		return
	}

	err = expense.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the expense."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Deleted expense successfully!"})
}

func updateExpense(context *gin.Context) {
	expenseId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse expense id."})
		return
	}

	// This will be for user assigment

	// expense, err := models.GetExpenseByID(expenseId)

	// if err != nil {
	// 	context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the expense."})
	// 	return
	// }

	var updatedExpense models.Expense
	err = context.ShouldBindJSON(&updatedExpense)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse request data."})
		return
	}

	updatedExpense.ID = expenseId

	err = updatedExpense.Update()

	fmt.Println(err)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update the expense."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Expense updated successfully!"})
}
