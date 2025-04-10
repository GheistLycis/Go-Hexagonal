package user

import (
	domain "Go-Hexagonal/src/user/domain"
	ports "Go-Hexagonal/src/user/ports"
)

type UserRepository struct {
	db ports.DBConnectionPort
}

func NewUserRepo(db ports.DBConnectionPort) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Create(u ports.UserPort, c string) (*domain.User, error) {
	user := &UserModel{
		ID:        u.GetID(),
		Status:    u.GetStatus(),
		Name:      u.GetName(),
		Email:     u.GetEmail(),
		Gender:    u.GetGender(),
		BirthDate: u.GetBirthDate(),
		CreatedBy: c,
	}

	if err := r.db.Insert(&user); err != nil {
		return nil, err
	}

	return &domain.User{
		ID:        user.ID,
		Status:    user.Status,
		Name:      user.Name,
		Email:     user.Email,
		Gender:    user.Gender,
		BirthDate: user.BirthDate,
	}, nil
}

func (r *UserRepository) Get(f ports.GetUserRepoFiltersDTO) (*domain.User, error) {
	user := &UserModel{}

	if err := r.db.Get(&user, f); err != nil {
		return nil, err
	}

	return &domain.User{
		ID:        user.ID,
		Status:    user.Status,
		Name:      user.Name,
		Email:     user.Email,
		Gender:    user.Gender,
		BirthDate: user.BirthDate,
	}, nil
}

func (r *UserRepository) List(f ports.ListUsersRepoFiltersDTO) ([]*domain.User, error) {
	users := []UserModel{}

	if err := r.db.List(&users); err != nil {
		return nil, err
	}

	listUsers := make([]*domain.User, len(users))

	for i, user := range users {
		listUsers[i] = &domain.User{
			ID:        user.ID,
			Status:    user.Status,
			Name:      user.Name,
			Email:     user.Email,
			Gender:    user.Gender,
			BirthDate: user.BirthDate,
		}
	}

	return listUsers, nil
}

func (r *UserRepository) Update(u ports.UserPort, ub string) (*domain.User, error) {
	user := &UserModel{
		ID:        u.GetID(),
		Status:    u.GetStatus(),
		Name:      u.GetName(),
		Email:     u.GetEmail(),
		Gender:    u.GetGender(),
		BirthDate: u.GetBirthDate(),
		UpdatedBy: &ub,
	}

	if err := r.db.Update(&user); err != nil {
		return nil, err
	}

	return &domain.User{
		ID:        user.ID,
		Status:    user.Status,
		Name:      user.Name,
		Email:     user.Email,
		Gender:    user.Gender,
		BirthDate: user.BirthDate,
	}, nil
}
