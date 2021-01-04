package routes

import (
	"github.com/gorilla/mux"
	"github.com/xxidbr9/MovieService/controllers"
)

// MovieRoutes : All Movie Route Handle In Here
func MovieRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/movies", controllers.GetAllMovies).Methods("GET")
	return router
}
