package user

import (
	"time"
)

type UserServicePort interface {
	Create(payload CreateUserServiceDTO, createdBy string) (*User, error)
	Disable(id string, updatedBy string) (*User, error)
	Enable(id string, updatedBy string) (*User, error)
	Get(filters GetUserServiceFiltersDTO) (*User, error)
	List(filters ListUsersServiceFiltersDTO) ([]*User, error)
}

type CreateUserServiceDTO struct {
	Name      string    `json:"name" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	Gender    Gender    `json:"gender" binding:"required"`
	BirthDate time.Time `json:"birthDate" binding:"required"`
}

type GetUserServiceFiltersDTO struct {
	ID    *string
	Email *string
}

type ListUsersServiceFiltersDTO struct {
	Name   *string
	Status *Status
	Gender *Gender
}
