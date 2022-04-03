package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID          uuid.UUID `gorm:"column:id;primaryKey"`
	Title       string    `gorm:"column:title"`
	Description string    `gorm:"column:description"`
	Category    string    `gorm:"column:category"`
	Price       int64     `gorm:"column:price"`
	UserId      uuid.UUID `gorm:"column:user_id" json:"-"`
	// User        User `gorm:"foreignKey:CreatedBy"`
}
