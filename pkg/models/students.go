package models

import "time"

type Student struct {
	ID         string    `gorm:"primaryKey"`
	Created_at time.Time `gorm:"autoCreateTime"`
	Updated_at time.Time `gorm:"autoUpdateTime"`
}
