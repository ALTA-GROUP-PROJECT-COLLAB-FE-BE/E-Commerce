package data

import (
	"errors"
	"project/e-commerce/features/cart"

	"gorm.io/gorm"
)

type CartRepo struct {
	DB *gorm.DB
}

func NewCartRepo(db *gorm.DB) cart.Data {
	return &CartRepo{
		DB: db,
	}
}

func (repo *CartRepo) SelectData(limit, offset, idFromToken int) (data []cart.Core, err error) {
	dataC := []Cart{}
	result := repo.DB.Preload("Product").Where("user_id=? AND status = ?", idFromToken, "On Process").Find(&dataC)
	if result.Error != nil {
		return []cart.Core{}, result.Error
	}
	return ToCoreList(dataC), nil
}

func (repo *CartRepo) CheckCart(idProd, idFromToken int) (isExist bool, idCart, qty int, err error) {
	dataC := Cart{}
	resCheck := repo.DB.Model(&Cart{}).Where("product_id = ? AND user_id = ? AND status = ?", idProd, idFromToken, "On Process").First(&dataC)
	if resCheck.Error != nil {
		return false, 0, 0, resCheck.Error
	}
	return true, int(dataC.ID), int(dataC.Qty), nil
}

func (repo *CartRepo) InsertData(data cart.Core) (row int, err error) {
	cart := FromCore(data)
	result := repo.DB.Create(&cart)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected != 1 {
		return 0, errors.New("failed to create data")
	}
	return int(result.RowsAffected), nil
}

func (repo *CartRepo) UpdateDataDB(qty, idCart, idFromToken int) (row int, err error) {
	dataC := Cart{}
	checkId := repo.DB.First(&dataC, idCart)

	if checkId.Error != nil {
		return 0, checkId.Error
	}
	if dataC.UserID != idFromToken {
		return -1, errors.New("you dont have access")
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

func (repo *CartRepo) DeleteDataDB(idCart, idFromToken int) (row int, err error) {
	dataC := Cart{}
	check := repo.DB.First(&dataC, idCart)
	if check.Error != nil {
		return 0, check.Error
	}
	if idFromToken != dataC.UserID {
		return -1, errors.New("you dont have access")
	}
	result := repo.DB.Delete(&Cart{}, idCart)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected != 1 {
		return 0, errors.New("delete data failed")
	}
	return int(result.RowsAffected), nil
}
