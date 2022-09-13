package migration

import (
	_mProducts "project/e-commerce/features/product/data"
	userModel "project/e-commerce/features/user/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {

	db.AutoMigrate(&userModel.User{})
	// db.AutoMigrate(&Product.Product{})
	db.AutoMigrate(&_mProducts.Product{})

}
