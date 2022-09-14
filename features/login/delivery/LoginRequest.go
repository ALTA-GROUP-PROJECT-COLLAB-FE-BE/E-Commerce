package delivery

import "project/e-commerce/features/login"

type AuthRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func ToCore(data AuthRequest) login.Core {
	return login.Core{
		Email:    data.Email,
		Password: data.Password,
	}
}
