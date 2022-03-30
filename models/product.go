package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid"`
	Title       string
	Description string
	Category    string
	Price       int64
}
