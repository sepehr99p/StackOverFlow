package models

type Question struct {
	UserId      int64    `json:"user_id"`
	TagId       int      `json:"tag_idz"`
	Description string   `json:"description"`
	Votes       int      `json:"votes"`
	Answers     []Answer `json:"answers"`
}

var mockQuestion = Question{
	UserId: 123, Description: "some description", Votes: 33, Answers: []Answer{mockAnswer, mockAnswer}, TagId: 33,
}
