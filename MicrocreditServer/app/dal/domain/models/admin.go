package models

type Admin struct {
	Id int `json:"id" db:"id" gorm:"primary_key"`
	Name string `json:"name" db:"name"`
	Surname string `json:"surname" db:"surname"`
	Email string `json:"email" db:"email"`
	Phone string `json:"phone" db:"phone"`
	Password string `json:"password" db:"password"`
	OrganizationId int `json:"organization_id" db:organization_id`
}
