package models

type ToDoList struct {
	ID     string `json:"id,omitempty"`
	Task   string `json:"task,omitempty"`
	Status string `json:"status,omitempty"`
}
