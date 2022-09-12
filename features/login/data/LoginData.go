package data

import (
	"project/e-commerce/features/login"

	"time"

	"gorm.io/gorm"
)

type User struct {
	Email     string `gorm:"primary key"`
	ID        uint   `gorm:"autoIncrement"`
	Name      string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func toCore(user User) login.LoginCore {

	var core = login.LoginCore{
		ID:       int(user.ID),
		Email:    user.Email,
		Password: user.Password,
	}

	return core
}
