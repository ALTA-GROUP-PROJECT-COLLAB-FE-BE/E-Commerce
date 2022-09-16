package migration

import (
	cartModel "project/e-commerce/features/cart/data"
	orderModel "project/e-commerce/features/order/data"
	Product "project/e-commerce/features/product/data"
	userModel "project/e-commerce/features/user/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {

	db.AutoMigrate(&userModel.User{})
	db.AutoMigrate(&Product.Product{})
	db.AutoMigrate(&cartModel.Cart{})
	db.AutoMigrate(&orderModel.Order{})
	db.AutoMigrate(&orderModel.Order{})
	db.AutoMigrate(&orderModel.OrderDetail{})
	db.AutoMigrate(&orderModel.Address{})
	db.AutoMigrate(&orderModel.Payment{})

}
