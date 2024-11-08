package user_service

import (
	"github.com/alaa-aqeel/school-managment-system/app/domain"
	"github.com/alaa-aqeel/school-managment-system/app/interfaces/observer"
)

type MapUsers map[string]domain.User

type UserService struct {
	users     MapUsers
	observers []observer.Observer[*domain.User]
}
