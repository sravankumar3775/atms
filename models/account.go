package models

type Account struct {
	AccountID	int `json:"accountID,omitempty"`
	UserName	string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Status string `json:"status,omitempty"`
}