package app

import "time"

type UserPort interface {
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
