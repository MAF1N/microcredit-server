package controllers

import (
	"github.com/revel/revel"
	"MicrocreditServer/app/dal/domain/models"
	"MicrocreditServer/app/dal"
)

type BankController struct {
	*revel.Controller
}

func (c BankController) List() revel.Result{
	var banks []models.Bank
	var bankRepository dal.BankRepository

	bankRepository.List(&banks)

	return c.RenderJSON(banks)
}

func (c BankController) Get(id int) revel.Result{
	var bank models.Bank
	var bankRepository dal.BankRepository

	bankRepository.Get(id, &bank)

	return c.RenderJSON(bank)
}

func (c BankController) Create() revel.Result{
	var bank models.Bank
	var bankRepository dal.BankRepository

	c.Params.BindJSON(&bank)
	bankRepository.Create(&bank)

	return c.RenderJSON("Your account was successfully created");
}

func (c BankController) Remove(id int) revel.Result{
	var bankRepository dal.BankRepository

	bankRepository.Delete(id)

	return c.RenderJSON(id)
}

func (c BankController) Update() revel.Result{
	var bankPut models.Bank
	var bankRepository dal.BankRepository

	c.Params.BindJSON(&bankPut)
	bankRepository.Update(&bankPut)

	return c.RenderJSON("Your information was updated!")
}
