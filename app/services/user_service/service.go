package user_service

import (
	"github.com/alaa-aqeel/school-managment-system/app/interfaces"
	"github.com/alaa-aqeel/school-managment-system/app/models/user"
	"github.com/alaa-aqeel/school-managment-system/app/observers"
)

func New() *UserService {
	return &UserService{
		users: make(MapUsers),
		observers: []interfaces.Observer[*user.User]{
			&observers.TimestampsObserver[*user.User]{},
			&observers.IdObserver[*user.User]{},
		},
	}
}

func (s *UserService) GetAll() MapUsers {

	return s.users
}

func (s *UserService) RunObserver(_user user.User, method string) {

	observers.ExceuteObservers(s.observers, "Creating", &_user)
}
