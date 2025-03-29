package user

import domain "Go-Hexagonal/src/user/domain"

type UserService struct {
	repo domain.UserRepoPort
}

func NewUserService(r domain.UserRepoPort) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) Create(p domain.CreateUserServiceDTO, createdBy string) (*domain.User, error) {
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
	user, err := s.repo.Get(domain.GetUserRepoFiltersDTO{ID: &id})
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
	user, err := s.repo.Get(domain.GetUserRepoFiltersDTO{ID: &id})
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

func (s *UserService) Get(f domain.GetUserServiceFiltersDTO) (*domain.User, error) {
	user, err := s.repo.Get(domain.GetUserRepoFiltersDTO(f))
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) List(f domain.ListUsersServiceFiltersDTO) ([]*domain.User, error) {
	users, err := s.repo.List(domain.ListUsersRepoFiltersDTO(f))
	if err != nil {
		return nil, err
	}

	return users, nil
}
