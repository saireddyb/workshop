package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", saireddy)
	http.ListenAndServe(":3000",nil)
}

func saireddy(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello world"))
}