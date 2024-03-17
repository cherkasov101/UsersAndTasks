package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// Function to check if the table exists and if not create completed_quests table
func CreateCompletedQuestsTable(DB *sql.DB) error {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS completed_quests (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		quest_id INTEGER,
		user_id INTEGER
	);
`
	_, err := DB.Exec(createTableSQL)
	return err
}
