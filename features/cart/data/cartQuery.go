package data

import (
	"errors"
	_cart "project/e-commerce/features/cart"

	"gorm.io/gorm"
)

type mysqlCartRepository struct {
	DB *gorm.DB
}

func NewCartRepository(db *gorm.DB) _cart.Data {
	return &mysqlCartRepository{
		DB: db,
	}
}

func (repo *mysqlCartRepository) SelectData(limit, offset, idFromToken int) (data []_cart.Core, err error) {
	dataCart := []Cart{}
	result := repo.DB.Preload("Product").Where("user_id=? AND status = ?", idFromToken, "on process").Find(&dataCart)
	if result.Error != nil {
		return []_cart.Core{}, result.Error
	}
	return toCoreList(dataCart), nil
}

func (repo *mysqlCartRepository) CheckCart(idProd, idFromToken int) (isExist bool, idCart, qty int, err error) {
	dataCart := Cart{}
	resultCheck := repo.DB.Model(&Cart{}).Where("product_id = ? AND user_id = ? AND status = ?", idProd, idFromToken, "on process").First(&dataCart)
	if resultCheck.Error != nil {
		return false, 0, 0, resultCheck.Error
	}
	return true, int(dataCart.ID), int(dataCart.Qty), nil
}

func (repo *mysqlCartRepository) InsertData(data _cart.Core) (row int, err error) {
	cart := FromCore(data)
	result := repo.DB.Create(&cart)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected != 1 {
		return 0, errors.New("failed to create data cart")
	}
	return int(result.RowsAffected), nil
}

func (repo *mysqlCartRepository) UpdateDataDB(qty, idCart, idFromToken int) (row int, err error) {
	dataCart := Cart{}
	idCheck := repo.DB.First(&dataCart, idCart)

	if idCheck.Error != nil {
		return 0, idCheck.Error
	}
	if dataCart.UserID != idFromToken {
		return -1, errors.New("you don't have access")
	}
	result := repo.DB.Model(&Cart{}).Where("id = ?", idCart).Update("qty", qty)

	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected != 1 {
		return 0, errors.New("failed to update data")
	}
	return int(result.RowsAffected), nil
}

func (repo *mysqlCartRepository) DeleteDataDB(idCart, idFromToken int) (row int, err error) {
	dataCart := Cart{}
	idCheck := repo.DB.First(&dataCart, idCart)
	if idCheck.Error != nil {
		return 0, idCheck.Error
	}
	if idFromToken != dataCart.UserID {
		return -1, errors.New("you don't have access")
	}
	result := repo.DB.Delete(&Cart{}, idCart)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected != 1 {
		return 0, errors.New("failed to delete data")
	}
	return int(result.RowsAffected), nil
}
