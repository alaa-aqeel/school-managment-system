package user_service

import "github.com/alaa-aqeel/school-managment-system/app/models/user"

type MapUsers map[string]user.User

type UserService struct {
	users MapUsers
}
