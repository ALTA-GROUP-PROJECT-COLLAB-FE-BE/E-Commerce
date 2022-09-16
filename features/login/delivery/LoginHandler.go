package delivery

import (
	"project/e-commerce/features/login"
	"project/e-commerce/utils/helper"

	"github.com/labstack/echo/v4"
)

type AuthDelivery struct {
	authUsecase login.UsecaseInterface
}

func New(e *echo.Echo, usecase login.UsecaseInterface) {

	handler := AuthDelivery{
		authUsecase: usecase,
	}

	e.POST("/login", handler.Auth)

}

func (h *AuthDelivery) Auth(c echo.Context) error {

	var req AuthRequest
	errBind := c.Bind(&req)
	if errBind != nil {
		return c.JSON(400, helper.FailedResponseHelper("wrong request"))
	}

	str := h.authUsecase.LoginAuthorized(req.Email, req.Password)
	if str == "please input email and password" || str == "email not found" || str == "wrong password" {
		return c.JSON(400, helper.FailedResponseHelper(str))
	} else if str == "failed to created token" {
		return c.JSON(500, helper.FailedResponseHelper(str))
	} else {
		return c.JSON(200, helper.SuccessDataResponseHelper("Login Success", str))
	}

}
