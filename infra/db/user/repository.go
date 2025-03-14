package db

import (
	app "Go-Hexagonal/app/user"
)


type UserRepo struct {}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

type UserModel struct {
	ID     string `db:"id"`
	Status   string  `db:"status"`
	Name   string  `db:"name"`
	Email  string`db:"email"`
	CreatedBy string `db:"created_by"`
	CreatedAt string `db:"created_at"`
	UpdatedBy string `db:"updated_by"`
	UpdatedAt string `db:"updated_at"`
}

func (r *UserRepo) Create(user app.UserI) (app.UserI, error) {
	user := &UserModel{
		ID: "",
		Status: "",
		Name: "",
		Email: "",
		CreatedBy: "",
		CreatedAt: "",
		UpdatedBy: "",
		UpdatedAt: "",
	}


	return user, nil
}

func (r *UserRepo) Get(filters app.GetUserRepoFilters) (app.UserI, error) {
	user := &UserModel{
		ID: "",
		Status: "",
		Name: "",
		Email: "",
		CreatedBy: "",
		CreatedAt: "",
		UpdatedBy: "",
		UpdatedAt: "",
	}


	return user, nil
}

func (r *UserRepo) List(filters app.ListUsersRepoFilters) ([]app.UserI, error) {
	user := &UserModel{
		ID: "",
		Status: "",
		Name: "",
		Email: "",
		CreatedBy: "",
		CreatedAt: "",
		UpdatedBy: "",
		UpdatedAt: "",
	}


	return user, nil
}

func (r *UserRepo) Update(user app.UserI) (app.UserI, error) {
	user := &UserModel{
		ID: "",
		Status: "",
		Name: "",
		Email: "",
		CreatedBy: "",
		CreatedAt: "",
		UpdatedBy: "",
		UpdatedAt: "",
	}


	return user, nil
}