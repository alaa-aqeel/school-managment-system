package user_service

import (
	uuid "github.com/alaa-aqeel/school-managment-system/adapters/observer/id"
	"github.com/alaa-aqeel/school-managment-system/adapters/observer/timestamps"
	"github.com/alaa-aqeel/school-managment-system/app/domain"
	"github.com/alaa-aqeel/school-managment-system/app/interfaces/observer"
	"github.com/alaa-aqeel/school-managment-system/app/utils"
)

func New() *UserService {
	return &UserService{
		users: make(MapUsers),
		observers: []observer.Observer[*domain.User]{
			&timestamps.TimestampsObserver[*domain.User]{},
			&uuid.IdObserver[*domain.User]{},
		},
	}
}

func (s *UserService) notifyObserver(method string, user *domain.User) {
	utils.ExceuteObservers(s.observers, method, user)
}

func (s *UserService) GetAll() MapUsers {

	return s.users
}
