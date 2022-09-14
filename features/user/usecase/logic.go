package usecase

import (
	"errors"
	"project/e-commerce/features/user"

	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userData user.DataInterface
}

func New(data user.DataInterface) user.UsecaseInterface {
	return &userUsecase{
		userData: data,
	}
}

func (usecase *userUsecase) PostData(data user.Core) (int, error) {
	if data.Name == "" {
		return -1, errors.New("masukan nama ")
	}
	if data.Email == "" {
		return -1, errors.New("masukan email ")
	}
	if data.Password == "" {
		return -1, errors.New("masukann password")
	}
	hashPass, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return -1, err
	}
	data.Password = string(hashPass)
	row, err := usecase.userData.InsertData(data)
	if err != nil {
		return -1, err
	}
	// row, err := usecase.userData.InsertData(data)
	return row, nil
}

func (usecase *userUsecase) GetData(id int) (user.Core, error) {
	results, err := usecase.userData.SelectData(id)
	return results, err
}

func (usecase *userUsecase) PostDataId(data user.Core, id int) (int, error) {
	row, err := usecase.userData.UpdateData(data, id)
	return row, err
}

func (usecase *userUsecase) DeleteAkun(id int) (int, error) {
	row, err := usecase.userData.DeleteData(id)
	return row, err
}
