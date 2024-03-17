package models

type CompletedQuest struct {
	ID      int `json:"id"`
	QuestId int `json:"quest_id"`
	UserId  int `json:"user_id"`
}
