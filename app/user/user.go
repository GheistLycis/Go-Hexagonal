package app

import (
	"errors"
	"fmt"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}


const (
	ENABLED  = "ATIVO"
	IN_ANALYSIS = "EM ANÁLISE"
	DISABLED = "INATIVO"
)

type User struct {
	ID     string  `valid:"uuidv4"`
	Status string  
	Name   string  
	Email  string  `valid:"email"`
}

func NewUser(name string, email string) *User {
	return &User{
		ID:     uuid.NewV4().String(),
		Status: IN_ANALYSIS,
		Name: name,
		Email: email,
	}
}

type UserI interface {
	Validate() (bool, error)
	Enable() error
	Disable() error
}

func (p *User) Validate() (bool, error) {
	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return false, err
	}

	if p.Status != ENABLED && p.Status != DISABLED && p.Status != IN_ANALYSIS {
		return false, errors.New(fmt.Sprintln("Status inválido: %s", p.Status))
	}

	return true, nil
}

func (p *User) Enable() error {
	p.Status = ENABLED

	_, err := p.Validate()

	return err
}

func (p *User) Disable() error {
	p.Status = DISABLED

	_, err := p.Validate()

	return err
}