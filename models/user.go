package models

type User struct {
	UserId     int    `json:"user_id" gorm:"primary_key"`
	UserName   string `json:"user_name"`
	Reputation int    `json:"reputation"`
	IsAdmin    bool   `json:"is_admin" gorm:"default:false"`
}
