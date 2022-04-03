package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"primaryKey"`
	Name     string    `gorm:"column:name"`
	Email    *string   `gorm:"column:email"`
	Phone    string    `gorm:"column:phone"`
	Password string    `gorm:"column:password" json:"-"`
	Address  string    `gorm:"column:address"`
	Birthday time.Time `gorm:"column:birthday"`
	Status   bool      `gorm:"column:status"`
	Token    string    `gorm:"column:token" json:"-"`
	Products []Product `gorm:"foreignKey:UserId" json:"products"`
}
