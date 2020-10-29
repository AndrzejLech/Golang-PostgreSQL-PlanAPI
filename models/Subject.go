package models

type Subject struct {
	ID 		string `json:"id"`
	Name 	string `json:"name"`
	Teacher string `json:"teacher"`
	Hour 	string `json:"hour"`
}