package app

import ports "Go-Hexagonal/app/user/ports"

type UserService struct {
	repo ports.UserRepoPort
}

func NewUserService(r ports.UserRepoPort) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) Create(p ports.CreateUserServiceDTO, createdBy string) (ports.UserPort, error) {
	user, err := NewUser(p.Name, p.Email, p.Gender, p.BirthDate)
	if err != nil {
		return nil, err
	}

	if _, err := s.repo.Create(user, createdBy); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Disable(id string, updatedBy string) (ports.UserPort, error) {
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

func (s *UserService) Enable(id string, updatedBy string) (ports.UserPort, error) {
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

func (s *UserService) Get(f ports.GetUserServiceFiltersDTO) (ports.UserPort, error) {
	user, err := s.repo.Get(ports.GetUserRepoFiltersDTO(f))
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) List(f ports.ListUsersServiceFiltersDTO) ([]ports.UserPort, error) {
	users, err := s.repo.List(ports.ListUsersRepoFiltersDTO(f))
	if err != nil {
		return nil, err
	}

	return users, nil
}
