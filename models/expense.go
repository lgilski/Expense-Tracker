package models

import (
	"time"

	"exampl.com/expense-tracker/db"
)

type Expense struct {
	ID          int64
	Description string    `binding:"required"`
	Category    string    `binding:"required"`
	Amount      float64   `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

func GetAllExpenses() ([]Expense, error) {
	// Get all expenses from db
	query := `SELECT * FROM expenses`
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var expenses []Expense

	// Go through all expenses rows - go through each expense in db
	for rows.Next() {
		var expense Expense

		// Populate single expense
		err := rows.Scan(&expense.ID, &expense.Description, &expense.Category, &expense.Amount, &expense.DateTime, &expense.UserID)

		if err != nil {
			return nil, err
		}

		// Add populated expense from db to expenses array
		expenses = append(expenses, expense)
	}

	return expenses, nil
}

func GetExpenseByID(expenseId int64) (*Expense, error) {
	query := `SELECT * FROM expenses WHERE id = ?`
	row := db.DB.QueryRow(query, expenseId)

	var expense Expense
	err := row.Scan(&expense.ID, &expense.Description, &expense.Category, &expense.Amount, &expense.DateTime, &expense.UserID)

	if err != nil {
		return nil, err
	}

	return &expense, nil
}

// Save adds new expense to the database.
// It also automatically adds new id.
func (e *Expense) Save() error {
	query := `
		INSERT INTO expenses(description, category, amount, dateTime, user_id)
		VALUES (?, ?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	// It runs this function every time when there is "return"
	defer stmt.Close()

	result, err := stmt.Exec(e.Description, e.Category, e.Amount, e.DateTime, e.UserID)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.ID = id
	return err
}

// Delete deletes an expense from database.
func (e Expense) Delete() error {
	query := `DELETE FROM expenses WHERE id = ?`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID)
	return err
}

func (e Expense) Update() error {
	query := `
	UPDATE expenses
	SET description = ?, category = ?, amount = ?, dateTime = ?
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Description, e.Category, e.Amount, e.DateTime, e.ID)
	return err
}
