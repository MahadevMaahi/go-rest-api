package router

import(
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	db "github.com/MahadevMaahi/go-rest-api/db"
	model "github.com/MahadevMaahi/go-rest-api/db/model"
)

func getMovies(w http.ResponseWriter, r *http.Request) {
	var dbRef db.DBKit = &(db.DBRef{})
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dbRef.GetMovieList())
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	var dbRef db.DBKit = &(db.DBRef{})
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range dbRef.GetMovieList() {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(dbRef.GetMovieList())
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	var dbRef db.DBKit = &(db.DBRef{})
	w.Header().Set("Content-Type", "application/json")
	var movie model.Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	dbRef.AddMovie(movie)
	json.NewEncoder(w).Encode(dbRef.GetMovieList())
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	var dbRef db.DBKit = &(db.DBRef{})
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range dbRef.GetMovieList() {
		if item.ID == params["id"] {
			dbRef.SetMovieList(append(dbRef.GetMovieList()[:index], dbRef.GetMovieList()[index+1:]...))
			break;
		}
	}
	var movie model.Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	dbRef.AddMovie(movie)
	json.NewEncoder(w).Encode(dbRef.GetMovieList())
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	var dbRef db.DBKit = &(db.DBRef{})
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range dbRef.GetMovieList() {
		if item.ID == params["id"] {
			dbRef.SetMovieList(append(dbRef.GetMovieList()[:index], dbRef.GetMovieList()[index+1:]...))
			break;
		}
	}
	json.NewEncoder(w).Encode(dbRef.GetMovieList())
}