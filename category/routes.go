package category

import (
	"github.com/go-chi/chi"
)

func Routes() chi.Router {
	r := chi.NewRouter()
	category := Category{}
	r.Post("/", category.Create)
	r.Get("/", category.GetAll)
	return r
}
