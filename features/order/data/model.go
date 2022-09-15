package data

import (
	"project/e-commerce/features/order"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Address       string
	Phone         int
	PaymentMethod string
	UserID        uint
	CartID        uint
	User          User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Cart          Cart `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Order    []Order `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Cart struct {
	gorm.Model
	Qty        uint
	TotalPrice uint
	Order      []Order `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func fromCore(dataCore order.Core) Order {
	dataModel := Order{
		Address:       dataCore.Address,
		Phone:         dataCore.Phone,
		PaymentMethod: dataCore.PaymentMethod,
		UserID:        uint(dataCore.UserID),
		CartID:        uint(dataCore.CartID),
	}
	return dataModel
}

func (data *Order) toCore() order.Core {
	return order.Core{
		Address:       data.Address,
		Phone:         data.Phone,
		PaymentMethod: data.PaymentMethod,
		User: order.User{
			ID: uint(data.UserID),
		},
		Cart: order.Cart{
			ID: uint(data.CartID),
		},
	}
}

func toCoreList(data []Order) []order.Core {
	var dataCore []order.Core
	for key := range data {
		dataCore = append(dataCore, data[key].toCore())
	}
	return dataCore
}
