package delivery

import (
	// "be11/apiclean/utils/helper"

	"net/http"
	"project/e-commerce/features/user"

	"project/e-commerce/middlewares"
	"project/e-commerce/utils/helper"
	"strconv"

	// UserResponse "project/e-commerce/features/user/delivery"

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
	e.POST("/users", handler.PostData)                                // Register User
	e.GET("/users/:id", handler.GetData, middlewares.JWTMiddleware()) // Lihat Profile

}

func (delivery *UserDelivery) PostData(c echo.Context) error {
	var dataRequest UserRequest
	errBind := c.Bind(&dataRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("error bind data"))
	}
	row, err := delivery.userUsecase.PostData(toCore(dataRequest))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("error insert data"))
	}
	if row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("error insert data"))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponseHelper("success insert data"))
}

func (delivery *UserDelivery) GetData(c echo.Context) error {

	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("param must be a number"))
	}

	result, err := delivery.userUsecase.GetData(idConv)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("failed to get data"))
	}

	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("succes get data", FromCore(result)))

}

// id := c.Param("id")
// idUser, _ := strconv.Atoi(id)
// result, errGet := h.userBusiness.GetUserById(idUser)
