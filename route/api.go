package route

import (
	"github.com/alaa-aqeel/school-managment-system/http/handlers"
	"github.com/go-chi/chi/v5"
)

func SetupRoute() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/hello-world", handlers.HelloWorldHandler)

	return r
}
