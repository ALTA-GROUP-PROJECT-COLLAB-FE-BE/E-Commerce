package product

import "time"

type Core struct {
	ID          int
	Name        string
	Price       int
	Qty         int
	Image       string
	Description string
	UserID      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	User        User
}

type User struct {
	ID   int
	Name string
}

type Business interface {
	GetAllData(limit, offset int) (data []Core, err error)
	CreateData(data Core) (row int, err error)
	GetProductById(idProd int) (product Core, err error)
	UpdateData(data Core, idProd, idFromToken int) (row int, err error)
	GetDataByMe(idUser int) (data []Core, err error)
	DeleteDataById(idUser, idFromToken int) (row int, err error)
}

type Data interface {
	SelectData(limit, offset int) (data []Core, err error)
	InsertData(data Core) (row int, err error)
	SelectDataById(idProd int) (product Core, err error)
	UpdateDataDB(data map[string]interface{}, idProd, idFromToken int) (row int, err error)
	GetDataByMeDB(idUser int) (data []Core, err error)
	DeleteDataByIdDB(idUser, idFromToken int) (row int, err error)
}
