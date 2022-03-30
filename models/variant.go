package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Variant struct {
	gorm.Model
	ID   uuid.UUID `gorm:"type:uuid"`
	Name string    `json:"name"`
}
