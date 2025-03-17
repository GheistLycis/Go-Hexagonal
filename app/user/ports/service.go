package app

import (
	"time"
)

type UserServicePort interface {
	Create(payload CreateUserServiceDTO, createdBy string) (UserPort, error)
	Disable(id string, updatedBy string) (UserPort, error)
	Enable(id string, updatedBy string) (UserPort, error)
	Get(filters GetUserServiceFiltersDTO) (UserPort, error)
	List(filters ListUsersServiceFiltersDTO) ([]UserPort, error)
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
