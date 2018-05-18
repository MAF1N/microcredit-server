package models

type Organization struct {
	Id int `json:"id" db:"id" gorm:"primary_key"`
	Title string `json:"title" db:"title"`
	Type string `json:"type" db:"type"`
	Amount float64 `json:"amount" db:"amount"`
	BankId int `json:"bank_id" db:"bank_id"`
	CreditAccounts []CreditAccount
}