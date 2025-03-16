package db

import (
	app "Go-Hexagonal/app/user"
	"time"

	"gorm.io/gorm"
)

type UserRepo struct {
	conn *gorm.DB
}

func NewUserRepo(c *gorm.DB) *UserRepo { // TODO: use generic interface
	return &UserRepo{conn: c}
}

type UserModel struct {
	gorm.Model
	ID        string     `gorm:"primaryKey"`
	Status    app.Status `gorm:"type:user_status"`
	Name      string
	Email     string     `gorm:"unique"`
	Gender    app.Gender `gorm:"type:user_gender"`
	BirthDate time.Time
	CreatedBy string
	CreatedAt time.Time
	UpdatedBy *string
	UpdatedAt *time.Time
}

func (r *UserRepo) Create(u app.UserI, createdBy string) (app.UserI, error) {
	user := &UserModel{
		ID:        u.GetID(),
		Status:    u.GetStatus(),
		Name:      u.GetName(),
		Email:     u.GetEmail(),
		Gender:    u.GetGender(),
		BirthDate: u.GetBirthDate(),
		CreatedBy: createdBy,
		CreatedAt: time.Now(),
	}

	if res := r.conn.Create(&user); res.Error != nil {
		return nil, res.Error
	}

	return &app.User{
		ID:        user.ID,
		Status:    user.Status,
		Name:      user.Name,
		Email:     user.Email,
		Gender:    user.Gender,
		BirthDate: user.BirthDate,
	}, nil
}

func (r *UserRepo) Get(f app.GetUserRepoFilters) (app.UserI, error) {
	user := &UserModel{}

	if res := r.conn.First(user, f); res.Error != nil {
		return nil, res.Error
	}

	return &app.User{
		ID:        user.ID,
		Status:    user.Status,
		Name:      user.Name,
		Email:     user.Email,
		Gender:    user.Gender,
		BirthDate: user.BirthDate,
	}, nil
}

func (r *UserRepo) List(f app.ListUsersRepoFilters) ([]app.UserI, error) {
	users := []UserModel{}

	if res := r.conn.Find(&users); res.Error != nil {
		return nil, res.Error
	}

	listUsers := make([]app.UserI, len(users))

	for i, user := range users {
		listUsers[i] = &app.User{
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

func (r *UserRepo) Update(u app.UserI, updatedBy string) (app.UserI, error) {
	now := time.Now()
	user := &UserModel{
		ID:        u.GetID(),
		Status:    u.GetStatus(),
		Name:      u.GetName(),
		Email:     u.GetEmail(),
		Gender:    u.GetGender(),
		BirthDate: u.GetBirthDate(),
		UpdatedBy: &updatedBy,
		UpdatedAt: &now,
	}

	if res := r.conn.Updates(&user); res.Error != nil {
		return nil, res.Error
	}

	return &app.User{
		ID:        user.ID,
		Status:    user.Status,
		Name:      user.Name,
		Email:     user.Email,
		Gender:    user.Gender,
		BirthDate: user.BirthDate,
	}, nil
}
