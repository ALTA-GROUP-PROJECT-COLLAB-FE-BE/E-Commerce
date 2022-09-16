package response

import Fcart "project/e-commerce/features/cart"

type Cart struct {
	ID         int    `json:"id"`
	Qty        int    `json:"qty"`
	TotalPrice int    `json:"total_price"`
	UserID     int    `json:"user_id"`
	Status     string `json:"status"` // <== Baru di tambahin
	Product    Product
}

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Qty         int    `json:"qty"`
	Description string `json:"description"`
}

func FromCore(data Fcart.Core) Cart {
	return Cart{
		ID:         data.ID,
		Qty:        data.Qty,
		TotalPrice: data.TotalPrice,
		UserID:     data.UserID,
		Status:     data.Status, // <== Baru di tambahin
		Product: Product{
			ID:          data.Product.ID,
			Name:        data.Product.Name,
			Price:       data.Product.Price,
			Qty:         data.Product.Qty,
			Description: data.Product.Description,
		},
	}
}

func FromCoreList(data []Fcart.Core) []Cart {
	result := []Cart{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}
