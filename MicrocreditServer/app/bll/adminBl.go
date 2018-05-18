package bll

import (
	"crypto/sha256"
	"encoding/base64"
	"MicrocreditServer/app/dal/domain/models"
	"errors"
	"time"
	"MicrocreditServer/app/dal"
	"github.com/dgrijalva/jwt-go"
	"MicrocreditServer/app"
)

type AdminModel struct {
	Id int `json:"id" db:"id" gorm:"primary_key"`
	Name string `json:"name" db:"name"`
	Surname string `json:"surname" db:"surname"`
	Email string `json:"email" db:"email"`
	Phone string `json:"phone" db:"phone"`
	Password string `json:"password" db:"password"`
	OrganizationId int `json:"organization_id" db:organization_id`
}

func NewAdminDal (model AdminModel) models.Admin{
	admin := models.Admin{
		Name: model.Name,
		Surname: model.Surname,
		Email: model.Email,
		Password: model.Password,
		Phone: model.Phone,
		OrganizationId:model.OrganizationId}
	return admin
}
func NewAdminModel(name string, surname string, email string, password string, phone string) AdminModel{
	adminBl := AdminModel{Name: name, Surname: surname, Email: email, Password: password, Phone: phone}
	adminBl.cryptPassword()
	return adminBl
}

func (adminBl *AdminModel) cryptPassword(){
	hasher := sha256.New()
	hasher.Write([]byte(adminBl.Password))
	adminBl.Password = base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}
func (adminBl AdminModel) getAdminFromDB() models.Admin{
	var uow dal.UnitOfWork
	var admin models.Admin
	uow.AdminRepository.GetByEmail(adminBl.Email, &admin)
	return admin
}
func (admin *AdminModel) checkPassword(password string) bool{
	admin.cryptPassword()
	if password == admin.Password {
		return true
	}
	return false
}

func (adminBl *AdminModel) GetToken() (string, error) {
	admin := adminBl.getAdminFromDB()

	passwordCorrect := adminBl.checkPassword(admin.Password)
	if !passwordCorrect {
		return "", errors.New("password incorrect")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims {
		"admin": true,
		"email": adminBl.Email,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	secretKey := app.GetSecretKey()
	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}