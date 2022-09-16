package delivery

import (
	"fmt"
	"net/http"
	"project/e-commerce/features/order"
	Ordereq "project/e-commerce/features/order/delivery/request"
	Orderes "project/e-commerce/features/order/delivery/response"
	middlewares "project/e-commerce/middlewares"

	"project/e-commerce/utils/helper"
	"strconv"

	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	orderBusiness order.UsecaseInterface
}

func NewOrderHandler(e *echo.Echo, orderBusiness order.UsecaseInterface) {
	handler := &OrderHandler{
		orderBusiness: orderBusiness,
	}
	e.POST("/orders", handler.PostOrder, middlewares.JWTMiddleware())
	e.PUT("/orders/:id/confirm", handler.Confirmed, middlewares.JWTMiddleware())
	e.PUT("/orders/:id/cancel", handler.Cancelled, middlewares.JWTMiddleware())
	e.GET("/orders/history", handler.History, middlewares.JWTMiddleware())
	e.GET("/orderdetail/:id", handler.OrderDetail)
	e.GET("/orders/me", handler.GetMyOrder, middlewares.JWTMiddleware())
}

func (h *OrderHandler) PostOrder(c echo.Context) error {
	orderReq := Ordereq.Order{}
	idFromToken := middlewares.ExtractToken(c)
	err := c.Bind(&orderReq)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("failed to bind data, check your input"))
	}
	dataOrder := Ordereq.ToCore(orderReq)
	dataOrder.UserID = idFromToken
	row, errCreate := h.orderBusiness.CreateData(dataOrder, orderReq.CartID)
	fmt.Println("row: ", row)
	if row == -1 {
		return c.JSON(400, helper.FailedResponseHelper("please make sure all fields are filled in correctly"))
	}

	if errCreate != nil {
		return c.JSON(400, helper.FailedResponseHelper("failed to create product, check your input"))
	}

	return c.JSON(200, helper.SuccessResponseHelper("succes order , please confirm your order"))
}

func (h *OrderHandler) Confirmed(c echo.Context) error {
	id := c.Param("id")
	idOrder, _ := strconv.Atoi(id)
	idFromToken := middlewares.ExtractToken(c)
	row, errCon := h.orderBusiness.ConfirmStatus(idOrder, idFromToken)
	if errCon != nil {
		return c.JSON(400, helper.FailedResponseHelper("you dont have access"))
	}
	if row == 0 {
		return c.JSON(400, helper.FailedResponseHelper("failed to update data"))
	}
	return c.JSON(200, helper.SuccessResponseHelper("order confirmed, your order will arrive to your address"))
}

func (h *OrderHandler) Cancelled(c echo.Context) error {
	id := c.Param("id")
	idOrder, _ := strconv.Atoi(id)
	idFromToken := middlewares.ExtractToken(c)
	row, errCon := h.orderBusiness.CancelStatus(idOrder, idFromToken)
	if errCon != nil {
		return c.JSON(400, helper.FailedResponseHelper("you dont have access"))
	}
	if row == 0 {
		return c.JSON(400, helper.FailedResponseHelper("failed to update data"))
	}
	return c.JSON(200, helper.FailedResponseHelper("Order canceled, please order again"))
}

func (h *OrderHandler) History(c echo.Context) error {
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")
	limitint, _ := strconv.Atoi(limit)
	offsetint, _ := strconv.Atoi(offset)
	idFromToken := middlewares.ExtractToken(c)
	result, err := h.orderBusiness.HistoryAll(limitint, offsetint, idFromToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("failed to get all data"))
	}
	return c.JSON(200, helper.SuccessDataResponseHelper("Successfully view your order history", Orderes.FromCoreList(result)))
}

func (h *OrderHandler) OrderDetail(c echo.Context) error {
	id := c.Param("id")
	idOrder, _ := strconv.Atoi(id)
	result, err := h.orderBusiness.OrderDetails(idOrder)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("failed to get all data"))
	}
	return c.JSON(200, helper.SuccessDataResponseHelper("Successfully view your order detail", Orderes.OrderFromCoreList(result)))
}

func (h *OrderHandler) GetMyOrder(c echo.Context) error {
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")
	limitint, _ := strconv.Atoi(limit)
	offsetint, _ := strconv.Atoi(offset)
	idFromToken := middlewares.ExtractToken(c)
	result, err := h.orderBusiness.GetMyDataOrder(limitint, offsetint, idFromToken)
	if err != nil {
		return c.JSON(402, helper.FailedResponseHelper("failed to get all data"))
	}
	return c.JSON(200, helper.SuccessDataResponseHelper("Successfully view your order", Orderes.FromCoreList(result)))
}
