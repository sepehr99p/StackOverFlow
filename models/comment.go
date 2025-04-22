package models

type Comment struct {
	CommentId   int    `json:"comment_id" gorm:"primary_key"`
	ParentId    int64  `json:"parent_id" gorm:"index"`
	ParentType  string `json:"parent_type" gorm:"index"`
	UserId      int64  `json:"user_id" gorm:"index"`
	DateCreated int64  `json:"date" gorm:"autoCreateTime;index"`
	Description string `json:"description"`
	Votes       int    `json:"votes" gorm:"index"`
}
