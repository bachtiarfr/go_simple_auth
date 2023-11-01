package entity

type User struct {
	ID          int    `json:"user_id"`
	Fullname    string `json:"fullname"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Age         int    `json:"age"`
	PhoneNumber string `json:"phone_number"`
}
