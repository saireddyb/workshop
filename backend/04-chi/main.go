package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", helloWorld)
	

	http.ListenAndServe(":3000", router)
}
func helloWorld(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("hello world"))
}