package usecase

import (
	"errors"
	"project/e-commerce/features/user"
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
		return -1, errors.New("Masukan Nama !")
	}
	if data.Email == "" {
		return -1, errors.New("Masukan Email !")
	}
	if data.Password == "" {
		return -1, errors.New("Masukann Password")
	}
	row, err := usecase.userData.InsertData(data)
	return row, err
}
