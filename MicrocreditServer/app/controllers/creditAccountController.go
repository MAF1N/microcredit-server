package controllers

import (
	"github.com/revel/revel"
	"MicrocreditServer/app/dal/domain/models"
	"MicrocreditServer/app/dal"
	"strconv"
)

type CreditAccountController struct {
	*revel.Controller
}

func (c CreditAccountController) List() revel.Result{
	var creditAccounts []models.CreditAccount
	var creditAccountRepository dal.CreditAccountRepository

	creditAccountRepository.List(&creditAccounts)

	return c.RenderJSON(creditAccounts)
}

func (c CreditAccountController) Get() revel.Result{
	var creditAccount models.CreditAccount
	var creditAccountRepository dal.CreditAccountRepository
	var organizationID, clientId int
	var err error

	organizationID, err = strconv.Atoi(c.Params.Get("orgId"))
	if err != nil {
		return c.RenderJSON(err)
	}

	clientId, err = strconv.Atoi(c.Params.Get("clientId"))

	if err != nil {
		return c.RenderJSON(err)
	}

	creditAccountRepository.Get(organizationID, clientId, &creditAccount)

	return c.RenderJSON(creditAccount)
}

func (c CreditAccountController) Create() revel.Result{
	var creditAccount models.CreditAccount
	var creditAccountRepository dal.CreditAccountRepository

	c.Params.BindJSON(&creditAccount)
	creditAccountRepository.Create(&creditAccount)

	return c.RenderJSON("Your credit account was successfully created");
}

func (c CreditAccountController) Remove(id int) revel.Result{
	var creditAccountRepository dal.CreditAccountRepository

	creditAccountRepository.Delete(id)

	return c.RenderJSON(id)
}

func (c CreditAccountController) Update() revel.Result{
	var creditAccountPut models.CreditAccount
	var creditAccountRepository dal.CreditAccountRepository

	c.Params.BindJSON(&creditAccountPut)
	creditAccountRepository.Update(&creditAccountPut)

	return c.RenderJSON("Your information was updated!")
}
