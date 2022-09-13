package delivery

import (
	"project/e-commerce/features/login"
	"project/e-commerce/utils/helper"

	"github.com/labstack/echo/v4"
)

type LoginHandler struct {
	userUsecase login.UsecaseInterface
}

func New(e *echo.Echo, usecase login.UsecaseInterface) {

	handler := LoginHandler{
		userUsecase: usecase,
	}

	e.POST("/login", handler.Login)

}

func (h *LoginHandler) Login(c echo.Context) error {

	var req Request
	errBind := c.Bind(&req)
	if errBind != nil {
		return c.JSON(400, errBind)
	}

	str, err := h.userUsecase.LoginAuthorized(req.Email, req.Password)
	if err != nil {
		return c.JSON(404, err)
	}
	return c.JSON(200, helper.SuccessDataResponseHelper("Login Success", str))

}
