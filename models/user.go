package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserId     int    `json:"user_id" gorm:"primary_key"`
	UserName   string `json:"user_name"`
	Reputation int    `json:"reputation"`
	//Questions  []*Question `json:"user_questions" gorm:"many2many:user_questions;foreignKey:Refer;joinForeignKey:QuestionId"`
}

var mockUser = User{
	UserId: 11, UserName: "sepehr", Reputation: 371,
}
