package data

import (
	"project/e-commerce/features/login"

	"log"

	"gorm.io/gorm"
)

type userLogin struct {
	DB *gorm.DB
}

func New(conn *gorm.DB) login.DataInterface {

	return &userLogin{
		DB: conn,
	}

}

func (repo *userLogin) LoginUser(email, password string) (login.LoginCore, error) {

	var logAuth User
	txEmail := repo.DB.Where("email = ?", email).First(&logAuth)
	if txEmail.Error != nil {
		log.Println("Error tx")
		return login.LoginCore{}, txEmail.Error
	}

	if txEmail.RowsAffected != 1 {
		log.Println("Error row")
		return login.LoginCore{}, txEmail.Error
	}

	var data = toCore(logAuth)

	return data, nil

}
