package models

type Email struct {
	From	string `json:"from"`
	To 		string `json:"to"`
	Subject string `json:"subject"`
	Body 	string `json:"body"`
	Date	string `json:"date"`
}