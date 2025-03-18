package user

type UserRepoPort interface {
	Create(user UserPort, createdBy string) (UserPort, error)
	Get(filters GetUserRepoFiltersDTO) (UserPort, error)
	List(filters ListUsersRepoFiltersDTO) ([]UserPort, error)
	Update(user UserPort, updatedBy string) (UserPort, error)
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
