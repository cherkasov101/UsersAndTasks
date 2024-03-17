package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"database/sql"

	"github.com/cherkasov101/UsersAndTasks/internal/models"
	"github.com/cherkasov101/UsersAndTasks/internal/services"
)

// Handler to signal that the user has completed the quest
func QuestDone(w http.ResponseWriter, r *http.Request, DB *sql.DB) {
	var CompletedQuest models.CompletedQuest
	err := json.NewDecoder(r.Body).Decode(&CompletedQuest)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	check, err := checkValidQuest(CompletedQuest.QuestId, CompletedQuest.UserId, DB)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	check2, err := checkQuestNotDone(CompletedQuest.QuestId, CompletedQuest.UserId, DB)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if !check || !check2 {
		log.Fatal(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	fmt.Println(CompletedQuest.QuestId, CompletedQuest.UserId)
	_, err = DB.Exec("INSERT INTO completed_quests(quest_id, user_id) VALUES(?, ?)", 
		CompletedQuest.QuestId, CompletedQuest.UserId)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	err = services.UpBalance(CompletedQuest.UserId, CompletedQuest.QuestId, DB)
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

// Function to check if the user hasn't completed the quest before
func checkQuestNotDone(quest_id int, user_id int, DB *sql.DB) (bool, error) {
	completedQuests, err := DB.Query("SELECT * FROM completed_quests")
	if err != nil {
		return false, err
	}
	defer completedQuests.Close()

	var completedQuestsList []models.CompletedQuest

	for completedQuests.Next() {
		var completedQuest models.CompletedQuest
		err := completedQuests.Scan(&completedQuest.ID, &completedQuest.QuestId, &completedQuest.UserId)
		if err != nil {
			return false, err
		}
		completedQuestsList = append(completedQuestsList, completedQuest)
	}
	fmt.Println("Quest already done", completedQuestsList, quest_id, user_id)
	for _, completedQuest := range completedQuestsList {
		if completedQuest.QuestId == quest_id && completedQuest.UserId == user_id {
			fmt.Println("Quest already done", completedQuestsList, quest_id, user_id)
			return false, nil
		}
	}
	return true, nil
}

// Function to check if the user and the quest exist
func checkValidQuest(quest_id int, user_id int, DB *sql.DB) (bool, error) {
	users, err := DB.Query("SELECT * FROM users")
	if err != nil {
		return false, err
	}
	defer users.Close()

	var userList []models.User

	for users.Next() {
		var user models.User
		err := users.Scan(&user.ID, &user.Name, &user.Balance)
		if err != nil {
			return false, err
		}
		userList = append(userList, user)
	}

	quests, err := DB.Query("SELECT * FROM quests")
	if err != nil {
		return false, err
	}
	defer quests.Close()

	var questList []models.Quest

	for quests.Next() {
		var quest models.Quest
		err := quests.Scan(&quest.ID, &quest.Name, &quest.Cost)
		if err != nil {
			return false, err
		}
		questList = append(questList, quest)
	}

	for _, user := range userList {
		if user.ID == user_id {
			for _, quest := range questList {
				if quest.ID == quest_id {
					return true, nil
				}
			}
		}
	}
	return false, nil
}
