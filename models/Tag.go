package models

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	TagId     int         `json:"tag_id" gorm:"primary_key"`
	Name      string      `json:"name"`
	Questions []*Question `json:"questions" gorm:"many2many:tag_questions"`
}
