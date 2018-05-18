package dal

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"MicrocreditServer/app/dal/domain/models"
	"errors"
)

type CreditAccountRepository struct{
}

func (cli *CreditAccountRepository) List(creditAccounts interface{}) error{

	db, err := gorm.Open(Driver, ConnectionString)

	if err != nil {
		return err
	}
	defer db.Close()

	db.SingularTable(true)
	db.Find(creditAccounts)
	return nil
}

func (cli *CreditAccountRepository) Get(organizationID int, userID int, dest interface{}) error{
	db, err := gorm.Open(Driver, ConnectionString)

	if err != nil {
		return err
	}
	defer db.Close()

	db.SingularTable(true)
	db.Where("organization_id = ? AND client_id = ?", organizationID, userID).First(dest)

	return nil
}

func (cli *CreditAccountRepository) Create(data interface{}) error{
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

func (cli *CreditAccountRepository) Update(data *models.CreditAccount) error{
	var creditAccount models.CreditAccount
	db, err := gorm.Open(Driver, ConnectionString)
	if err != nil {
		return err
	}
	defer db.Close()
	db.SingularTable(true)

	db.First(&creditAccount, data.OrganizationId, data.ClientId)
	creditAccount.Amount = data.Amount
	creditAccount.CardId = data.CardId
	db.Save(&creditAccount)


	return nil
}
func (cli *CreditAccountRepository) Delete(ID int) error {
	var creditAccount models.CreditAccount
	db, err := gorm.Open(Driver, ConnectionString)

	if err != nil {
		return err
	}
	defer db.Close()
	db.SingularTable(true)

	db.First(&creditAccount, ID)
	db.Delete(&creditAccount)

	return nil
}