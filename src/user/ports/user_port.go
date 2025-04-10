package user

import (
	domain "Go-Hexagonal/src/user/domain"
	"time"

	"github.com/google/uuid"
)

type UserPort interface {
	// Validate validates the user for business rules
	Validate() error

	// Enable sets user Status to ENABLED and calls .Validate(). It returns an error if user is already enabled or Validate returns any errors.
	Enable() error

	// Disable sets user Status to DISABLED and calls .Validate(). It returns an error if user is already disabled or Validate returns any errors.
	Disable() error

	GetID() uuid.UUID
	GetStatus() domain.Status
	GetBirthDate() time.Time
	GetName() string
	GetEmail() string
	GetGender() domain.Gender
}
