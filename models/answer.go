package models

import "gorm.io/gorm"

type Answer struct {
	gorm.Model
	AnswerId    int    `json:"answer_id" gorm:"primary_key"`
	QuestionId  int    `json:"question_id" gorm:"foreignkey:FkOtherId;references:QuestionId"`
	UserId      int64  `json:"user_id"`
	DateCreated int64  `json:"date" gorm:"autoCreateTime"`
	Description string `json:"description"`
	Votes       int    `json:"votes"`
}

type Comment struct {
}

var mockAnswer = Answer{UserId: 234, Description: "some answer description", Votes: 44, DateCreated: 33344}
