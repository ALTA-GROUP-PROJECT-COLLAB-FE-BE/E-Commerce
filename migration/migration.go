package migration

import (
	Product "project/e-commerce/features/product"
	Users "project/e-commerce/features/user"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {

	db.AutoMigrate(&Users.User{})
	db.AutoMigrate(&Product.Product{})

}
