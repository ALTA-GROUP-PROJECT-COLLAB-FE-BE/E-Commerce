package presentation

import (
	"net/http"
	"project/e-commerce/features/product"
	_requestProduct "project/e-commerce/features/product/delivery/request"
	_responseProduct "project/e-commerce/features/product/delivery/response"
	"project/e-commerce/middlewares"
	"strconv"

	_helper "project/e-commerce/utils/helper"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productBusiness product.Business
}

func NewProductHandler(e *echo.Echo, productBusiness product.Business) {
	handler := &ProductHandler{
		productBusiness: productBusiness,
	}
	e.POST("/product", handler.PostProduct, middlewares.JWTMiddleware())
	e.GET("/product", handler.GetAll)
	e.PUT("/product/:id", handler.PutProduct, middlewares.JWTMiddleware())
	e.GET("/product/:id", handler.GetById, middlewares.JWTMiddleware())
	e.GET("/product/me", handler.GetByMe, middlewares.JWTMiddleware()) //endpoint untuk melihat profile
	e.DELETE("/product/:id", handler.DeleteById, middlewares.JWTMiddleware())

}

func (h *ProductHandler) GetAll(c echo.Context) error {
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")
	limitint, _ := strconv.Atoi(limit)
	offsetint, _ := strconv.Atoi(offset)
	result, err := h.productBusiness.GetAllData(limitint, offsetint)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.FailedResponseHelper("failed to get all products"))
	}
	return c.JSON(http.StatusOK, _helper.SuccessDataResponseHelper("success", _responseProduct.FromCoreList(result)))
}

func (h *ProductHandler) PostProduct(c echo.Context) error {
	prodReq := _requestProduct.Product{}
	id := middlewares.ExtractToken(c)
	err := c.Bind(&prodReq)
	prodReq.UserID = id
	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.FailedResponseHelper("failed to bind data, check your input"))
	}

	dataProduct := _requestProduct.ToCore(prodReq)
	row, errCreate := h.productBusiness.CreateData(dataProduct)
	if row == -1 {
		return c.JSON(http.StatusBadRequest, _helper.FailedResponseHelper("please make sure all fields are filled in correctly"))
	}

	if errCreate != nil {
		return c.JSON(http.StatusInternalServerError, _helper.FailedResponseHelper("failed to create product, check your input"))
	}

	return c.JSON(http.StatusOK, _helper.SuccessResponseHelper("success"))
}

func (h *ProductHandler) GetById(c echo.Context) error {
	id := c.Param("id")
	idUser, _ := strconv.Atoi(id)
	result, errGet := h.productBusiness.GetProductById(idUser)
	if errGet != nil {
		return c.JSON(http.StatusInternalServerError, _helper.FailedResponseHelper("failed to get data product"))
	}
	return c.JSON(http.StatusOK, _helper.SuccessDataResponseHelper("success", _responseProduct.FromCore(result)))
}

func (h *ProductHandler) PutProduct(c echo.Context) error {
	id := c.Param("id")
	idProd, _ := strconv.Atoi(id)
	idFromToken := middlewares.ExtractToken(c)
	prodReq := _requestProduct.Product{}
	err := c.Bind(&prodReq)

	if err != nil {
		return c.JSON(http.StatusBadRequest, _helper.FailedResponseHelper("failed to bind data, check your input"))
	}
	dataProduct := _requestProduct.ToCore(prodReq)
	row, errUpd := h.productBusiness.UpdateData(dataProduct, idProd, idFromToken)
	if errUpd != nil {
		return c.JSON(http.StatusInternalServerError, _helper.FailedResponseHelper("you dont have access"))
	}
	if row == 0 {
		return c.JSON(http.StatusBadRequest, _helper.FailedResponseHelper("failed to update data"))
	}
	return c.JSON(http.StatusOK, _helper.SuccessResponseHelper("success"))
}

func (h *ProductHandler) GetByMe(c echo.Context) error {
	idFromToken := middlewares.ExtractToken(c)
	result, err := h.productBusiness.GetDataByMe(idFromToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.FailedResponseHelper("failed to get your product data"))
	}
	return c.JSON(http.StatusOK, _helper.SuccessDataResponseHelper("success", _responseProduct.FromCoreList(result)))
}

func (h *ProductHandler) DeleteById(c echo.Context) error {
	id := c.Param("id")
	idProd, _ := strconv.Atoi(id)
	idFromToken := middlewares.ExtractToken(c)
	row, errDel := h.productBusiness.DeleteDataById(idProd, idFromToken)
	if errDel != nil {
		return c.JSON(http.StatusInternalServerError, _helper.FailedResponseHelper("you dont have access"))
	}
	if row != 1 {
		return c.JSON(http.StatusBadRequest, _helper.FailedResponseHelper("failed to delete data user"))
	}

	return c.JSON(http.StatusOK, _helper.SuccessResponseHelper("success"))
}
