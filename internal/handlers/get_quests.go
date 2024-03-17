package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/cherkasov101/UsersAndTasks/internal/models"
)

// Handler to get all quests
func GetQuests(w http.ResponseWriter, r *http.Request, DB *sql.DB) {
	quests, err := DB.Query("SELECT * FROM quests")
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer quests.Close()

	var questList []models.Quest

	for quests.Next() {
		var quest models.Quest
		err := quests.Scan(&quest.ID, &quest.Name, &quest.Cost)
		if err != nil {
			log.Fatal(err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		questList = append(questList, quest)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(questList); err != nil {
		log.Fatal(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
