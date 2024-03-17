package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/cherkasov101/UsersAndTasks/internal/models"
)

// Handler to get all users
func GetUsers(w http.ResponseWriter, r *http.Request, DB *sql.DB) {
	users, err := DB.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer users.Close()

	var userList []models.User

	for users.Next() {
		var user models.User
		err := users.Scan(&user.ID, &user.Name, &user.Balance)
		if err != nil {
			log.Fatal(err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		userList = append(userList, user)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(userList); err != nil {
		log.Fatal(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
