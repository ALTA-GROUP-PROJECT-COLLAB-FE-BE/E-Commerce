package delivery

import (
	"net/http"
	"project/e-commerce/features/user"

	"project/e-commerce/middlewares"
	"project/e-commerce/utils/helper"
	"strconv"

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
	e.POST("/registration", handler.PostData)                                       // Daftar akun
	e.GET("/readprofile/:id", handler.GetData, middlewares.JWTMiddleware())         // Lihat Profile
	e.PUT("/updateprofile/:id", handler.PostDataId, middlewares.JWTMiddleware())    // Update Akun
	e.DELETE("/deleteaccount/:id", handler.DeleteAkun, middlewares.JWTMiddleware()) // Delte Akun
}

func (delivery *UserDelivery) PostData(c echo.Context) error {
	var dataRequest UserRequest
	errBind := c.Bind(&dataRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("error binding data"))
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

func (delivery *UserDelivery) PostDataId(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	idToken := middlewares.ExtractToken(c)

	if idToken != idConv {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("you dont have acces"))
	}

	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("param must be a number"))
	}

	var dataUpdate UserRequest
	errBind := c.Bind(&dataUpdate)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("error binding data"))
	}

	row, err := delivery.userUsecase.PostDataId(toCore(dataUpdate), idConv)
	if err != nil || row == 0 {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("failed to update data"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("success update data"))
}

func (delivery *UserDelivery) DeleteAkun(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	idToken := middlewares.ExtractToken(c)

	if idToken != idConv {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("you dont have acces"))
	}

	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("param must be a number"))
	}

	row, err := delivery.userUsecase.DeleteAkun(idConv)
	if err != nil || row == 0 {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("failed to delete account"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("succes delte account"))
}
