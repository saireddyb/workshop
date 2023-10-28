package functions

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"workshop/model"

	"github.com/go-chi/chi/v5"
)


var movies = []model.Movies{
	{Id: 0, Name: "KGF", Rating: 10, Year: 2018, ImgUrl: "somerandomurl", Category: "Action", Watched: false},
	{Id: 3, Name: "KGF2", Rating: 8, Year: 2022, ImgUrl: "somerandomurl1", Category: "Action2", Watched: false},
	{Id: 5, Name: "Avator", Rating: 10, Year: 2012, ImgUrl: "somerandomurl3", Category: "Animation", Watched: false},
	{Id: 99, Name: "Avator2", Rating: 9, Year: 2022, ImgUrl: "somerandomurl4", Category: "Thriller", Watched: false},
}

func ListMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(movies)
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var resultId int;
	for index , movie	:= range movies {
		if movie.Id == id {
			resultId = index
			break
		}
	}
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(movies[resultId])
}

func UpdateMovie(w http.ResponseWriter, r *http.Request){
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var body model.Movies
	_ = json.NewDecoder(r.Body).Decode(&body)
	for index, movie := range movies {
		if movie.Id == id {
			movie.Watched = body.Watched
			movies[index] = movie
			break
		}
	}
	w.Write([]byte("Movie is now watched"))
}

func AddMovie(w http.ResponseWriter, r *http.Request){
	var newMovie model.Movies
	_ = json.NewDecoder(r.Body).Decode(&newMovie)

	if(newMovie.Id == 0 || newMovie.Name == "") {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	for _, movie := range movies {
		if newMovie.Id == movie.Id {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	movies = append(movies, newMovie)
	fmt.Print(newMovie)
}