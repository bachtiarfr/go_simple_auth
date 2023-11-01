package entity

type LoginResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

type UserResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    []User `json:"data"`
}
