package user

import (
	app "Go-Hexagonal/src/user/app"
	domain "Go-Hexagonal/src/user/domain"
	"time"

	"gorm.io/gorm"
)

type UserRepository struct {
	conn *gorm.DB
}

func NewUserRepo(c *gorm.DB) *UserRepository { // TODO: use generic interface for DB adapter
	return &UserRepository{conn: c}
}

type UserModel struct {
	gorm.Model
	ID        string        `gorm:"primaryKey"`
	Status    domain.Status `gorm:"type:user_status"`
	Name      string
	Email     string        `gorm:"unique"`
	Gender    domain.Gender `gorm:"type:user_gender"`
	BirthDate time.Time
	CreatedBy string
	CreatedAt time.Time
	UpdatedBy *string
	UpdatedAt *time.Time
}

func (r *UserRepository) Create(u domain.UserPort, createdBy string) (domain.UserPort, error) {
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

	return &domain.User{
		ID:        user.ID,
		Status:    user.Status,
		Name:      user.Name,
		Email:     user.Email,
		Gender:    user.Gender,
		BirthDate: user.BirthDate,
	}, nil
}

func (r *UserRepository) Get(f app.GetUserRepoFiltersDTO) (domain.UserPort, error) {
	user := &UserModel{}

	if res := r.conn.First(user, f); res.Error != nil {
		return nil, res.Error
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

func (r *UserRepository) List(f app.ListUsersRepoFiltersDTO) ([]domain.UserPort, error) {
	users := []UserModel{}

	if res := r.conn.Find(&users); res.Error != nil {
		return nil, res.Error
	}

	listUsers := make([]domain.UserPort, len(users))

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

func (r *UserRepository) Update(u domain.UserPort, updatedBy string) (domain.UserPort, error) {
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

	return &domain.User{
		ID:        user.ID,
		Status:    user.Status,
		Name:      user.Name,
		Email:     user.Email,
		Gender:    user.Gender,
		BirthDate: user.BirthDate,
	}, nil
}
