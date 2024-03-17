package handlers

import (
	"encoding/json"
	"net/http"
	"database/sql"
	"log"

	"github.com/cherkasov101/UsersAndTasks/internal/models"
)

// Function to get history of completed quests for the user
func GetHistory(w http.ResponseWriter, r *http.Request, DB *sql.DB) {
	userID := r.URL.Query().Get("user_id")

	if userID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	var userName string
	err := DB.QueryRow("SELECT name FROM users WHERE id = ?", userID).Scan(&userName)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	completedQuests, err := DB.Query("SELECT * FROM completed_quests WHERE user_id = ?", userID)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer completedQuests.Close()

	//var completedQuestsList []models.CompletedQuest
	var questNames []string

	for completedQuests.Next() {
		var completedQuest models.CompletedQuest
		err := completedQuests.Scan(&completedQuest.ID, &completedQuest.QuestId, &completedQuest.UserId)
		if err != nil {
			log.Fatal(err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		questName := ""
		err = DB.QueryRow("SELECT name FROM quests WHERE id = ?", completedQuest.QuestId).Scan(&questName)
		if err != nil {
			log.Fatal(err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		questNames = append(questNames, questName)
		//completedQuestsList = append(completedQuestsList, completedQuest)
	}

	response := struct {
		UserName          string   `json:"user_name"`
		CompletedQuests   []string `json:"completed_quests"`
	}{
		UserName:        userName,
		CompletedQuests: questNames,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
