package user_service

func New() *UserService {
	return &UserService{
		users: make(MapUsers),
	}
}

func (s *UserService) GetAll() MapUsers {
	return s.users
}
