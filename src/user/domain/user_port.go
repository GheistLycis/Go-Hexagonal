package user

import (
	"time"

	"github.com/google/uuid"
)

type UserPort interface {
	Validate() error
	Enable() error
	Disable() error
	GetID() uuid.UUID
	GetStatus() Status
	GetBirthDate() time.Time
	GetName() string
	GetEmail() string
	GetGender() Gender
}
