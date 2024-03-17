package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// Function to check if the table exists and if not create quests table
func CreateQuestsTable(DB *sql.DB) error {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS quests (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        cost INTEGER
    );
`
	_, err := DB.Exec(createTableSQL)
	return err
}
