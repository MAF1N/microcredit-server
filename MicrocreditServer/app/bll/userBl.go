package bll

import (
	"crypto/sha256"
	"encoding/base64"
	"unicode"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
	"github.com/revel/revel"
	"MicrocreditServer/app"
	"MicrocreditServer/app/dal"
	"MicrocreditServer/app/dal/domain/models"
)


type UserModel struct {
	Id int `json:"id" db:"id" gorm:"primary_key"`
	Name string `json:"name" db:"name"`
	Surname string `json:"surname" db:"surname"`
	Email string `json:"email" db:"email"`
	Phone string `json:"phone" db:"phone"`
	Password string `json:"password" db:"password"`
}

func NewClient (user UserModel) models.Client{
	client := models.Client{Name: user.Name, Surname: user.Surname, Email: user.Email, Password: user.Password, Phone: user.Phone}
	return client
}

func AuthUserModel(email string, password string) UserModel{
	userBl := UserModel{Email: email, Password: password}
	userBl.cryptPassword()
	return userBl
}

func (userBl *UserModel) cryptPassword(){
	hasher := sha256.New()
	hasher.Write([]byte(userBl.Password))
	userBl.Password = base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}

func (userBl  *UserModel) checkPassword(password string) bool{
	userBl.cryptPassword()
	if password == userBl.Password {
		return true
	}
	return false
}

func (userBl *UserModel) verifyPassword() (sevenOrMore, number, lower, upper, special bool) {
	lettersCount := 0
	for _, char := range userBl.Password {
		switch {
		case unicode.IsNumber(char):
			number = true
		case unicode.IsUpper(char):
			upper = true
		case unicode.IsLower(char):
			lower = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			special = true
		default:
			//return false, false, false, false
		}
		lettersCount++;
	}
	sevenOrMore = lettersCount >= 7
	return
}

func (userBl UserModel) getUserFromDB() models.Client{
	var uow dal.UnitOfWork
	var user models.Client
	uow.ClientRepository.GetByEmail(userBl.Email, &user)
	return user
}

func (userBl UserModel) GetToken() (string, error) {
	user := userBl.getUserFromDB()

	passwordCorrect := userBl.checkPassword(user.Password)
	if !passwordCorrect {
		return "", errors.New("password incorrect")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims {
		"admin": false,
		"email": userBl.Email,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	secretKey := app.GetSecretKey()
	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (userBl UserModel) Registration() error {
	sevenOrMore, number, lower, upper, _ := userBl.verifyPassword()
	var unitOfWork dal.UnitOfWork
	if !(sevenOrMore && number && lower && upper) {
		revel.INFO.Println(sevenOrMore)
		revel.INFO.Println(number)
		revel.INFO.Println(lower)
		revel.INFO.Println(upper)
		return errors.New("password must contain number, lower, upper and length more or 7")
	}

	if userBl.Email == "" {
		return errors.New("email must'n empty")
	}
	userBl.cryptPassword()
	user := NewClient(userBl)
	err := unitOfWork.ClientRepository.Create(&user)

	if err != nil {
		return err
	}

	return nil
}