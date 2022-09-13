package data

import (
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

// var user models.User
// tx := config.DB.First(&user, id)
// if tx.Error != nil {
// 	return entities.UserCore{}, tx.Error
// }

// return models.ToEntities(user), nil
