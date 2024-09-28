package handlers

import (
	"net/http"

	"github.com/alaa-aqeel/school-managment-system/app/models/user"
	"github.com/alaa-aqeel/school-managment-system/app/services/user_service"
	"github.com/alaa-aqeel/school-managment-system/helpers/response"
)

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {

	userService := user_service.New()
	userService.CreateUser(user.User{
		Username: "admin 2",
		Password: "admin 2",
		IsActive: true,
	})
	userService.CreateUser(user.User{
		Username: "admin",
		Password: "admin",
		IsActive: true,
	})
	response.Response(w).Json(userService.GetAll())
}
