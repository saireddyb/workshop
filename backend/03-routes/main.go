package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/helloworld", helloworld)
	http.HandleFunc("/welcome", saireddy)
	http.ListenAndServe(":3000",nil)
}

func helloworld(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello worlds"))
}
func saireddy(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Weolcome to go lang"))
}