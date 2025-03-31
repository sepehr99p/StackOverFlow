package models

type Comment struct {
	CommentId   int    `json:"comment_id" gorm:"primary_key"`
	ParentId    int64  `json:"parent_id"`
	ParentType  string `json:"parent_type"` // "question" or "answer"
	UserId      int64  `json:"user_id"`
	DateCreated int64  `json:"date" gorm:"autoCreateTime"`
	Description string `json:"description"`
	Votes       int    `json:"votes"`
}
