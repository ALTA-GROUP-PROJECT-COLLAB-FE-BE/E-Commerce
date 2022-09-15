package data

import (
	"project/e-commerce/features/order"

	"gorm.io/gorm"
)

type orderData struct {
	db *gorm.DB
}

func new(db *gorm.DB) order.DataInterface {
	return &orderData{
		db: db,
	}
}

func (repo *orderData) InsertData(data order.Core) (int, error) {
	dataModel := fromCore(data)
	tx := repo.db.Create(&dataModel)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil
}
