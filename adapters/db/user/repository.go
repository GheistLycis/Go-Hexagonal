package db

import (
	app "Go-Hexagonal/app/user"
	app_ports "Go-Hexagonal/app/user/ports"
	"time"

	"gorm.io/gorm"
)

type UserRepository struct {
	conn *gorm.DB
}

func NewUserRepo(c *gorm.DB) *UserRepository { // TODO: use generic interface
	return &UserRepository{conn: c}
}

type UserModel struct {
	gorm.Model
	ID        string           `gorm:"primaryKey"`
	Status    app_ports.Status `gorm:"type:user_status"`
	Name      string
	Email     string           `gorm:"unique"`
	Gender    app_ports.Gender `gorm:"type:user_gender"`
	BirthDate time.Time
	CreatedBy string
	CreatedAt time.Time
	UpdatedBy *string
	UpdatedAt *time.Time
}

func (r *UserRepository) Create(u app_ports.UserPort, createdBy string) (app_ports.UserPort, error) {
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

func (r *UserRepository) Get(f app_ports.GetUserRepoFiltersDTO) (app_ports.UserPort, error) {
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

func (r *UserRepository) List(f app_ports.ListUsersRepoFiltersDTO) ([]app_ports.UserPort, error) {
	users := []UserModel{}

	if res := r.conn.Find(&users); res.Error != nil {
		return nil, res.Error
	}

	listUsers := make([]app_ports.UserPort, len(users))

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

func (r *UserRepository) Update(u app_ports.UserPort, updatedBy string) (app_ports.UserPort, error) {
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
