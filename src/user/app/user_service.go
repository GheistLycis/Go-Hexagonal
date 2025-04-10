package user

import (
	domain "Go-Hexagonal/src/user/domain"
	ports "Go-Hexagonal/src/user/ports"
)

type UserService struct {
	repo ports.UserRepoPort
}

func NewUserService(r ports.UserRepoPort) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) Create(p ports.CreateUserServiceDTO, createdBy string) (*domain.User, error) {
	user, err := domain.NewUser(p.Name, p.Email, p.Gender, p.BirthDate)
	if err != nil {
		return nil, err
	}

	if _, err := s.repo.Create(user, createdBy); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Disable(id string, updatedBy string) (*domain.User, error) {
	user, err := s.repo.Get(ports.GetUserRepoFiltersDTO{ID: &id})
	if err != nil {
		return nil, err
	}

	if err := user.Disable(); err != nil {
		return nil, err
	}

	if _, err := s.repo.Update(user, updatedBy); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Enable(id string, updatedBy string) (*domain.User, error) {
	user, err := s.repo.Get(ports.GetUserRepoFiltersDTO{ID: &id})
	if err != nil {
		return nil, err
	}

	if err := user.Enable(); err != nil {
		return nil, err
	}

	if _, err := s.repo.Update(user, updatedBy); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Get(f ports.GetUserServiceFiltersDTO) (*domain.User, error) {
	user, err := s.repo.Get(ports.GetUserRepoFiltersDTO(f))
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) List(f ports.ListUsersServiceFiltersDTO) ([]*domain.User, error) {
	users, err := s.repo.List(ports.ListUsersRepoFiltersDTO(f))
	if err != nil {
		return nil, err
	}

	return users, nil
}
