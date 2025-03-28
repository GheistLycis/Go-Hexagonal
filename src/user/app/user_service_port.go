package user

import (
	domain "Go-Hexagonal/src/user/domain"
	"time"
)

type UserServicePort interface {
	Create(payload CreateUserServiceDTO, createdBy string) (*domain.User, error)
	Disable(id string, updatedBy string) (*domain.User, error)
	Enable(id string, updatedBy string) (*domain.User, error)
	Get(filters GetUserServiceFiltersDTO) (*domain.User, error)
	List(filters ListUsersServiceFiltersDTO) ([]*domain.User, error)
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
