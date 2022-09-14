package usecase

import (
	"errors"
	"project/e-commerce/features/cart"
)

type cartUsecase struct {
	cartData cart.Data
}

func NewCartUsecase(dataCart cart.Data) cart.Business {
	return &cartUsecase{
		cartData: dataCart,
	}
}

func (cu *cartUsecase) GetAllData(limit, offset, idFromToken int) (data []cart.Core, err error) {
	data, err = cu.cartData.SelectData(limit, offset, idFromToken)
	for k, v := range data {
		data[k].TotalPrice = v.Qty * v.Product.Price
	}
	return data, err
}
func (cu *cartUsecase) CreateData(data cart.Core) (row int, err error) {
	if data.Qty == 0 || data.Product.ID == 0 {
		return -1, errors.New("please make sure all field are field in coreectly")
	}
	isExist, idCart, qty, _ := cu.cartData.CheckCart(data.ProductID, data.UserID)
	if isExist {
		row, err = cu.cartData.UpdateDataDB(qty+1, idCart, data.UserID)
	} else {
		row, err = cu.cartData.InsertData(data)
	}
	return row, err
}
func (cu *cartUsecase) UpdateData(qty, idCart, idFromToken int) (row int, err error) {
	row, err = cu.cartData.UpdateDataDB(qty, idCart, idFromToken)
	return row, err
}
func (cu *cartUsecase) DeleteData(idCart, idFromToken int) (row int, err error) {
	row, err = cu.cartData.DeleteDataDB(idCart, idFromToken)
	return row, err
}
