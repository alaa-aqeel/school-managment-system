package user_service

import (
	"time"

	"github.com/alaa-aqeel/school-managment-system/app/models/user"
	"github.com/alaa-aqeel/school-managment-system/helpers"
	"github.com/alaa-aqeel/school-managment-system/helpers/hash"
)

func (s *UserService) CreateUser(user user.User) error {
	user.ID = helpers.UUID()
	hash, err := hash.Make(user.Password)
	if err != nil {
		return err
	}
	user.Password = hash
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	s.users[user.ID] = user

	return nil
}
