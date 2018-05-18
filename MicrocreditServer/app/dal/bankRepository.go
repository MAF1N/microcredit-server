package dal

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"MicrocreditServer/app/dal/domain/models"
	"errors"
)

type BankRepository struct{
}

func (cli *BankRepository) List(banks interface{}) error{

	db, err := gorm.Open(Driver, ConnectionString)

	if err != nil {
		return err
	}
	defer db.Close()

	db.SingularTable(true)
	db.Find(banks)
	return nil
}

func (cli *BankRepository) Get(ID int, dest interface{}) error{
	db, err := gorm.Open(Driver, ConnectionString)

	if err != nil {
		return err
	}
	defer db.Close()

	db.SingularTable(true)
	db.First(dest, ID)

	return nil
}

func (cli *BankRepository) Create(data interface{}) error{
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

func (cli *BankRepository) Update(data *models.Bank) error{
	var bank models.Bank
	db, err := gorm.Open(Driver, ConnectionString)
	if err != nil {
		return err
	}
	defer db.Close()
	db.SingularTable(true)

	db.First(&bank, data.Id)
	bank.Title = data.Title
	bank.Address = data.Address
	bank.ContactPhone = data.ContactPhone
	db.Save(&bank)


	return nil
}
func (cli *BankRepository) Delete(ID int) error {
	var bank models.Bank
	db, err := gorm.Open(Driver, ConnectionString)

	if err != nil {
		return err
	}
	defer db.Close()
	db.SingularTable(true)

	db.First(&bank, ID)
	db.Delete(&bank)

	return nil
}