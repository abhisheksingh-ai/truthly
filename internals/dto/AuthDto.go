package dto

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LogInRes struct {
	UserId   string `json:"userId,omitempty"`
	UserName string `json:"userName,omitempty"`
	Token    string `json:"token,omitempty"`
}
