package cart

import (
	"time"
)

type Core struct {
	ID         int
	Qty        int
	Status     string
	TotalPrice int
	UserID     int
	ProductID  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Product    Product
	User       User
}

type Product struct {
	ID          int
	Name        string
	Price       int
	Qty         int
	Description string
	UserID      int
	User        User
}

type User struct {
	ID    int
	Name  int
	Email int
}

type Business interface {
	GetAllData(limit, offset, idFromToken int) (data []Core, err error)
	CreateData(data Core) (row int, err error)
	UpdateData(qty, idCart, idFromToken int) (row int, err error)
	DeleteData(idProd, idFromToken int) (row int, err error)
}

type Data interface {
	SelectData(limit, offset, idFromToken int) (data []Core, err error)
	InsertData(data Core) (row int, err error)
	UpdateDataDB(qty, idCart, idFromToken int) (row int, err error)
	DeleteDataDB(idProd, idFromToken int) (row int, err error)
	CheckCart(idProd, idFromToken int) (isExist bool, idCart, qty int, err error)
}
