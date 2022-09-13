package data

import (
	"errors"
	"project/e-commerce/features/product"

	"gorm.io/gorm"
)

type ProductRepo struct {
	DB *gorm.DB
}

func New(db *gorm.DB) product.DataInterface {
	return &ProductRepo{
		DB: db,
	}
}

func (repo *ProductRepo) GetProduct(limit, offset int) (data []product.Core, err error) {
	dataProduct := []Product{}
	result := repo.DB.Preload("User").Limit(limit).Offset(offset).Find(&dataProduct)
	if result.Error != nil {
		return []product.Core{}, result.Error
	}
	return ToCoreList(dataProduct), nil
}

func (repo *ProductRepo) InsertProduct(add product.Core) (row int, err error) {
	product := FromCore(add)
	result := repo.DB.Create(&product)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected != 1 {
		return 0, errors.New("failed to insert data")
	}
	return int(result.RowsAffected), nil
}

func (repo *ProductRepo) SelectProductById(id int) (data product.Core, err error) {
	dataProduct := Product{}
	result := repo.DB.Preload("User").Find(&dataProduct, id)
	if result.Error != nil {
		return product.Core{}, result.Error
	}
	return dataProduct.ToCore(), nil

}

func (repo *ProductRepo) UpdateProduct(data map[string]interface{}, id, token int) (row int, err error) {
	dataProduct := Product{}
	res := repo.DB.First(&dataProduct, id)
	if res.Error != nil {
		return -1, res.Error
	}
	if dataProduct.UserID != uint(token) {
		return -1, errors.New("you can't access")
	}
	result := repo.DB.Model(&Product{}).Where("id = ?", id).Updates(data)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected != 1 {
		return 0, errors.New("failed to update data")
	}
	return int(result.RowsAffected), nil
}

func (repo *ProductRepo) GetMyProduct(id int) (data []product.Core, err error) {
	dataProduct := []Product{}
	result := repo.DB.Preload("User").Find(&dataProduct).Where("userid = ?", id)
	if result.Error != nil {
		return []product.Core{}, result.Error
	}
	return ToCoreList(dataProduct), nil
}

func (repo *ProductRepo) DeleteMyProduct(id, token int) (row int, err error) {
	data := Product{}
	CheckId := repo.DB.First(&data, id)
	if CheckId.Error != nil {
		return 0, CheckId.Error
	}
	if id != int(data.UserID) {
		return -1, errors.New("you can't access")
	}
	result := repo.DB.Delete(&Product{}, id)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected != 1 {
		return 0, errors.New("failed to delete data")
	}
	return int(result.RowsAffected), nil
}
