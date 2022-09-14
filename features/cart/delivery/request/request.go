package request

type Cart struct {
	IdProduct int `json:"product_id" form:"product_id"`
	Qty       int `json:"qty" form:"qty"`
}
