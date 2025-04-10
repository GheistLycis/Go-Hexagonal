package user

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var validate = validator.New() // TODO: implement singleton validator

type User struct {
	ID        uuid.UUID `validate:"uuid4"`
	Status    Status
	Name      string
	Email     string `validate:"email"`
	Gender    Gender
	BirthDate time.Time
}

func NewUser(name string, email string, gender Gender, birthDate time.Time) (*User, error) {
	user := &User{
		ID:        uuid.New(),
		Status:    IN_ANALYSIS,
		Name:      name,
		Email:     email,
		Gender:    gender,
		BirthDate: birthDate,
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) Validate() error {
	if err := validate.Struct(u); err != nil {
		return err
	}

	if u.Name == "" {
		return errors.New("o nome não pode ser vazio")
	}

	if u.Status != ENABLED && u.Status != DISABLED && u.Status != IN_ANALYSIS {
		return fmt.Errorf("status inválido: %s", u.Status)
	}

	if u.Gender != MALE && u.Gender != FEMALE && u.Gender != OTHER {
		return fmt.Errorf("gênero inválido: %s", u.Gender)
	}

	return nil
}

func (u *User) Enable() error {
	if u.Status == ENABLED {
		return errors.New("o usuário já está ativo")
	}

	u.Status = ENABLED

	return u.Validate()
}

func (u *User) Disable() error {
	if u.Status == DISABLED {
		return errors.New("o usuário já está inativo")
	}

	u.Status = DISABLED

	return u.Validate()
}

func (u *User) GetID() uuid.UUID        { return u.ID }
func (u *User) GetStatus() Status       { return u.Status }
func (u *User) GetName() string         { return u.Name }
func (u *User) GetEmail() string        { return u.Email }
func (u *User) GetGender() Gender       { return u.Gender }
func (u *User) GetBirthDate() time.Time { return u.BirthDate }
