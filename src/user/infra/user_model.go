package user

import (
	domain "Go-Hexagonal/src/user/domain"
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	ID        string        `gorm:"primaryKey"`
	Status    domain.Status `gorm:"type:user_status"`
	Name      string
	Email     string        `gorm:"unique"`
	Gender    domain.Gender `gorm:"type:user_gender"`
	BirthDate time.Time
	CreatedBy string
	CreatedAt time.Time
	UpdatedBy *string
	UpdatedAt *time.Time
}
