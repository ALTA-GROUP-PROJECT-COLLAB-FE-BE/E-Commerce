package order

import "time"

type Core struct {
	ID            uint
	Address       string
	Phone         int
	PaymentMethod string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	UserID        uint
	CartID        uint
	User          User
	Cart          Cart
}

type User struct {
	ID   uint
	Name string
}

type Cart struct {
	ID         uint
	Qty        uint
	TotalPrice uint
}

type UsecaseInterface interface {
	PostData(data Core) (row int, err error)
	GetData(id int) (data []Core, err error)
}

type DataInterface interface {
	InsertData(data Core) (row int, err error)
	SelectData(id int) (data []Core, err error)
}
