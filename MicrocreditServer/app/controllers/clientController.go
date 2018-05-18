package controllers

import (
	"github.com/revel/revel"
	"MicrocreditServer/app/dal/domain/models"
	"MicrocreditServer/app/dal"
	"github.com/dgrijalva/jwt-go"
	"errors"
	"MicrocreditServer/app"
	"MicrocreditServer/app/bll"
)

type ClientController struct {
	*revel.Controller
}

func ValidateToken(tokenString string) (bool, error){
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error){
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			revel.INFO.Println("UNEXPECTED SIGNIN METHOD: %v", token.Header["alg"])
			return nil, errors.New("unexpected signin method")
		}
		return app.GetSecretKey(), nil
	})

	if err != nil {
		return false, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		revel.INFO.Println(claims["email"], claims["exp"])
		return true, nil
	} else {
		return false, nil
	}

}

func (c ClientController) List() revel.Result{
	var clients []models.Client
	var uow dal.UnitOfWork

	uow.ClientRepository.List(&clients)

	return c.RenderJSON(clients)
}

func (c ClientController) Get(id int) revel.Result{
	var client models.Client
	var uow dal.UnitOfWork

	uow.ClientRepository.Get(id, &client)

	return c.RenderJSON(client)
}
func (c ClientController) Auth() revel.Result{
	var jsonData map[string] string
	c.Params.BindJSON(&jsonData)
	var userBl = new(bll.UserModel)
	var adminBl = new(bll.AdminModel)
	adminBl.Password = jsonData["password"]
	adminBl.Email = jsonData["email"]
	token, _:= adminBl.GetToken()
	if token != "" {
		return c.RenderJSON(token)
	}
	userBl.Password = jsonData["password"]
	userBl.Email = jsonData["email"]
	tokenUser, err := userBl.GetToken()
	if err != nil{
		return c.RenderJSON(err)
	}
	return c.RenderJSON(tokenUser)
}

func (c ClientController) Create() revel.Result{
	var user bll.UserModel
	c.Params.BindJSON(&user)
	user.Registration()

	return c.RenderJSON("You have been registered")
}

func (c ClientController) Remove(id int) revel.Result{
	var uow dal.UnitOfWork
	uow.ClientRepository.Delete(id)
	return c.RenderJSON(id)
}

func (c ClientController) Update() revel.Result{
	var clientPut models.Client
	var uow dal.UnitOfWork

	c.Params.BindJSON(&clientPut)
	uow.ClientRepository.Update(&clientPut)

	return c.RenderJSON("Your information was updated!")
}
