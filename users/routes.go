package users

import "github.com/go-chi/chi"

func Routes() chi.Router {
	router := chi.NewRouter()
	user := User{}
	router.Get("/", user.GetAll)
	router.Post("/", user.Register)
	router.Post("/auth", user.Login)
	return router
}
