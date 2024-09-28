package user_service

import (
	"fmt"

	"github.com/alaa-aqeel/school-managment-system/app/models/user"
	"github.com/alaa-aqeel/school-managment-system/app/observers"
)

func New() *UserService {
	return &UserService{
		users: make(MapUsers),
		observers: []observers.Observer[*user.User]{
			&observers.TimestampsObserver[*user.User]{},
			&observers.IdObserver[*user.User]{},
		},
	}
}

func (s *UserService) GetAll() MapUsers {
	fmt.Println(s.users)
	return s.users
}
