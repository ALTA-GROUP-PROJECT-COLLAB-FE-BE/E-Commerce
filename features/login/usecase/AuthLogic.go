package usecase

import (
	"log"
	"project/e-commerce/features/login"
	"project/e-commerce/middlewares"

	"golang.org/x/crypto/bcrypt"
)

type loginUsecase struct {
	dataLogin login.DataInterface
}

func New(data login.DataInterface) login.UsecaseInterface {
	return &loginUsecase{
		dataLogin: data,
	}
}

func (usecase *loginUsecase) LoginAuthorized(email, password string) (string, error) {

	var err error
	if email == "" || password == "" {
		return "", err
	}

	results, errEmail := usecase.dataLogin.LoginUser(email, password)
	if errEmail != nil {
		return "", errEmail
	}

	errPw := bcrypt.CompareHashAndPassword([]byte(results.Password), []byte(password))
	if errPw != nil {
		log.Println("Error pw")
		return "", err
	}

	token, errToken := middlewares.CreateToken(int(results.ID))

	if errToken != nil {
		return "", err
	}

	if token == "" {
		return "", err
	}

	return token, nil

}
