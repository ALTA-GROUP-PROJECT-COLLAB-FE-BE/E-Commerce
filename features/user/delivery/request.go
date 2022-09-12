package delivery

import "project/e-commerce/features/user"

type UserRequest struct {
	Name     string `json:"name" form:"name"`
	PPicture string `json:"name" form:"name"`
	Username string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func toCore(data UserRequest) user.Core {
	return user.Core{
		Name:     data.Name,
		PPicture: data.PPicture,
		Username: data.Username,
		Email:    data.Email,
		Password: data.Password,
	}
}
