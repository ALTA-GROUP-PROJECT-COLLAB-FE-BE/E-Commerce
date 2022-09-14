package data

import (
	"project/e-commerce/features/cart"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	Qty       int
	Status    string
	UserID    int
	ProductID int
	User      User    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Product   Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Product  []Product
}

type Product struct {
	gorm.Model
	Name        string
	Price       int
	Qty         int
	Image       string
	Description string
	UserID      uint
	Cart        []Cart `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User        User
}

func (data *Cart) toCore() cart.Core {
	return cart.Core{
		ID:        int(data.ID),
		Qty:       data.Qty,
		Status:    data.Status,
		UserID:    data.UserID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		Product: cart.Product{
			ID:          int(data.Product.ID),
			Name:        data.Product.Name,
			Price:       data.Product.Price,
			Qty:         data.Product.Qty,
			Description: data.Product.Description,
		},
	}
}

func ToCoreList(data []Cart) []cart.Core {
	result := []cart.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func FromCore(core cart.Core) Cart {
	return Cart{
		Qty:       core.Qty,
		Status:    core.Status,
		UserID:    core.UserID,
		ProductID: core.Product.ID,
	}
}
