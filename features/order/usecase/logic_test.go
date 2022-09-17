package usecase

import (
	"errors"
	"project/e-commerce/features/order"

	"project/e-commerce/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddOrder(t *testing.T) {
	repo := new(mocks.OrderData)
	// mockOrder := transaction.Core{Quantity: 10, TotalPrice: 100000, OrderStatus: "waiting", CartID: 12}
	result := order.Core{ID: 1, PaymentID: 1, UserID: 1, AddressID: 2, TotalPrice: 20000, Status: "on process"}
	// payment := []order.PaymentCore{{PaymentName: "Mastercard", PaymentCode: "1223", CardNumber: "6123-1121-1134-1125"}}

	t.Run("order success", func(t *testing.T) {

		repo.On("InsertOrder", mock.Anything, mock.Anything).Return(1, errors.New("error")).Once()

		useCase := NewOrderBusiness(repo)

		res, err := useCase.CreateData(result, []int{result.ID})
		assert.Error(t, err)
		assert.Equal(t, -1, res)
		repo.AssertExpectations(t)
	})
}

// func TestInsertAddress(t *testing.T) {

// 	repo := new(mocks.OrderData)
// 	dataInput := order.AddressCore{ID: 1, City: "Depok", Province: "Jawa Barat", PostalCode: "16519", Street: "jalan margonda"}

// 	t.Run("Empty data.", func(t *testing.T) {
// 		// dataInput := cart.Core{ProductID: 0}
// 		repo.On("InsertData", mock.Anything, mock.Anything)
// 		usecase := NewOrderBusiness(repo)
// 		result, err := usecase.CreateData([]int{dataInput.ID})
// 		assert.EqualError(t, err, "please make sure all fields are filled in correctly")
// 		assert.Equal(t, -1, result)
// 		repo.AssertExpectations(t)
// 	})
// 	resAdd := order.AddressCore{City: "Depok", Province: "Jawa Barat", PostalCode: "16519", Street: "jalan margonda"}

// 	t.Run("order success", func(t *testing.T) {

// 		resAdd.New("InsertAddress", mock.Anything, mock.Anything).Return(1, errors.New("error")).Once()

// 		useCase := NewOrderBusiness(repo)

// 		res, err := useCase.CreateData(result, []int{result.ID})
// 		assert.Error(t, err)
// 		assert.Equal(t, -1, res)
// 		repo.AssertExpectations(t)
// 	})
// }
