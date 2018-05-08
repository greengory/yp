package users

import "github.com/go-chi/chi"

func Routes() chi.Router {
	router := chi.NewRouter()

	user := User{}
	router.Get("/", user.GetAll)
	return router
}
