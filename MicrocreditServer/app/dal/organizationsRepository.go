package dal

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"MicrocreditServer/app/dal/domain/models"
	"errors"
)

type OrganizationRepository struct{
}

func (cli *OrganizationRepository) List(organizations interface{}) error{

	db, err := gorm.Open(Driver, ConnectionString)
	if err != nil {
		return err
	}
	defer db.Close()

	db.SingularTable(true)

	db.Find(organizations)

	return nil
}

func (cli *OrganizationRepository) Get(ID int, organization interface{}) error{
	db, err := gorm.Open(Driver, ConnectionString)
	if err != nil {
		return err
	}
	defer db.Close()

	db.SingularTable(true)

	db.First(organization, ID)

	return nil
}

func (cli *OrganizationRepository) Create(data interface{}) error{
	db, err := gorm.Open(Driver, ConnectionString)
	if err != nil {
		return err
	}
	defer db.Close()

	db.SingularTable(true)

	isNew := db.NewRecord(data)
	if !isNew {
		errors.New("Organization exists")
	}

	db.Create(data)

	isNew = db.NewRecord(data)
	if isNew {
		errors.New("Organization not added")
	}

	return nil
}

func (cli *OrganizationRepository) Update(data *models.Organization) error{
	var organization models.Organization

	db, err := gorm.Open(Driver, ConnectionString)
	if err != nil {
		return err
	}
	defer db.Close()

	db.SingularTable(true)

	db.First(&organization, data.Id)

	organization.Title = data.Title
	organization.Type = data.Type
	organization.BankId = data.Id
	organization.Amount = data.Amount

	db.Save(&organization)

	return nil
}

func (cli *OrganizationRepository) Delete(ID int) error {
	var organization models.Organization

	db, err := gorm.Open(Driver, ConnectionString)

	if err != nil {
		return err
	}
	defer db.Close()

	db.SingularTable(true)

	db.First(&organization, ID)
	db.Delete(&organization)

	return nil
}