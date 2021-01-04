package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/xxidbr9/MovieService/routes"
)

// Request :
type Request struct {
	Hello string `json:"hello"`
}

func main() {
	r := routes.InitRouters()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(Request{Hello: "Hello World!"})
	}).Methods("GET")
	fmt.Println("\n Server Run ")
	http.ListenAndServe(":8001", r)
}
