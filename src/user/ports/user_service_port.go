package user

import (
	domain "Go-Hexagonal/src/user/domain"
	"time"
)

type UserServicePort interface {
	// Create creates a new user
	Create(payload CreateUserServiceDTO, createdBy string) (*domain.User, error)

	// Disable sets user's Status to DISABLED. It returns an error if user with given ID doesn't exist or any business rules are not met
	Disable(id string, updatedBy string) (*domain.User, error)

	// Update sets user's Status to ENABLED. It returns an error if user with given ID doesn't exist or any business rules are not met
	Enable(id string, updatedBy string) (*domain.User, error)

	// Get returns the first user matching the filters
	Get(filters GetUserServiceFiltersDTO) (*domain.User, error)

	// List returns all users matching the filters
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
