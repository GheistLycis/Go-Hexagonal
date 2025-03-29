package user

type UserRepoPort interface {
	Create(user UserPort, createdBy string) (*User, error)
	Get(filters GetUserRepoFiltersDTO) (*User, error)
	List(filters ListUsersRepoFiltersDTO) ([]*User, error)
	Update(user UserPort, updatedBy string) (*User, error)
}

type GetUserRepoFiltersDTO struct {
	ID    *string
	Email *string
}

type ListUsersRepoFiltersDTO struct {
	Name   *string
	Status *Status
	Gender *Gender
}
