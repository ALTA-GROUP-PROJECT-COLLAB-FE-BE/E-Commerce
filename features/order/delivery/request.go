package delivery

import "project/e-commerce/features/order"

type OrderRequest struct {
	Address       string `json:"address" form:"address"`
	Phone         int    `json:"phone" form:"phone"`
	PaymentMethod string `json:"payment method" form:"payment method"`
	UserID        uint
	CartID        uint
}

func toCore(data OrderRequest) order.Core {
	return order.Core{
		Address:       data.Address,
		Phone:         data.Phone,
		PaymentMethod: data.PaymentMethod,
		User: order.User{
			ID: data.UserID,
		},
		Cart: order.Cart{
			ID: data.CartID,
		},
	}
}
