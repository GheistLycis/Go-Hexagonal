package user

import (
	domain "Go-Hexagonal/src/user/domain"
	"time"
)

type UserServicePort interface {
	Create(payload CreateUserServiceDTO, createdBy string) (domain.UserPort, error)
	Disable(id string, updatedBy string) (domain.UserPort, error)
	Enable(id string, updatedBy string) (domain.UserPort, error)
	Get(filters GetUserServiceFiltersDTO) (domain.UserPort, error)
	List(filters ListUsersServiceFiltersDTO) ([]domain.UserPort, error)
}

type CreateUserServiceDTO struct {
	Name      string        `json:"name" binding:"required"`
	Email     string        `json:"email" binding:"required"`
	Gender    domain.Gender `json:"gender" binding:"required"`
	BirthDate time.Time     `json:"birthDate" binding:"required"`
}

type GetUserServiceFiltersDTO struct {
	ID    *string
	Email *string
}

type ListUsersServiceFiltersDTO struct {
	Name   *string
	Status *domain.Status
	Gender *domain.Gender
}
