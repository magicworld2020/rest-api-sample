package model

type User struct {
	UserID   string `json:"user_id"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
}

type CreateUserResponse struct {
	Message string `json:"message"`
	User    User   `json:"user"`
}
