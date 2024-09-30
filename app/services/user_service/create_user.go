package user_service

import (
	"github.com/alaa-aqeel/school-managment-system/app/models/user"
	"github.com/alaa-aqeel/school-managment-system/helpers/hash"
)

func (s *UserService) CreateUser(_user user.User) error {
	hash, err := hash.Make(_user.Password)
	if err != nil {
		return err
	}
	_user.Password = hash
	s.RunObserver(_user, "Creating")
	s.users[_user.ID] = _user

	return nil
}

func (s *UserService) UpdateUser(_user user.User) error {
	s.RunObserver(_user, "Updating")
	s.users[_user.ID] = _user

	return nil
}
