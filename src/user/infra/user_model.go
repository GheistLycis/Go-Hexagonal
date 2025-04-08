package user

import (
	domain "Go-Hexagonal/src/user/domain"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	ID        uuid.UUID     `gorm:"type:uuid;default:gen_random_uuid()"`
	Status    domain.Status `gorm:"type:user_status"`
	Name      string
	Email     string        `gorm:"unique"`
	Gender    domain.Gender `gorm:"type:user_gender"`
	BirthDate time.Time
	CreatedBy string
	UpdatedBy *string
}

func (UserModel) TableName() string {
	return "user"
}
