package user

import domain "Go-Hexagonal/src/user/domain"

type UserService struct {
	repo UserRepoPort
}

func NewUserService(r UserRepoPort) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) Create(p CreateUserServiceDTO, createdBy string) (domain.UserPort, error) {
	user, err := domain.NewUser(p.Name, p.Email, p.Gender, p.BirthDate)
	if err != nil {
		return nil, err
	}

	if _, err := s.repo.Create(user, createdBy); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Disable(id string, updatedBy string) (domain.UserPort, error) {
	user, err := s.repo.Get(GetUserRepoFiltersDTO{ID: &id})
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

func (s *UserService) Enable(id string, updatedBy string) (domain.UserPort, error) {
	user, err := s.repo.Get(GetUserRepoFiltersDTO{ID: &id})
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

func (s *UserService) Get(f GetUserServiceFiltersDTO) (domain.UserPort, error) {
	user, err := s.repo.Get(GetUserRepoFiltersDTO(f))
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) List(f ListUsersServiceFiltersDTO) ([]domain.UserPort, error) {
	users, err := s.repo.List(ListUsersRepoFiltersDTO(f))
	if err != nil {
		return nil, err
	}

	return users, nil
}
