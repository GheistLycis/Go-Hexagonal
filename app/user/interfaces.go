package app

// USER

type UserI interface {
	Validate() (bool, error)
	Enable() error
	Disable() error
	GetID() string
	GetStatus() Status
	GetName() string
	GetEmail() string
	GetGender() Gender
}

type Status string

const (
	ENABLED     Status = "ATIVO"
	IN_ANALYSIS Status = "EM ANÁLISE"
	DISABLED    Status = "INATIVO"
)

type Gender string

const (
	MALE   Gender = "MASCULINO"
	FEMALE Gender = "FEMININO"
	OTHER  Gender = "OUTRO"
)

// SERVICE

type UserServiceI interface {
	Create(payload CreateUserServicePayload) (UserI, error)
	Disable(id string) (UserI, error)
	Enable(id string) (UserI, error)
	Get(filters GetUserServiceFilters) (UserI, error)
	List(filters ListUsersServiceFilters) ([]UserI, error)
}

type CreateUserServicePayload struct {
	Name   string
	Email  string
	Gender Gender
}

type GetUserServiceFilters struct {
	ID    *string
	Email *string
}

type ListUsersServiceFilters struct {
	Name   *string
	Status *Status
	Gender *Gender
}

// REPOSITORY

type UserRepoI interface {
	Create(user UserI) (UserI, error)
	Get(filters GetUserRepoFilters) (UserI, error)
	List(filters ListUsersRepoFilters) ([]UserI, error)
	Update(user UserI) (UserI, error)
}

type GetUserRepoFilters struct {
	ID    *string
	Email *string
}

type ListUsersRepoFilters struct {
	Name   *string
	Status *Status
	Gender *Gender
}
