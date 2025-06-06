package user

import (
	user_mock "Go-Hexagonal/src/user/mocks"
	user "Go-Hexagonal/src/user/ports"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	service := NewUserService(
		user_mock.NewMockUserRepoPort(nil),
	)

	user, err := service.Create(user.CreateUserServiceDTO{
		Name:      "Name",
		Email:     "mock@email.com",
		Gender:    "Gender",
		BirthDate: time.Now(),
	}, "Created By")

}
