package usecase

import "project/e-commerce/features/order"

type orderUsercase struct {
	orderData order.DataInterface
}

func New(data order.DataInterface) order.UsecaseInterface {
	return &orderUsercase{
		orderData: data,
	}
}

func (usecase *orderUsercase) PostData(data order.Core) (int, error) {
	row, err := usecase.orderData.InsertData(data)
	return row, err
}

func (usecase *orderUsercase) GetData(id int) ([]order.Core, error) {
	results, err := usecase.orderData.SelectData(id)
	return results, err
}
