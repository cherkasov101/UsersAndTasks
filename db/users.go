package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// Function to check if the table exists and if not create users table
func CreateUsersTable(DB *sql.DB) error {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        balance INTEGER
	);
`
	_, err := DB.Exec(createTableSQL)
	return err
}
