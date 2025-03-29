package models

type Tag struct {
	TagId int    `json:"tag_id" gorm:"primary_key"`
	Name  string `json:"name"`
}
