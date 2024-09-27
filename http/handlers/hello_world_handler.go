package handlers

import (
	"net/http"

	"github.com/alaa-aqeel/school-managment-system/config"
)

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello world ^"))
	w.Write([]byte(config.Database.GetEnv("DB_HOST")))
}
