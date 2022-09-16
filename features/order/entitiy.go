package order

import "time"

type Core struct {
	ID          int
	PaymentID   int
	UserID      int
	AddressID   int
	TotalPrice  int
	Status      string
	Address     AddressCore
	Payment     PaymentCore
	OrderDetail OrderDetail
	CreatedAt   time.Time
}

type OrderDetail struct {
	ID          int
	OrderID     int
	ProductID   int
	ProductName string
	Price       int
	Qty         int
}
type PaymentCore struct {
	ID          int
	PaymentName string
	CardNumber  string
	PaymentCode string
}

type AddressCore struct {
	ID         int
	City       string
	Province   string
	PostalCode string
	Street     string
}

type UsecaseInterface interface {
	CreateData(data Core, idCart []int) (row int, err error)
	ConfirmStatus(idOrder, idFromToken int) (row int, err error)
	CancelStatus(orderID, idFromToken int) (row int, err error)
	HistoryAll(limitint, offsetint, idFromToken int) (data []Core, err error)
	OrderDetails(orderID int) (data []OrderDetail, err error)
	GetMyDataOrder(limitint, offsetint, idFromToken int) (data []Core, err error)
}

type Data interface {
	InsertOrder(data Core) (orderID int, err error)
	InsertAddress(AddressCore) (addID int, err error)
	InsertPayment(PaymentCore) (payID int, err error)
	TotalPrice(idCart []int) (totalPrice int, err error)
	SelectCart(idCart []int, userID int) (item []OrderDetail, err error)
	InsertOrderDetail(data OrderDetail, orderID int) (row int, err error)
	ConfirmStatusData(idOrder, idFromToken int) (row int, err error)
	CancelStatusData(orderID, idFromToken int) (row int, err error)
	HistoryAllData(limitint, offsetint, idFromToken int) (data []Core, err error)
	OrderDetailDB(orderID int) (data []OrderDetail, err error)
	GetMyDataOrderDB(limitint, offsetint, idFromToken int) (data []Core, err error)
}
