package models

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	QuestionId  int64     `json:"question_id" gorm:"primary_key"`
	UserId      int64     `json:"user_id"`
	TagId       int       `json:"tag_idz"`
	Description string    `json:"description"`
	Votes       int       `json:"votes"`
	DateCreated int64     `json:"date" gorm:"autoCreateTime"`
	Answers     []*Answer `json:"answers" gorm:"many2many:question_answers;foreignKey:QuestionId;joinForeignKey:QuestionId"`
}

var mockQuestion = Question{
	UserId: 123, Description: "some description", Votes: 33, Answers: []*Answer{&mockAnswer, &mockAnswer}, TagId: 33,
	QuestionId: 33,
}
