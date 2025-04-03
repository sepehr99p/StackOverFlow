package models

type Answer struct {
	AnswerId    int    `json:"answer_id" gorm:"primary_key"`
	QuestionId  int64  `json:"question_id" gorm:"index"` // Foreign key to Question
	UserId      int64  `json:"user_id"`
	DateCreated int64  `json:"date" gorm:"autoCreateTime"`
	Description string `json:"description"`
	Votes       int    `json:"votes"`
	//Comments    []*Comment `json:"comments" gorm:"foreignKey:ParentId;constraint:OnDelete:CASCADE;"`
}
