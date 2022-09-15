package factory

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	userData "project/e-commerce/features/user/data"
	userDelivery "project/e-commerce/features/user/delivery"
	userUsecase "project/e-commerce/features/user/usecase"

	productData "project/e-commerce/features/product/data"
	productDelivery "project/e-commerce/features/product/delivery"
	productUseCase "project/e-commerce/features/product/usecase"

	cartData "project/e-commerce/features/cart/data"
	cartDelivery "project/e-commerce/features/cart/delivery"
	cartUsecase "project/e-commerce/features/cart/usecase"

	loginData "project/e-commerce/features/login/data"
	authDelivery "project/e-commerce/features/login/delivery"
	authUsecase "project/e-commerce/features/login/usecase"

	orderData "project/e-commerce/features/order/data"
	orderDelivery "project/e-commerce/features/order/delivery"
	orderUsecase "project/e-commerce/features/order/usecase"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	userDataFactory := userData.New(db)
	userUsecaseFactory := userUsecase.New(userDataFactory)
	userDelivery.New(e, userUsecaseFactory)

	productData := productData.NewProductRepository(db)
	productUsecase := productUseCase.NewProductBusiness(productData)
	productDelivery.NewProductHandler(e, productUsecase)

	cartDataFactory := cartData.NewCartRepo(db)
	cartUsecase := cartUsecase.NewCartUsecase(cartDataFactory)
	cartDelivery.NewCartHandler(e, cartUsecase)

	authDataFactory := loginData.New(db)
	authUsecaseFactory := authUsecase.New(authDataFactory)
	authDelivery.New(e, authUsecaseFactory)

	orderDataFactory := orderData.New(db)
	orderUsecaseFactory := orderUsecase.New(orderDataFactory)
	orderDelivery.New(e, orderUsecaseFactory)
}
