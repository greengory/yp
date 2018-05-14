package business

import "github.com/go-chi/chi"

func Routes() chi.Router {
	r := chi.NewRouter()
	business := Business{}
	r.Post("/", business.Create)
	return r
}
