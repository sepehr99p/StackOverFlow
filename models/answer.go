package models

type Answer struct {
	AnswerId        int    `json:"answer_id" gorm:"primary_key"`
	QuestionId      int64  `json:"question_id" gorm:"index"`                     // Foreign key to Question
	UserId          int    `json:"user_id" gorm:"index"`                         // Foreign key to User
	IsCorrectAnswer bool   `json:"is_correct_answer" gorm:"default:false;index"` // Index for filtering correct answers
	DateCreated     int64  `json:"date" gorm:"autoCreateTime;index"`             // Index for sorting by date
	Description     string `json:"description"`
	Votes           int    `json:"votes" gorm:"index"` // Index for sorting by votes
	//Comments    []*Comment `json:"comments" gorm:"foreignKey:ParentId;constraint:OnDelete:CASCADE;"`
}
