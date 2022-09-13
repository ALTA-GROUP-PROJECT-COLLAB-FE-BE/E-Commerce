package user

import (
	"time"
)

type Core struct {
	ID        uint
	Name      string
	PPicture  string
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UsecaseInterface interface {
	PostData(data Core) (row int, err error)
	GetData(id int) (data Core, err error)
}

type DataInterface interface {
	InsertData(data Core) (row int, err error)
	SelectData(id int) (data Core, err error)
}
