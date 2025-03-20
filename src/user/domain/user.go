package user

import (
	"errors"
	"fmt"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type User struct {
	ID        string    `valid:"uuidv4"`
	Status    Status    `valid:"-"`
	Name      string    `valid:"-"`
	Email     string    `valid:"email"`
	Gender    Gender    `valid:"-"`
	BirthDate time.Time `valid:"-"`
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

func NewUser(name string, email string, gender Gender, BirthDate time.Time) (UserPort, error) {
	user := &User{
		ID:        uuid.NewV4().String(),
		Status:    IN_ANALYSIS,
		Name:      name,
		Email:     email,
		Gender:    gender,
		BirthDate: BirthDate,
	}

	if _, err := user.Validate(); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) Validate() (bool, error) {
	_, err := govalidator.ValidateStruct(u)
	if err != nil {
		return false, err
	}

	if u.Name == "" {
		return false, errors.New("o nome não pode ser vazio")
	}

	if u.Status != ENABLED && u.Status != DISABLED && u.Status != IN_ANALYSIS {
		return false, fmt.Errorf("status inválido: %s", u.Status)
	}

	if u.Gender != MALE && u.Gender != FEMALE && u.Gender != OTHER {
		return false, fmt.Errorf("gênero inválido: %s", u.Gender)
	}

	return true, nil
}

func (u *User) Enable() error {
	if u.Status == ENABLED {
		return errors.New("o usuário já está ativo")
	}

	u.Status = ENABLED

	_, err := u.Validate()

	return err
}

func (u *User) Disable() error {
	if u.Status == DISABLED {
		return errors.New("o usuário já está inativo")
	}

	u.Status = DISABLED

	_, err := u.Validate()

	return err
}

func (u *User) GetID() string           { return u.ID }
func (u *User) GetStatus() Status       { return u.Status }
func (u *User) GetName() string         { return u.Name }
func (u *User) GetEmail() string        { return u.Email }
func (u *User) GetGender() Gender       { return u.Gender }
func (u *User) GetBirthDate() time.Time { return u.BirthDate }
