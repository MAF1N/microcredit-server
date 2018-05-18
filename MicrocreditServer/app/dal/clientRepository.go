package dal

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"MicrocreditServer/app/dal/domain/models"
	"errors"
)

type ClientRepository struct{
}

func (cli *ClientRepository) List(clients interface{}) error{
	db, err := gorm.Open(Driver, ConnectionString)
	if err != nil {
		return err
	}
	defer db.Close()
	db.SingularTable(true)

	db.Find(clients)
	return nil
}

func (cli *ClientRepository) GetByEmail(email string, dest *models.Client) error {
	db, err := gorm.Open(Driver, ConnectionString)

	if err != nil {
		return err
	}
	defer db.Close()
	db.SingularTable(true)

	err = db.Where("email = ?", email).First(dest).Error
	if err != nil {
		return err
	}
	return nil
}

func (cli *ClientRepository) Get(ID int, dest interface{}) error{
	db, err := gorm.Open(Driver, ConnectionString)

	if err != nil {
		return err
	}
	defer db.Close()
	db.SingularTable(true)

	db.First(dest, ID)
	return nil
}

func (cli *ClientRepository) Create(data interface{}) error{
	db, err := gorm.Open(Driver, ConnectionString)
	if err != nil {
		return err
	}
	defer db.Close()
	db.SingularTable(true)

	isNew := db.NewRecord(data)
	if !isNew {
		errors.New("Not New")
	}

	db.Create(data)

	isNew = db.NewRecord(data)
	if isNew {
		errors.New("Not added")
	}

	return nil
}

func (cli *ClientRepository) Update(data *models.Client) error{
	var client models.Client
	db, err := gorm.Open(Driver, ConnectionString)

	if err != nil {
		return err
	}
	defer db.Close()

	db.SingularTable(true)

	db.First(&client, data.Id)
	client.Email = data.Email
	client.Phone = data.Phone
	client.Name = data.Name
	client.Surname = data.Surname
	db.Save(&client)


	return nil
}

func (cli *ClientRepository) Delete(ID int) error {
	db, err := gorm.Open(Driver, ConnectionString)
	var client models.Client

	if err != nil {
		return err
	}
	defer db.Close()

	db.SingularTable(true)

	db.First(&client, ID)
	db.Delete(&client)

	return nil
}