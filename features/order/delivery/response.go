package delivery

import "project/e-commerce/features/order"

type OrderResponse struct {
	ID            uint   `json:"id"`
	Address       string `json:"address"`
	Phone         int    `json:"phone"`
	PaymentMethod string `json:"paymentmethod"`
	User          User
	Cart          Cart
}

type User struct {
	ID uint `json:"id"`
}

type Cart struct {
	ID         uint `json:"id"`
	TotalPrice uint `json:"totalprice"`
}

func fromCore(data order.Core) OrderResponse {
	return OrderResponse{
		ID:            data.ID,
		Address:       data.Address,
		Phone:         data.Phone,
		PaymentMethod: data.PaymentMethod,
		User: User{
			ID: data.UserID,
		},
		Cart: Cart{
			ID: data.CartID,
		},
	}
}

func fromCoreList(data []order.Core) []OrderResponse {
	var dataResponse []OrderResponse
	for _, v := range data {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
