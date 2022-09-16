package delivery

import (
	"net/http"
	Fcart "project/e-commerce/features/cart"
	_requestCart "project/e-commerce/features/cart/delivery/request"
	_responseCart "project/e-commerce/features/cart/delivery/response"
	middlewares "project/e-commerce/middlewares"
	_helper "project/e-commerce/utils/helper"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CartHandler struct {
	cartBusiness Fcart.Business
}

func NewCartHandler(e *echo.Echo, cartBusiness Fcart.Business) {
	handler := &CartHandler{
		cartBusiness: cartBusiness,
	}
	e.GET("/cart", handler.GetAll, middlewares.JWTMiddleware())
	e.POST("/cart", handler.PostCart, middlewares.JWTMiddleware())
	e.PUT("/cart/:id", handler.UpdateCart, middlewares.JWTMiddleware())
	e.DELETE("/cart/:id", handler.DeleteCart, middlewares.JWTMiddleware())

}
func (h *CartHandler) GetAll(c echo.Context) error {
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")
	limitint, _ := strconv.Atoi(limit)
	offsetint, _ := strconv.Atoi(offset)
	idFromToken := middlewares.ExtractToken(c)
	result, err := h.cartBusiness.GetAllData(limitint, offsetint, idFromToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.FailedResponseHelper("failed to get all data"))
	}

	return c.JSON(http.StatusOK, _helper.SuccessDataResponseHelper("success get all carts", _responseCart.FromCoreList(result)))
}

func (h *CartHandler) PostCart(c echo.Context) error {
	idFromToken := middlewares.ExtractToken(c)
	cartReq := _requestCart.Cart{}
	err := c.Bind(&cartReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.FailedResponseHelper("failed to bind data, check your input"))
	}
	dataCart := Fcart.Core{}
	dataCart.Product.ID = cartReq.IdProduct
	dataCart.Qty = cartReq.Qty
	dataCart.Status = "on process"
	dataCart.UserID = idFromToken
	row, errCreate := h.cartBusiness.CreateData(dataCart)
	if row == -1 {
		return c.JSON(http.StatusBadRequest, _helper.FailedResponseHelper("please make sure all fields are filled in correctly"))
	}
	if errCreate != nil {
		return c.JSON(http.StatusInternalServerError, _helper.FailedResponseHelper("failed to add to cart"))
	}

	return c.JSON(http.StatusOK, _helper.SuccessResponseHelper("success add to cart"))
}

func (h *CartHandler) UpdateCart(c echo.Context) error {
	id := c.Param("id")
	idCart, _ := strconv.Atoi(id)
	idFromToken := middlewares.ExtractToken(c)
	cartReq := _requestCart.Cart{}
	err := c.Bind(&cartReq)

	if err != nil {
		return c.JSON(http.StatusBadRequest, _helper.FailedResponseHelper("failed to bind data, check your input"))
	}
	qty := cartReq.Qty
	row, errUpd := h.cartBusiness.UpdateData(qty, idCart, idFromToken)
	if errUpd != nil {
		return c.JSON(http.StatusInternalServerError, _helper.FailedResponseHelper("you dont have access"))
	}
	if row == 0 {
		return c.JSON(http.StatusBadRequest, _helper.FailedResponseHelper("failed to update data"))
	}
	return c.JSON(http.StatusOK, _helper.SuccessResponseHelper("success update cart"))
}

func (h *CartHandler) DeleteCart(c echo.Context) error {
	id := c.Param("id")
	idProd, _ := strconv.Atoi(id)
	idFromToken := middlewares.ExtractToken(c)
	row, errDel := h.cartBusiness.DeleteData(idProd, idFromToken)
	if errDel != nil {
		return c.JSON(http.StatusInternalServerError, _helper.FailedResponseHelper("you dont have access"))
	}
	if row != 1 {
		return c.JSON(http.StatusBadRequest, _helper.FailedResponseHelper("failed to delete data user"))
	}
	return c.JSON(http.StatusOK, _helper.SuccessResponseHelper("success delete cart"))
}
