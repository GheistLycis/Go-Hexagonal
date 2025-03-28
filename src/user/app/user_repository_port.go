package user

import domain "Go-Hexagonal/src/user/domain"

type UserRepoPort interface {
	Create(user domain.UserPort, createdBy string) (*domain.User, error)
	Get(filters GetUserRepoFiltersDTO) (*domain.User, error)
	List(filters ListUsersRepoFiltersDTO) ([]*domain.User, error)
	Update(user domain.UserPort, updatedBy string) (*domain.User, error)
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
