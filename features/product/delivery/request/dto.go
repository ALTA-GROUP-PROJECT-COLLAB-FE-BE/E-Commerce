package request

import (
	"project/e-commerce/features/product"
)

type Product struct {
	Name        string `json:"name" form:"name"`
	Price       int    `json:"price" form:"price"`
	Qty         int    `json:"qty" form:"qty"`
	Image       string `json:"image" form:"image"`
	Description string `json:"description" form:"description"`
	UserID      int
}

func ToCore(req Product) product.Core {
	return product.Core{
		Name:        req.Name,
		Price:       req.Price,
		Qty:         req.Qty,
		Image:       req.Image,
		Description: req.Description,
		User: product.User{
			ID: req.UserID,
		},
	}
}
