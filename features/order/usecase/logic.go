package usecase

import (
	"errors"
	"fmt"
	"project/e-commerce/features/order"
)

type orderUseCase struct {
	orderData order.Data
}

func NewOrderBusiness(dataorder order.Data) order.UsecaseInterface {
	return &orderUseCase{
		orderData: dataorder,
	}
}

func (uc *orderUseCase) CreateData(data order.Core, idCart []int) (row int, err error) {

	if data.Address.City == "" || data.Address.Street == "" || data.Address.PostalCode == "" || data.Address.Province == "" {
		return -1, errors.New("please make sure all fields are filled in correctly")
	}
	idAddress, errAddress := uc.orderData.InsertAddress(data.Address)
	if errAddress != nil {
		return 0, errAddress
	}

	if data.Payment.PaymentName == "" || data.Payment.CardNumber == "" || data.Payment.PaymentCode == "" {
		return -1, errors.New("please make sure all fields are filled in correctly")
	}
	idPayment, errPayment := uc.orderData.InsertPayment(data.Payment)
	if errPayment != nil {
		return 0, errPayment
	}
	data.PaymentID = idPayment
	data.AddressID = idAddress
	data.Status = "on process"
	totalPrice, errTotalPrice := uc.orderData.TotalPrice(idCart)
	if errTotalPrice != nil {
		return 0, errTotalPrice
	}
	data.TotalPrice = totalPrice
	idOrder, errOrder := uc.orderData.InsertOrder(data)
	if errOrder != nil {
		return 0, errOrder
	}
	resultCart, errCart := uc.orderData.SelectCart(idCart, data.UserID)
	if errCart != nil {
		return 0, errCart
	}
	for i := 0; i < len(resultCart); i++ {
		fmt.Println("result: ", resultCart[i])
		row, err = uc.orderData.InsertOrderDetail(resultCart[i], idOrder)
	}
	return row, err
}

func (uc *orderUseCase) ConfirmStatus(orderID, idFromToken int) (row int, err error) {
	row, err = uc.orderData.ConfirmStatusData(orderID, idFromToken)
	return row, err
}

func (uc *orderUseCase) CancelStatus(orderID, idFromToken int) (row int, err error) {
	row, err = uc.orderData.CancelStatusData(orderID, idFromToken)
	return row, err
}

func (uc *orderUseCase) HistoryAll(limitint, offsetint, idFromToken int) (data []order.Core, err error) {
	data, err = uc.orderData.HistoryAllData(limitint, offsetint, idFromToken)
	return data, err
}

func (uc *orderUseCase) OrderDetails(orderID int) (data []order.OrderDetail, err error) {
	data, err = uc.orderData.OrderDetailDB(orderID)
	return data, err
}

func (uc *orderUseCase) GetMyDataOrder(limitint, offsetint, idFromToken int) (data []order.Core, err error) {
	data, err = uc.orderData.GetMyDataOrderDB(limitint, offsetint, idFromToken)
	return data, err
}
