package controllers

import (
	"github.com/revel/revel"
	"MicrocreditServer/app/dal/domain/models"
	"MicrocreditServer/app/dal"
)

type OrganizationController struct {
	*revel.Controller
}

func (c OrganizationController) List() revel.Result{
	var Organizations []models.Organization
	var OrganizationRepository dal.OrganizationRepository

	token := c.Request.Header.Get("Authorization")[7:]

	tokenValid, err := ValidateToken(token)
	if err != nil {
		return c.RenderJSON(err)
	}

	if !tokenValid {
		return c.RenderJSON("Token is not valid!")
	}

	OrganizationRepository.List(&Organizations)

	return c.RenderJSON(Organizations)
}

func (c OrganizationController) Get(id int) revel.Result{
	var organization models.Organization
	var OrganizationRepository dal.OrganizationRepository

	OrganizationRepository.Get(id, &organization)

	return c.RenderJSON(organization)
}

func (c OrganizationController) Create() revel.Result{
	var organization models.Organization
	var OrganizationRepository dal.OrganizationRepository

	c.Params.BindJSON(&organization)
	OrganizationRepository.Create(&organization)

	return c.RenderJSON("Your account was successfully created");
}

func (c OrganizationController) Remove(id int) revel.Result{
	var organizationRepository dal.OrganizationRepository

	organizationRepository.Delete(id)

	return c.RenderJSON(id)
}

func (c OrganizationController) Update() revel.Result{
	var organization models.Organization
	var OrganizationRepository dal.OrganizationRepository

	c.Params.BindJSON(&organization)
	OrganizationRepository.Update(&organization)

	return c.RenderJSON("Your organization information was updated!")
}
