package models

type Tag struct {
	TagId     int         `json:"tag_id" gorm:"primary_key"`
	Name      string      `json:"name" gorm:"uniqueIndex"`
	Questions []*Question `json:"questions" gorm:"many2many:tag_questions"`
}
