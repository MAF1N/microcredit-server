package models

type CreditAccount struct {
	OrganizationId int `json:"organization_id" db:"organization_id" gorm:"primary_key"`
	ClientId int `json:"client_id" db:"client_id"`
	CardId string `json:"card_id" db:"card_id"`
	Amount float64 `json:"amount" db:"amount"`
	CreditPercent float32 `json:"credit_percent" db:"credit_percent"`
}