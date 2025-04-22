package models

type User struct {
	UserId     int    `json:"user_id" gorm:"primary_key"`
	UserName   string `json:"user_name" gorm:"uniqueIndex"`
	Reputation int    `json:"reputation" gorm:"index"`
	IsAdmin    bool   `json:"is_admin" gorm:"default:false;index"`
	Password   string `json:"password"`
}
