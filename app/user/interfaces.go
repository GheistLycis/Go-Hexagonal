package app

import "time"

// USER

type UserI interface {
	Validate() (bool, error)
	Enable() error
	Disable() error
	GetID() string
	GetStatus() Status
	GetBirthDate() time.Time
	GetName() string
	GetEmail() string
	GetGender() Gender
}

/*
const (

	ENABLED     Status = "ATIVO"
	IN_ANALYSIS Status = "EM ANÁLISE"
	DISABLED    Status = "INATIVO"

)
*/
type Status string

const (
	ENABLED     Status = "ATIVO"
	IN_ANALYSIS Status = "EM ANÁLISE"
	DISABLED    Status = "INATIVO"
)

/*
const (

	MALE   Gender = "MASCULINO"
	FEMALE Gender = "FEMININO"
	OTHER  Gender = "OUTRO"

)
*/
type Gender string

const (
	MALE   Gender = "MASCULINO"
	FEMALE Gender = "FEMININO"
	OTHER  Gender = "OUTRO"
)

// SERVICE

type UserServiceI interface {
	Create(payload CreateUserServicePayload, createdBy string) (UserI, error)
	Disable(id string, updatedBy string) (UserI, error)
	Enable(id string, updatedBy string) (UserI, error)
	Get(filters GetUserServiceFilters) (UserI, error)
	List(filters ListUsersServiceFilters) ([]UserI, error)
}

type CreateUserServicePayload struct {
	Name      string    `json:"name" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	Gender    Gender    `json:"gender" binding:"required"`
	BirthDate time.Time `json:"birthDate" binding:"required"`
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
	Create(user UserI, createdBy string) (UserI, error)
	Get(filters GetUserRepoFilters) (UserI, error)
	List(filters ListUsersRepoFilters) ([]UserI, error)
	Update(user UserI, updatedBy string) (UserI, error)
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
