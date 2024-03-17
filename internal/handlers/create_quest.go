package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/cherkasov101/UsersAndTasks/internal/models"
)

// Handler to create a quest
func CreateQuest(w http.ResponseWriter, r *http.Request, DB *sql.DB) {
	var quest models.Quest
	err := json.NewDecoder(r.Body).Decode(&quest)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	_, err = DB.Exec("INSERT INTO quests(name, cost) VALUES(?, ?)", quest.Name, quest.Cost)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"status": "ok"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
}
