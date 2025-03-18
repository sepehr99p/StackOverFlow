package models

type User struct {
	UserId     int    `json:"user_id"`
	UserName   string `json:"user_name"`
	Reputation int    `json:"reputation"`
}

var mockUser = User{
	UserId: 11, UserName: "sepehr", Reputation: 371,
}
