package business

import (
	"errors"
	"project/e-commerce/features/product"
)

type productUseCase struct {
	productData product.Data
}

func NewProductBusiness(dataProduct product.Data) product.Business {
	return &productUseCase{
		productData: dataProduct,
	}
}

func (uc *productUseCase) GetAllData(limit, offset int) (resp []product.Core, err error) {
	resp, err = uc.productData.SelectData(limit, offset)
	return resp, err
}

func (uc *productUseCase) CreateData(input product.Core) (row int, err error) {
	if input.Name == "" || input.Price == 0 || input.Qty == 0 || input.Image == "" || input.Description == "" {
		return -1, errors.New("please make sure all fields are filled in correctly")
	}
	row, err = uc.productData.InsertData(input)
	return row, err
}

func (uc *productUseCase) GetProductById(idProd int) (data product.Core, err error) {
	data, err = uc.productData.SelectDataById(idProd)
	return data, err
}

func (uc *productUseCase) UpdateData(input product.Core, idProd, idFromToken int) (row int, err error) {
	prodReq := map[string]interface{}{}
	if input.Name != "" {
		prodReq["name"] = input.Name
	}
	if input.Price != 0 {
		prodReq["price"] = input.Price
	}
	prodReq["qty"] = input.Qty
	if input.Image != "" {
		prodReq["image"] = input.Image
	}
	if input.Description != "" {
		prodReq["description"] = input.Description
	}
	row, err = uc.productData.UpdateDataDB(prodReq, idProd, idFromToken)
	return row, err
}

func (uc *productUseCase) GetDataByMe(idUser int) (data []product.Core, err error) {
	data, err = uc.productData.GetDataByMeDB(idUser)
	return data, err
}

func (uc *productUseCase) DeleteDataById(idProd, idFromToken int) (row int, err error) {
	row, err = uc.productData.DeleteDataByIdDB(idProd, idFromToken)
	return row, err
}
