package data

import (
	product "project/e-commerce/features/cart/data"
	"project/e-commerce/features/order"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	PaymentID   int
	UserID      int
	TotalPrice  int
	AddressID   int
	Status      string
	Address     Address
	Payment     Payment
	OrderDetail []OrderDetail
}

type OrderDetail struct {
	ID          int `gorm:"autoIncrement"`
	OrderID     int
	ProductID   int
	ProductName string `gorm:"column:product_name"`
	Price       int
	Qty         int
}

type Cart struct {
	gorm.Model
	Qty       int
	Status    string
	ProductID int
	UserID    int
	Product   product.Product
}

type Address struct {
	ID         int `gorm:"autoIncrement"`
	City       string
	Province   string
	PostalCode string
	Street     string
}

type Payment struct {
	ID          int `gorm:"autoIncrement"`
	PaymentName string
	CardNumber  string
	PaymentCode string
}

func (data Order) toCore() order.Core {
	return order.Core{
		ID:         int(data.ID),
		TotalPrice: data.TotalPrice,
		UserID:     data.UserID,
		Status:     data.Status,
		Address: order.AddressCore{
			City:       data.Address.City,
			Province:   data.Address.Province,
			PostalCode: data.Address.PostalCode,
			Street:     data.Address.Street,
		},
		Payment: order.PaymentCore{
			PaymentName: data.Payment.PaymentName,
		},
		CreatedAt: data.CreatedAt,
	}
}

func ToCoreList(data []Order) []order.Core {
	result := []order.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func toOrderDetailCore(data OrderDetail) order.OrderDetail {
	return order.OrderDetail{
		ID:          int(data.ID),
		OrderID:     data.OrderID,
		ProductName: data.ProductName,
		Price:       data.Price,
		Qty:         data.Qty,
	}
}
func ToOrderDetailCoreList(data []OrderDetail) []order.OrderDetail {
	result := []order.OrderDetail{}
	for key := range data {
		result = append(result, toOrderDetailCore(data[key]))
	}
	return result
}

func (data *Cart) toOrderDetailFromCart() order.OrderDetail {
	return order.OrderDetail{
		ID:        int(data.ID),
		Qty:       data.Qty,
		ProductID: int(data.Product.ID),
		Price:     data.Product.Price,
	}
}

func ToOrderDetailCoreListFromCart(data []Cart) []order.OrderDetail {
	result := []order.OrderDetail{}
	for key := range data {
		result = append(result, data[key].toOrderDetailFromCart())
	}
	return result
}

func fromCore(core order.Core) Order {
	return Order{
		PaymentID:  core.PaymentID,
		AddressID:  core.AddressID,
		TotalPrice: core.TotalPrice,
		UserID:     core.UserID,
		Status:     core.Status,
	}
}

func FromOrderDetailCoreList(data []order.OrderDetail) []OrderDetail {
	result := []OrderDetail{}
	for key := range data {
		result = append(result, fromOrderDetailCore(data[key]))
	}

	return result
}

func fromAddressCore(addressCore order.AddressCore) Address {
	return Address{
		City:       addressCore.City,
		Province:   addressCore.Province,
		PostalCode: addressCore.PostalCode,
		Street:     addressCore.Street,
	}
}

func fromPaymentCore(paymentCore order.PaymentCore) Payment {
	return Payment{
		CardNumber:  paymentCore.CardNumber,
		PaymentCode: paymentCore.PaymentCode,
		PaymentName: paymentCore.PaymentName,
	}
}

func fromOrderDetailCore(orderDetail order.OrderDetail) OrderDetail {
	return OrderDetail{
		ProductID:   orderDetail.ProductID,
		ProductName: orderDetail.ProductName,
		Price:       orderDetail.Price,
		Qty:         orderDetail.Qty,
	}
}
