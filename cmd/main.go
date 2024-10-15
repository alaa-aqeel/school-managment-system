package main

import (
	"net/http"

	"github.com/alaa-aqeel/school-managment-system/helpers/cfg"
	"github.com/alaa-aqeel/school-managment-system/route"
	"github.com/go-chi/chi/v5"
)

func main() {
	cfg.InitConfig()
	r := chi.NewRouter()
	r.Mount("/api/v1", route.SetupRoute())
	http.ListenAndServe(":8080", r)
}
