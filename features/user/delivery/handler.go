package delivery

import (
	// "be11/apiclean/utils/helper"

	"net/http"
	"project/e-commerce/features/user"
	"project/e-commerce/utils/helper"

	"github.com/labstack/echo/v4"
)

type UserDelivery struct {
	userUsecase user.UsecaseInterface
}

// ini komen
func New(e *echo.Echo, usecase user.UsecaseInterface) {
	handler := &UserDelivery{
		userUsecase: usecase,
	}
	e.POST("/users", handler.PostData)
}

func (delivery *UserDelivery) PostData(c echo.Context) error {
	var dataRequest UserRequest
	errBind := c.Bind(&dataRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("error bind data"))
	}
	// var userCoreData = entities.RequestToCore(userData)
	// inputPassByte := []byte(userCoreData.Password)
	// hashPass, _ := bcrypt.GenerateFromPassword(inputPassByte, bcrypt.DefaultCost)
	// userCoreData.Password = string(hashPass)

	row, err := delivery.userUsecase.PostData(toCore(dataRequest))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("error insert data"))
	}
	if row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("error insert data"))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponseHelper("success insert data"))
}
