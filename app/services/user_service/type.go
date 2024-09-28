package user_service

import (
	"github.com/alaa-aqeel/school-managment-system/app/models/user"
	"github.com/alaa-aqeel/school-managment-system/app/observers"
)

type MapUsers map[string]user.User

type UserService struct {
	users     MapUsers
	observers []observers.Observer[*user.User]
}
