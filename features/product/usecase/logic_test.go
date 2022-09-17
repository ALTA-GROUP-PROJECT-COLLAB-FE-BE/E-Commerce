package usecase

import (
	"errors"
	"project/e-commerce/features/product"
	"project/e-commerce/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddProduct(t *testing.T) {
	repo := new(mocks.ProductData)
	dataInput := product.Core{
		ID:          0,
		Name:        "Nike Air Max",
		Price:       1000000,
		Qty:         1,
		Image:       "sepatu.jpg",
		Description: "Sneakers Casual Original",
		UserID:      1,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		User:        product.User{},
	}

	t.Run("Success insert data", func(t *testing.T) {
		repo.On("InsertData", mock.Anything).Return(1, nil).Once()

		usecase := NewProductBusiness(repo)
		result, err := usecase.CreateData(dataInput)
		assert.NoError(t, err)
		assert.Equal(t, 1, result)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Insert data.", func(t *testing.T) {
		repo.On("InsertData", mock.Anything).Return(1, errors.New("error")).Once()

		usecase := NewProductBusiness(repo)
		result, err := usecase.CreateData(dataInput)
		assert.Error(t, err)
		assert.Equal(t, 1, result)
		repo.AssertExpectations(t)
	})
	t.Run("Empty data.", func(t *testing.T) {
		dataInput := product.Core{
			ID:          0,
			Name:        "",
			Price:       0,
			Qty:         0,
			Image:       "",
			Description: "",
			UserID:      0,
			CreatedAt:   time.Time{},
			UpdatedAt:   time.Time{},
			User:        product.User{},
		}

		usecase := NewProductBusiness(repo)
		result, err := usecase.CreateData(dataInput)
		assert.EqualError(t, err, "please make sure all fields are filled in correctly")
		assert.Equal(t, -1, result)
		repo.AssertExpectations(t)
	})
}

func TestGetAll(t *testing.T) {
	repo := new(mocks.ProductData)
	getData := []product.Core{{
		ID:          0,
		Name:        "Nike Air Max",
		Price:       1000000,
		Qty:         1,
		Image:       "sepatu.jpg",
		Description: "Sneakers Casual Original",
		UserID:      1,
	}}
	t.Run("Success Get all data.", func(t *testing.T) {
		repo.On("SelectData", mock.Anything, mock.Anything).Return(getData, nil).Once()

		usecase := NewProductBusiness(repo)
		result, err := usecase.GetAllData(1, 1)
		assert.NoError(t, err)
		assert.Equal(t, getData, result)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Get all data.", func(t *testing.T) {
		repo.On("SelectData", mock.Anything, mock.Anything).Return([]product.Core{}, errors.New("error")).Once()

		usecase := NewProductBusiness(repo)
		result, err := usecase.GetAllData(1, 1)
		assert.Error(t, err)
		assert.Equal(t, []product.Core([]product.Core{}), result)
		repo.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	repo := new(mocks.ProductData)
	dataId := product.Core{
		ID:          0,
		Name:        "Nike Air Max",
		Price:       1000000,
		Qty:         1,
		Image:       "sepatu.jpg",
		Description: "Sneakers Casual Original",
		UserID:      1,
	}
	t.Run("Success Get data by Id.", func(t *testing.T) {
		repo.On("SelectDataById", mock.Anything).Return(dataId, nil).Once()

		usecase := NewProductBusiness(repo)
		result, err := usecase.GetProductById(1)
		assert.NoError(t, err)
		assert.Equal(t, dataId, result)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Get data by Id.", func(t *testing.T) {
		repo.On("SelectDataById", mock.Anything).Return(product.Core{}, errors.New("Error")).Once()

		usecase := NewProductBusiness(repo)
		result, err := usecase.GetProductById(1)
		assert.Error(t, err)
		assert.NotEqual(t, 1, result)
		repo.AssertExpectations(t)
	})

}
func TestUpdateData(t *testing.T) {
	repo := new(mocks.ProductData)
	upData := product.Core{
		Name:        "Nike Air Max",
		Price:       1000000,
		Qty:         1,
		Image:       "sepatu.jpg",
		Description: "Sneakers Casual Original",
	}
	t.Run("Success Update data", func(t *testing.T) {
		repo.On("UpdateDataDB", mock.Anything, mock.Anything, mock.Anything).Return(1, nil).Once()

		usecase := NewProductBusiness(repo)

		result, err := usecase.UpdateData(upData, 1, 1)
		assert.NoError(t, err)
		assert.Equal(t, 1, result)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Update data", func(t *testing.T) {
		repo.On("UpdateDataDB", mock.Anything, mock.Anything, mock.Anything).Return(-1, errors.New("Error")).Once()

		usecase := NewProductBusiness(repo)

		result, err := usecase.UpdateData(upData, 1, 1)
		assert.Error(t, err)
		assert.Equal(t, -1, result)
		repo.AssertExpectations(t)
	})

}

func TestDeleteData(t *testing.T) {
	repo := new(mocks.ProductData)
	t.Run("Success Delete data.", func(t *testing.T) {
		repo.On("DeleteDataByIdDB", mock.Anything, mock.Anything).Return(1, nil).Once()

		usecase := NewProductBusiness(repo)

		result, err := usecase.DeleteDataById(1, 1)
		assert.Nil(t, err)
		assert.Equal(t, 1, result)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Delete data", func(t *testing.T) {
		repo.On("DeleteDataByIdDB", mock.Anything, mock.Anything).Return(-1, errors.New("Error")).Once()

		usecase := NewProductBusiness(repo)

		result, err := usecase.DeleteDataById(1, 1)
		assert.Error(t, err)
		assert.Equal(t, -1, result)
		repo.AssertExpectations(t)
	})
}
