package app

import (
	"errors"
	"fmt"
	"time"

	ports "Go-Hexagonal/app/user/ports"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type User struct {
	ID        string       `valid:"uuidv4"`
	Status    ports.Status `valid:"-"`
	Name      string       `valid:"-"`
	Email     string       `valid:"email"`
	Gender    ports.Gender `valid:"-"`
	BirthDate time.Time    `valid:"-"`
}

func NewUser(name string, email string, gender ports.Gender, BirthDate time.Time) (*User, error) {
	user := &User{
		ID:        uuid.NewV4().String(),
		Status:    ports.IN_ANALYSIS,
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

	if u.Status != ports.ENABLED && u.Status != ports.DISABLED && u.Status != ports.IN_ANALYSIS {
		return false, fmt.Errorf("status inválido: %s", u.Status)
	}

	if u.Gender != ports.MALE && u.Gender != ports.FEMALE && u.Gender != ports.OTHER {
		return false, fmt.Errorf("gênero inválido: %s", u.Gender)
	}

	return true, nil
}

func (u *User) Enable() error {
	if u.Status == ports.ENABLED {
		return errors.New("o usuário já está ativo")
	}

	u.Status = ports.ENABLED

	_, err := u.Validate()

	return err
}

func (u *User) Disable() error {
	if u.Status == ports.DISABLED {
		return errors.New("o usuário já está inativo")
	}

	u.Status = ports.DISABLED

	_, err := u.Validate()

	return err
}

func (u *User) GetID() string           { return u.ID }
func (u *User) GetStatus() ports.Status { return u.Status }
func (u *User) GetName() string         { return u.Name }
func (u *User) GetEmail() string        { return u.Email }
func (u *User) GetGender() ports.Gender { return u.Gender }
func (u *User) GetBirthDate() time.Time { return u.BirthDate }
