package models

type Answer struct {
	AnswerId    int    `json:"answer_id" gorm:"primary_key"`
	QuestionId  int    `json:"question_id" gorm:"foreignkey:FkOtherId;references:QuestionId"`
	UserId      int64  `json:"user_id"`
	Description string `json:"description"`
	Votes       int    `json:"votes"`
}

var mockAnswer = Answer{UserId: 234, Description: "some answer description", Votes: 44}
