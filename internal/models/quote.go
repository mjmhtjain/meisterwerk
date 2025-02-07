package models

type Quote struct {
	ID     string `json:"id"`
	Author string `json:"author"`
	Status string `json:"status"`
}
