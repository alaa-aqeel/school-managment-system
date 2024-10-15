package user_service

import (
	"github.com/alaa-aqeel/school-managment-system/app/domain"
	"github.com/alaa-aqeel/school-managment-system/app/utils"
)

func (s *UserService) CreateUser(_user domain.User) error {
	hash, err := utils.MakeHash(_user.Password)
	if err != nil {
		return err
	}
	_user.Password = hash
	s.notifyObserver("Creating", &_user)
	s.users[_user.ID] = _user

	return nil
}

func (s *UserService) UpdateUser(_user domain.User) error {
	s.notifyObserver("Updating", &_user)
	s.users[_user.ID] = _user

	return nil
}
