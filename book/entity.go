package book

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Description string
	Price       float64
	Rating      uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
