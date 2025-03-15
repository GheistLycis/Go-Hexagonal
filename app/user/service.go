package app

type UserService struct {
	repo UserRepoI
}

func NewUserService(r UserRepoI) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) Create(p CreateUserServicePayload, createdBy string) (UserI, error) {
	user, err := NewUser(p.Name, p.Email, p.Gender, p.BirthDate)
	if err != nil {
		return nil, err
	}

	if _, err := s.repo.Create(user, createdBy); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Disable(id string, updatedBy string) (UserI, error) {
	user, err := s.repo.Get(GetUserRepoFilters{ID: &id})
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

func (s *UserService) Enable(id string, updatedBy string) (UserI, error) {
	user, err := s.repo.Get(GetUserRepoFilters{ID: &id})
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

func (s *UserService) Get(f GetUserServiceFilters) (UserI, error) {
	user, err := s.repo.Get(GetUserRepoFilters(f))
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) List(f ListUsersServiceFilters) ([]UserI, error) {
	users, err := s.repo.List(ListUsersRepoFilters(f))
	if err != nil {
		return nil, err
	}

	return users, nil
}
