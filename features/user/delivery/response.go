package delivery

type UserResponse struct {
	Name     string `json:"name" form:"name"`
	PPicture string `json:"ppicture" form:"ppicture"`
	Username string `json:"username" form:"username"`
}
