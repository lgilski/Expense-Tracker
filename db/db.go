package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createExpensesTable := `
	CREATE TABLE IF NOT EXISTS expenses (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		description TEXT NOT NULL,
		category TEXT NOT NULL,
		amount REAL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER
	)
	`
	// FOREIGN KEY(user_id) REFERENCES users(id)

	_, err := DB.Exec(createExpensesTable)

	if err != nil {
		panic("Could not create expenses table.")
	}
}
