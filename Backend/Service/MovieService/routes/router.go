package routes

import "github.com/gorilla/mux"

// InitRouters : It is for initialing router for all routes
func InitRouters() *mux.Router {
	router := mux.NewRouter()
	router = MovieRoutes(router)
	return router
}
