package user

import domain "Go-Hexagonal/src/user/domain"

type UserRepoPort interface {
	Create(user domain.UserPort, createdBy string) (domain.UserPort, error)
	Get(filters GetUserRepoFiltersDTO) (domain.UserPort, error)
	List(filters ListUsersRepoFiltersDTO) ([]domain.UserPort, error)
	Update(user domain.UserPort, updatedBy string) (domain.UserPort, error)
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
