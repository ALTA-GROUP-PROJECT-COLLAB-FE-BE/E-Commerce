package response

import (
	"project/e-commerce/features/order"
	"time"
)

type Order struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	TotalPrice  int       `json:"total_price"`
	Status      string    `json:"status"`
	Address     Address   `json:"product"`
	PaymentName string    `json:"payment_name"`
	CreatedAt   time.Time `json:"created_at"`
}

type Address struct {
	Street     string `json:"street"`
	City       string `json:"city"`
	Province   string `json:"province"`
	PostalCode string `json:"postal_code"`
}

func FromCore(data order.Core) Order {
	return Order{
		ID:         data.ID,
		UserID:     data.UserID,
		TotalPrice: data.TotalPrice,
		Status:     data.Status,
		Address: Address{
			Street:     data.Address.Street,
			City:       data.Address.City,
			Province:   data.Address.Province,
			PostalCode: data.Address.PostalCode,
		},
		PaymentName: data.Payment.PaymentName,
		CreatedAt:   data.CreatedAt,
	}
}

func FromCoreList(data []order.Core) []Order {
	result := []Order{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}

type OrderDetail struct {
	ID          int    `json:"id"`
	OrderID     int    `json:"order_id"`
	ProductName string `json:"product_name"`
	Price       int    `json:"price"`
	Qty         int    `json:"qty"`
}

func OrderFromCore(data order.OrderDetail) OrderDetail {
	return OrderDetail{
		ID:          data.ID,
		OrderID:     data.OrderID,
		ProductName: data.ProductName,
		Price:       data.Price,
		Qty:         data.Qty,
	}
}

func OrderFromCoreList(data []order.OrderDetail) []OrderDetail {
	result := []OrderDetail{}
	for key := range data {
		result = append(result, OrderFromCore(data[key]))
	}
	return result
}
