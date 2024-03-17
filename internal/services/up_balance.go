package services

import (
	"database/sql"
)

func UpBalance(user_id int, quest_id int, DB *sql.DB) error {
	var cost int
	err := DB.QueryRow("SELECT cost FROM quests WHERE id = ?", quest_id).Scan(&cost)
	if err != nil {
		return err
	}

	var balance int
	err = DB.QueryRow("SELECT balance FROM users WHERE id = ?", user_id).Scan(&balance)
	if err != nil {
		return err
	}

	balance += cost

	_, err = DB.Exec("UPDATE users SET balance = balance + ? WHERE id = ?", balance, user_id)
	return err
}