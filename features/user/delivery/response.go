package delivery

import "project/e-commerce/features/user"

type UserResponse struct {
	Name     string `json:"name" form:"name"`
	PPicture string `json:"ppicture" form:"ppicture"`
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
}

func FromCore(data user.Core) UserResponse {
	return UserResponse{
		Name:     data.Name,
		PPicture: data.PPicture,
		Username: data.Username,
		Email:    data.Email,
	}
}
