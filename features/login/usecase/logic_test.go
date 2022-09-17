package usecase

import (
	"project/e-commerce/features/login"
	"project/e-commerce/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLogin(t *testing.T) {
	repo := new(mocks.UserLogin)
	dataInput := login.Core{ID: 1, Email: "Hery@gmail.com", Password: "$3b$89$6tSIi7BiTknraN3A9tRX/iuytTycJa/oWI8Ft9KcrZNF3ehbjRtes"}
	t.Run("success password.", func(t *testing.T) {
		repo.On("LoginUser", mock.Anything, mock.Anything).Return(dataInput, nil).Once()

		usecase := New(repo)
		result := usecase.LoginAuthorized("Hery@gmail.com", "asdf")
		assert.NotEqual(t, "please input email and password", result)
		assert.NotEqual(t, "email not found", result)
		assert.NotEqual(t, "wrong password", result)
		assert.NotEqual(t, "error to created token", result)
		repo.AssertExpectations(t)
	})
	t.Run("success login", func(t *testing.T) {
		usecase := New(repo)
		result := usecase.LoginAuthorized("Hery@gmail.com", "")
		assert.Equal(t, "please input email and password", result)
		repo.AssertExpectations(t)
	})
	t.Run("Empty email.", func(t *testing.T) {
		usecase := New(repo)
		result := usecase.LoginAuthorized("", "Hery")
		assert.Equal(t, "please input email and password", result)
		repo.AssertExpectations(t)
	})

}
