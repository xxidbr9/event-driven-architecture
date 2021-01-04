package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/xxidbr9/MovieService/models"
)

// GetAllMovies : Get All Movie "TEST"
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Movie{ID: 1, Name: "Avatar The Lagend Of Japan"})
}
