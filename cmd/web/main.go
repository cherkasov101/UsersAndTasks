package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"

	"github.com/cherkasov101/UsersAndTasks/db"
	"github.com/cherkasov101/UsersAndTasks/internal/handlers"
)

var DB *sql.DB

func main() {
	var err error
	DB, err = db.ConnectDB()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer DB.Close()

	err = db.LaunchDB(DB)
	if err != nil {
		log.Fatal(err)
		return
	}

	r := mux.NewRouter()
	r.HandleFunc("/create-user", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateUser(w, r, DB)
	}).Methods(http.MethodPost)
	r.HandleFunc("/create-quest", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateQuest(w, r, DB)
	}).Methods(http.MethodPost)
	r.HandleFunc("/quest-done", func(w http.ResponseWriter, r *http.Request) {
		handlers.QuestDone(w, r, DB)
	}).Methods(http.MethodPost)
	r.HandleFunc("/get-history", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetHistory(w, r, DB)
	}).Methods(http.MethodGet)
	r.HandleFunc("/get-users", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetUsers(w, r, DB)
	}).Methods(http.MethodGet)
	r.HandleFunc("/get-quests", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetQuests(w, r, DB)
	}).Methods(http.MethodGet)

	http.ListenAndServe(":8080", r)
}
