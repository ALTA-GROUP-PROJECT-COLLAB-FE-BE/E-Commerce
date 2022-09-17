package usecase

import (
	"errors"
	"project/e-commerce/features/user"
	"project/e-commerce/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPostData(t *testing.T) {
	repo := new(mocks.UserData)
	returnData := user.Core{Name: "jackma", PPicture: "image-2.jpg", Username: "jackma_", Email: "jackma@gmail.com", Password: "123"}
	t.Run("success post daca", func(t *testing.T) {
		repo.On("InsertData", mock.Anything).Return(1, nil).Once()

		useCase := New(repo)
		res, err := useCase.PostData(returnData)
		assert.NoError(t, err)
		assert.Equal(t, 1, res)
		repo.AssertExpectations(t)
	})
	t.Run("error create data", func(t *testing.T) {

		repo.On("InsertData", mock.Anything).Return(-1, errors.New("there is some error")).Once()

		useCase := New(repo)
		res, err := useCase.PostData(returnData)
		assert.EqualError(t, err, "there is some error")
		assert.Equal(t, -1, res)
		repo.AssertExpectations(t)
	})
	t.Run("empty data", func(t *testing.T) {

		returnData.Name = ""
		returnData.Password = ""
		returnData.Email = ""

		useCase := New(repo)
		res, err := useCase.PostData(returnData)
		assert.EqualError(t, err, "masukan nama ")
		assert.Equal(t, -1, res)
		repo.AssertExpectations(t)
	})

}

func TestGetData(t *testing.T) {
	repo := new(mocks.UserData)
	returnData := user.Core{Name: "jackma", PPicture: "image-2.jpg", Username: "jackma_", Email: "jackma@gmail.com"}
	t.Run("get data success", func(t *testing.T) {
		repo.On("SelectData", mock.Anything).Return(returnData, nil).Once()

		useCase := New(repo)
		res, err := useCase.GetData(1)
		assert.NoError(t, err)
		assert.Equal(t, res, returnData)
		repo.AssertExpectations(t)
	})

	t.Run("error get data", func(t *testing.T) {
		repo.On("SelectData", mock.Anything).Return(user.Core{}, errors.New("error")).Once()
		useCase := New(repo)
		res, err := useCase.GetData(1)
		assert.Error(t, err)
		assert.NotEqual(t, 1, res)
		repo.AssertExpectations(t)
	})
}

func TestDeleteData(t *testing.T) {
	repo := new(mocks.UserData)
	t.Run("success delete data", func(t *testing.T) {

		repo.On("DeleteData", mock.Anything).Return(-1, nil).Once()

		useCase := New(repo)

		delete, err := useCase.DeleteAkun(1)
		assert.Nil(t, err)
		assert.Equal(t, -1, delete)
		repo.AssertExpectations(t)
	})

	t.Run("error delete data", func(t *testing.T) {

		repo.On("DeleteData", mock.Anything).Return(-1, errors.New("error")).Once()

		useCase := New(repo)

		res, err := useCase.DeleteAkun(1)
		assert.Error(t, err)
		assert.Equal(t, -1, res)
		repo.AssertExpectations(t)
	})
}

func TestPostDataId(t *testing.T) {
	repo := new(mocks.UserData)
	returnData := user.Core{
		ID:        0,
		Name:      "jackma",
		PPicture:  "image-2.jpg",
		Username:  "jackma_",
		Email:     "jackma@gmail.com",
		Password:  "asdf",
		CreatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
	}

	t.Run("Update data success", func(t *testing.T) {
		repo.On("UpdateData", mock.Anything, mock.Anything).Return(1, nil).Once()

		useCase := New(repo)
		res, err := useCase.PostDataId(returnData, -1)
		assert.NoError(t, err)
		assert.Equal(t, 1, res)
		repo.AssertExpectations(t)
	})

}
