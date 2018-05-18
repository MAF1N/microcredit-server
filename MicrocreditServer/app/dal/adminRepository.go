package dal

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"MicrocreditServer/app/dal/domain/models"
	"errors"
)

type AdminRepository struct{
}

func (admRep *AdminRepository) List(admins interface{}) error{
	db, err := gorm.Open(Driver, ConnectionString)
	if err != nil {
		return err
	}
	defer db.Close()
	db.SingularTable(true)

	db.Find(admins)
	return nil
}

func (admRep *AdminRepository) GetByEmail(email string, dest *models.Admin) error {
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

func (admRep *AdminRepository) Get(ID int, dest interface{}) error{
	db, err := gorm.Open(Driver, ConnectionString)

	if err != nil {
		return err
	}
	defer db.Close()
	db.SingularTable(true)

	db.First(dest, ID)
	return nil
}

func (admRep *AdminRepository) Create(data interface{}) error{
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

func (admRep *AdminRepository) Update(data *models.Admin) error{
	var admin models.Admin
	db, err := gorm.Open(Driver, ConnectionString)

	if err != nil {
		return err
	}
	defer db.Close()

	db.SingularTable(true)

	db.First(&admin, data.Id)
	admin.Email = data.Email
	admin.Phone = data.Phone
	admin.Name = data.Name
	admin.Surname = data.Surname
	db.Save(&admin)


	return nil
}

func (admRep *AdminRepository) Delete(ID int) error {
	db, err := gorm.Open(Driver, ConnectionString)
	var admin models.Admin

	if err != nil {
		return err
	}
	defer db.Close()

	db.SingularTable(true)

	db.First(&admin, ID)
	db.Delete(&admin)

	return nil
}