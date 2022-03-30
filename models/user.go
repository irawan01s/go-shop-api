package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid"`
	Name     string
	Email    *string
	Phone    string
	Password string
	Address  string
	Birthday time.Time
	Token    string
}
