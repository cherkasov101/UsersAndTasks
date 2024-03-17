package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// Function to create the tables
func LaunchDB(DB *sql.DB) error {
	err := CreateUsersTable(DB)
	if err != nil {
		return err
	}

	err = CreateQuestsTable(DB)
	if err != nil {
		return err
	}

	err = CreateCompletedQuestsTable(DB)
	if err != nil {
		return err
	}

	return nil
}
