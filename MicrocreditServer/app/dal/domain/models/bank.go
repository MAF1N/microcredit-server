package models

type Bank struct {
	Id int `json:"id" db:"id" gorm:"primary_key"`
	Title string `json:"title" db:"title"`
	Address string `json:"address" db:"address"`
	ContactPhone string `json:"contact_phone" db:"contact_phone"`
}