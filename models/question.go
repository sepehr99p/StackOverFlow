package models

type Question struct {
	QuestionId  int64  `json:"question_id" gorm:"primary_key"`
	UserId      int    `json:"user_id" gorm:"index"` // Foreign key to User
	Description string `json:"description"`
	Votes       int    `json:"votes" gorm:"index"`               // Index for sorting by votes
	DateCreated int64  `json:"date" gorm:"autoCreateTime;index"` // Index for sorting by date
	Tags        []*Tag `json:"tags" gorm:"many2many:tag_questions"`
	//Answers     []*Answer  `json:"answers" gorm:"foreignKey:QuestionId"` // One-to-Many
	//Comments    []*Comment `json:"comments" gorm:"foreignKey:ParentId;constraint:OnDelete:CASCADE;"`
}
