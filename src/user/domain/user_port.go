package user

import "time"

type UserPort interface {
	Validate() (isValid bool, err error)
	Enable() error
	Disable() error
	GetID() string
	GetStatus() Status
	GetBirthDate() time.Time
	GetName() string
	GetEmail() string
	GetGender() Gender
}
