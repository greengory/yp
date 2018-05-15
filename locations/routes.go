package locations

import (
	"github.com/go-chi/chi"
)

func Routes() chi.Router {
	r := chi.NewRouter()

	location := Location{}
	r.Post("/", location.Create)
	r.Get("/", location.GetAll)

	return r
}
