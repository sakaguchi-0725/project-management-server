package domain

import "time"

type Task struct {
	ID        uint `gorm:"primaryKey"`
	Title     string
	Desc      string `gorm:"column:description"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
