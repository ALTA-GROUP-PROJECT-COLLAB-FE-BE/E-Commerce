package data

import (
	"project/e-commerce/features/product"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Image       string
	Name        string
	Description string
	Price       int
	Qty         int
	UserID      uint
	User        User
}

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Product  []Product
}

func (data *Product) ToCore() product.Core {
	return product.Core{
		ID:          int(data.ID),
		Image:       data.Image,
		Name:        data.Name,
		Description: data.Description,
		Qty:         data.Qty,
		Price:       data.Price,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
		User: product.User{
			ID:   int(data.UserID),
			Name: data.User.Name,
		},
	}
}

func ToCoreList(data []Product) []product.Core {
	result := []product.Core{}
	for i := range data {
		result = append(result, data[i].ToCore())
	}
	return result
}

func FromCore(core product.Core) Product {
	return Product{
		Image:       core.Image,
		Name:        core.Name,
		Description: core.Description,
		Price:       core.Price,
		Qty:         core.Qty,
		UserID:      uint(core.UserID),
	}
}
