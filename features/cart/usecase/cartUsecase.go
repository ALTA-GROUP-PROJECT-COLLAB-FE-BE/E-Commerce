package usecase

import (
	"errors"
	Fcart "project/e-commerce/features/cart"
)

type cartUseCase struct {
	cartData Fcart.Data
}

func NewCartBusiness(datacart Fcart.Data) Fcart.Business {
	return &cartUseCase{
		cartData: datacart,
	}
}

func (uc *cartUseCase) GetAllData(limit, offset, idFromToken int) (data []Fcart.Core, err error) {
	data, err = uc.cartData.SelectData(limit, offset, idFromToken)
	for k, v := range data {
		data[k].TotalPrice = v.Qty * v.Product.Price
	}
	return data, err
}

func (uc *cartUseCase) CreateData(data Fcart.Core) (row int, err error) {
	if data.Qty == 0 || data.Product.ID == 0 {
		return -1, errors.New("please make sure all fields are filled in correctly")
	}
	isExist, idCart, qty, _ := uc.cartData.CheckCart(data.Product.ID, data.UserID)
	if isExist {
		row, err = uc.cartData.UpdateDataDB(qty+1, idCart, data.UserID)
	} else {
		row, err = uc.cartData.InsertData(data)
	}

	return row, err
}

func (uc *cartUseCase) UpdateData(qty, idCart, idFromToken int) (row int, err error) {
	row, err = uc.cartData.UpdateDataDB(qty, idCart, idFromToken)
	return row, err
}

func (uc *cartUseCase) DeleteData(idCart, idFromToken int) (row int, err error) {
	row, err = uc.cartData.DeleteDataDB(idCart, idFromToken)
	return row, err
}
