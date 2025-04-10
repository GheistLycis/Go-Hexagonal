package user

import domain "Go-Hexagonal/src/user/domain"

type UserRepoPort interface {
	// Create creates a new user in the DB
	Create(user UserPort, createdBy string) (*domain.User, error)

	// Get returns the first user found in the DB matching the filters
	Get(filters GetUserRepoFiltersDTO) (*domain.User, error)

	// List returns all users found in the DB matching the filters
	List(filters ListUsersRepoFiltersDTO) ([]*domain.User, error)

	// Update updates given user in the DB, searching it by ID. It returns an error if user is not found
	Update(user UserPort, updatedBy string) (*domain.User, error)
}

type GetUserRepoFiltersDTO struct {
	ID    *string
	Email *string
}

type ListUsersRepoFiltersDTO struct {
	Name   *string
	Status *domain.Status
	Gender *domain.Gender
}
