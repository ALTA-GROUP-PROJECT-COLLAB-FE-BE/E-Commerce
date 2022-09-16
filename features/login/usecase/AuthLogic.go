package usecase

import (
	"project/e-commerce/features/login"
	"project/e-commerce/middlewares"
)

type authUsecase struct {
	authData login.DataInterface
}

func New(data login.DataInterface) login.UsecaseInterface {
	return &authUsecase{
		authData: data,
	}
}

func (usecase *authUsecase) LoginAuthorized(email, password string) string {

	if email == "" || password == "" {
		return "please input email and password"
	}

	results, errEmail := usecase.authData.LoginUser(email)
	if errEmail != nil {
		return "email not found"
	}

	token, errToken := middlewares.CreateToken(int(results.ID))

	if errToken != nil {
		return "error to created token"
	}

	return token

}
