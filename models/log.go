package models

import "time"

type Log struct {
	ID               uint `gorm:"primaryKey"`
	UserID           uint
	Action           string
	EntityType       string
	EntityID         uint
	ReputationChange int
	Description      string
	CreatedAt        time.Time
}
