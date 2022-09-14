package data

import (
	"errors"
	"project/e-commerce/features/user"

	"gorm.io/gorm"
)

type userData struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.DataInterface {
	return &userData{
		db: db,
	}
}

func (repo *userData) InsertData(data user.Core) (int, error) {
	dataModel := fromCore(data)
	tx := repo.db.Create(&dataModel)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil
}

func (repo *userData) SelectData(id int) (user.Core, error) {
	var user User
	tx := repo.db.First(&user, id)
	if tx.Error != nil {
		return user.toCore(), tx.Error
	}

	return user.toCore(), nil
}

func (repo *userData) UpdateData(data user.Core, id int) (row int, err error) {
	tx := repo.db.Model(&User{}).Where("id = ?", id).Updates(fromCore(data))
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("failed to update data")
	}

	return int(tx.RowsAffected), nil
}

func (repo *userData) DeleteData(id int) (row int, err error) {
	tx := repo.db.Delete(&User{}, id)
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("failed delete akun")
	}
	return int(tx.RowsAffected), nil
}
