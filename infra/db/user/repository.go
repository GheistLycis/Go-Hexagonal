package db

import (
	app "Go-Hexagonal/app/user"
	"time"
)

type UserRepo struct{}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

type UserModel struct {
	ID        string     `db:"id"`
	Status    app.Status `db:"status"`
	Name      string     `db:"name"`
	Email     string     `db:"email"`
	Gender    app.Gender `db:"gender"`
	CreatedBy string     `db:"created_by"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedBy string     `db:"updated_by"`
	UpdatedAt time.Time  `db:"updated_at"`
}

func (r *UserRepo) Create(user app.UserI) (app.UserI, error) {
	return nil, nil
}

func (r *UserRepo) Get(filters app.GetUserRepoFilters) (app.UserI, error) {
	return nil, nil
}

func (r *UserRepo) List(filters app.ListUsersRepoFilters) ([]app.UserI, error) {
	return nil, nil
}

func (r *UserRepo) Update(user app.UserI) (app.UserI, error) {
	return nil, nil
}
