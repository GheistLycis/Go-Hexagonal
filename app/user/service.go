package app


type UserService struct {
	repo UserRepoI
}

func NewUserService(r UserRepoI) *UserService {
	return &UserService{repo: r}
}


type GetUserRepoFilters struct {
	ID *string
	Email *string
}

type ListUsersRepoFilters struct {
	Name *string
	Status *string
}

type UserRepoI interface {
	Create(user UserI) (UserI, error)
	Get(filters GetUserRepoFilters) (UserI, error)
	List(filters ListUsersRepoFilters) ([]UserI, error)
	Update(user UserI) (UserI, error)
}


type CreateUserServicePayload struct {
	Name string
	Email string
}

type GetUserServiceFilters struct {
	ID *string
	Email *string
}

type ListUsersServiceFilters struct {
	Name *string
	Status *string
}

type UserServiceI interface {
	Create(payload CreateUserServicePayload) (UserI, error)
	Disable(id string) (UserI, error)
	Enable(id string) (UserI, error)
	Get(filters GetUserServiceFilters) (UserI, error)
	List(filters ListUsersServiceFilters) ([]UserI, error)
}

func (s *UserService) Create(p CreateUserServicePayload) (User, error) {
	user := NewUser()
	user.Name = name
	user.Price = price
	_, err := user.IsValid()
	if err != nil {
		return &User{}, err
	}
	result, err := s.Persistence.Save(user)
	if err != nil {
		return &User{}, err
	}
	return result, nil
}

func (s *UserService) Disable(id string) (User, error) {
	err := user.Disable()
	if err != nil {
		return &User{}, err
	}
	result, err := s.Persistence.Save(user)
	if err != nil {
		return &User{}, err
	}
	return result, nil
}

func (s *UserService) Enable(id string) (User, error) {
	err := user.Enable()
	if err != nil {
		return &User{}, err
	}
	result, err := s.Persistence.Save(user)
	if err != nil {
		return &User{}, err
	}
	return result, nil
}

func (s *UserService) Get(f GetUserServiceFilters) (User, error) {
	filters := GetUserRepoFilters{ID: &id}
	user, err := s.repo.Get(filters)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) List(f ListUsersServiceFilters) ([]User, error) {
	filters := GetUserRepoFilters{ID: &id}
	user, err := s.repo.Get(filters)

	if err != nil {
		return nil, err
	}

	return user, nil
}
