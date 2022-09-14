package delivery

import (
	"net/http"
	"project/e-commerce/features/cart"
	requestCart "project/e-commerce/features/cart/delivery/request"
	"project/e-commerce/middlewares"
	"project/e-commerce/utils/helper"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CartHandler struct {
	cartBusiness cart.Business
}

func NewCartHandler(e *echo.Echo, cartBusiness cart.Business) {
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
	IntLimit, _ := strconv.Atoi(limit)
	IntOffset, _ := strconv.Atoi(offset)
	idFromToken := middlewares.ExtractToken(c)
	result, err := h.cartBusiness.GetAllData(IntLimit, IntOffset, idFromToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("failed get all data"))
	}
	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("success get data", result))
}

func (h *CartHandler) PostCart(c echo.Context) error {
	idFromToken := middlewares.ExtractToken(c)
	cartReq := requestCart.Cart{}
	err := c.Bind(&cartReq)
	if err != nil {
		return c.JSON(200, helper.FailedResponseHelper("failed to get data check your input"))
	}
	dataCart := cart.Core{}
	dataCart.Product.ID = cartReq.IdProduct
	dataCart.Qty = cartReq.Qty
	dataCart.Status = "On Process"
	dataCart.UserID = idFromToken
	row, errCreate := h.cartBusiness.CreateData(dataCart)
	if row == -1 {
		return c.JSON(400, helper.FailedResponseHelper("column cannot be empty"))
	}
	if errCreate != nil {
		return c.JSON(406, helper.FailedResponseHelper("failed add to cart"))
	}
	return c.JSON(200, helper.SuccessResponseHelper("success"))
}

func (h *CartHandler) UpdateCart(c echo.Context) error {
	id := c.Param("id")
	cartId, _ := strconv.Atoi(id)
	idToken := middlewares.ExtractToken(c)
	reqC := requestCart.Cart{}
	err := c.Bind(&reqC)

	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("failed to get data check your input"))
	}
	qty := reqC.Qty
	row, errUp := h.cartBusiness.UpdateData(qty, cartId, idToken)
	if errUp != nil {
		return c.JSON(401, helper.FailedResponseHelper("you can't access"))
	}
	if row == 0 {
		return c.JSON(400, helper.FailedResponseHelper("failed to update data"))
	}
	return c.JSON(200, helper.SuccessResponseHelper("succes to update data"))
}

func (h *CartHandler) DeleteCart(c echo.Context) error {
	id := c.Param("id")
	prodId, _ := strconv.Atoi(id)
	tokenId := middlewares.ExtractToken(c)
	row, delEr := h.cartBusiness.DeleteData(prodId, tokenId)
	if delEr != nil {
		return c.JSON(401, helper.FailedResponseHelper("you can't access"))
	}
	if row != 1 {
		return c.JSON(400, helper.FailedResponseHelper("failed to delete data"))
	}
	return c.JSON(200, helper.SuccessResponseHelper("success to delete data"))
}
