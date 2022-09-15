package delivery

import (
	"net/http"
	"project/e-commerce/features/order"
	"project/e-commerce/middlewares"
	"project/e-commerce/utils/helper"
	"strconv"

	"github.com/labstack/echo/v4"
)

type OrderDelivery struct {
	orderUsecase order.UsecaseInterface
}

func New(e *echo.Echo, usecase order.UsecaseInterface) {
	handler := &OrderDelivery{
		orderUsecase: usecase,
	}
	e.POST("/checkout", handler.PostData)                                    // membuat pesanan
	e.GET("/orderhistory:/id", handler.GetData, middlewares.JWTMiddleware()) //order history
}

func (delivery *OrderDelivery) PostData(c echo.Context) error {
	var dataRequest OrderRequest
	errBind := c.Bind(&dataRequest)

	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("error binding data"))
	}

	row, err := delivery.orderUsecase.PostData(toCore(dataRequest))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("error insert data"))
	}
	if row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("error insert data"))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponseHelper("success insert data"))
}

func (delivery *OrderDelivery) GetData(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	idtoken := middlewares.ExtractToken(c)

	if idtoken != idConv {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("you dont have acces"))
	}
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("param must be a number"))
	}

	results, err := delivery.orderUsecase.GetData(idConv)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("error get data"))
	}

	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("success", fromCoreList(results)))
}
