package models

type Question struct {
	QuestionId  int64      `json:"question_id" gorm:"primary_key"`
	UserId      int        `json:"user_id" gorm:"index"` // Foreign key to User
	TagId       int        `json:"tag_idz"`
	Description string     `json:"description"`
	Votes       int        `json:"votes"`
	DateCreated int64      `json:"date" gorm:"autoCreateTime"`
	Tags        []*Tag     `json:"tags" gorm:"many2many:tag_questions"`
	Answers     []*Answer  `json:"answers" gorm:"foreignKey:QuestionId"` // One-to-Many
	Comments    []*Comment `json:"comments" gorm:"foreignKey:ParentId;constraint:OnDelete:CASCADE;"`
}

var mockQuestion = Question{
	UserId: 123, Description: "some description", Votes: 33, Answers: []*Answer{&mockAnswer, &mockAnswer}, TagId: 33,
	QuestionId: 33,
}
