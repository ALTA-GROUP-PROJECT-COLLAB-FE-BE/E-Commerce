package product

import (
	"time"
)

type Core struct {
	ID          int
	Image       string
	Name        string
	Description string
	Price       int
	Qty         int
	UserID      uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	User        User
}

type User struct {
	ID   int
	Name string
}

type UsecaseInterface interface {
	GetAllProduct(limit, offset int) (data []Core, err error) // masukin ke usecase
	// AddProduct(data Core) (row int, err error)
	// GetProductById(id int) (product Core, err error)
}

type DataInterface interface {
	GetProduct(limit, offset int) (data []Core, err error)
	InsertProduct(data Core) (row int, err error)
	SelectProductById(id int) (product Core, err error)
	UpdateProduct(data map[string]interface{}, id, token int) (row int, err error)
	GetMyProduct(id int) (data []Core, err error)
	DeleteMyProduct(id, token int) (row int, err error)
}
