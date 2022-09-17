package usecase

import (
	"errors"
	"project/e-commerce/features/cart"
	"project/e-commerce/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCartPost(t *testing.T) {
	repo := new(mocks.CartData)
	dataInput := cart.Core{ProductID: 1}

	t.Run("Empty data.", func(t *testing.T) {
		dataInput := cart.Core{ProductID: 0}

		usecase := NewCartBusiness(repo)
		result, err := usecase.CreateData(dataInput)
		assert.EqualError(t, err, "please make sure all fields are filled in correctly")
		assert.Equal(t, -1, result)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Insert data.", func(t *testing.T) {
		// repo.On("InsertData", mock.Anything).Return(1, errors.New("error")).Once()
		// dataInput := cart.Core{ProductID: 1}

		usecase := NewCartBusiness(repo)
		result, err := usecase.CreateData(dataInput)
		assert.Error(t, err)
		assert.Equal(t, -1, result)
		repo.AssertExpectations(t)
	})

}

func TestDeleteCart(t *testing.T) {
	repo := new(mocks.CartData)

	t.Run("Success Delete data.", func(t *testing.T) {
		repo.On("DeleteDataDB", mock.Anything, mock.Anything).Return(1, nil).Once()

		usecase := NewCartBusiness(repo)

		result, err := usecase.DeleteData(1, 1)
		assert.Nil(t, err)
		assert.Equal(t, 1, result)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Delete data", func(t *testing.T) {
		repo.On("DeleteDataDB", mock.Anything, mock.Anything).Return(-1, errors.New("Error")).Once()

		usecase := NewCartBusiness(repo)

		result, err := usecase.DeleteData(1, 1)
		assert.Error(t, err)
		assert.Equal(t, -1, result)
		repo.AssertExpectations(t)
	})

}

func TestGetAll(t *testing.T) {
	repo := new(mocks.CartData)
	dataInput := []cart.Core{{ID: 1, Status: "on process", ProductID: 1, UserID: 1, TotalPrice: 200000, Qty: 1}}
	t.Run("Success Get data by Token.", func(t *testing.T) {
		repo.On("SelectData", mock.Anything, mock.Anything, mock.Anything).Return(dataInput, nil).Once()
		usecase := NewCartBusiness(repo)
		result, err := usecase.GetAllData(0, 1, 2)
		assert.NoError(t, err)
		assert.Equal(t, dataInput, result)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Get data by Token.", func(t *testing.T) {
		repo.On("SelectData", mock.Anything, mock.Anything, mock.Anything).Return([]cart.Core{}, errors.New("Error")).Once()

		usecase := NewCartBusiness(repo)
		result, err := usecase.GetAllData(1, 1, 1)
		assert.Error(t, err)
		assert.NotEqual(t, 1, result)
		repo.AssertExpectations(t)
	})

}
