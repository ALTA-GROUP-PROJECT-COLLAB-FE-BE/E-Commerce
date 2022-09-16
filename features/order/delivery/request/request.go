package request

import "project/e-commerce/features/order"

type Order struct {
	CartID  []int   `json:"cart_id" form:"cart_id"`
	Address Address `json:"address" form:"address"`
	Payment Payment `json:"payment" form:"payment"`
}

type Address struct {
	Street     string `json:"street" form:"street"`
	City       string `json:"city" form:"city"`
	Province   string `json:"province" form:"province"`
	PostalCode string `json:"postal_code" form:"postal_code"`
}

type Payment struct {
	PaymentName string `json:"payment_name" form:"payment_name"`
	CardNumber  string `json:"card_number" form:"card_number"`
	PaymentCode string `json:"payment_code" form:"payment_code"`
}

func ToCore(req Order) order.Core {
	return order.Core{
		Address: order.AddressCore{
			City:       req.Address.City,
			Province:   req.Address.Province,
			PostalCode: req.Address.PostalCode,
			Street:     req.Address.Street,
		},
		Payment: order.PaymentCore{
			PaymentName: req.Payment.PaymentName,
			CardNumber:  req.Payment.CardNumber,
			PaymentCode: req.Payment.PaymentCode,
		},
	}
}
