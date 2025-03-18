package models

type Answer struct {
	UserId      int64  `json:"user_id"`
	Description string `json:"description"`
	Votes       int    `json:"votes"`
}

var mockAnswer = Answer{UserId: 234, Description: "some answer description", Votes: 44}
