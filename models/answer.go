package models

type Answer struct {
	AnswerId    int    `json:"answer_id" gorm:"primary_key"`
	QuestionId  int64  `json:"question_id" gorm:"index"` // Foreign key to Question
	UserId      int64  `json:"user_id"`
	DateCreated int64  `json:"date" gorm:"autoCreateTime"`
	Description string `json:"description"`
	Votes       int    `json:"votes"`
}

var mockAnswer = Answer{UserId: 234, Description: "some answer description", Votes: 44, DateCreated: 33344}
