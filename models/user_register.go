package models

type UserRegister struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
