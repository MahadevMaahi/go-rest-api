package router

import(
	"fmt"
	"github.com/gorilla/mux"
	db "github.com/MahadevMaahi/go-rest-api/db"
)

type MuxRouter struct {}

type RouterKit interface {
	SetUpRouter() *mux.Router
}

func (router *MuxRouter) SetUpRouter() *mux.Router {
	var dbRef db.DBKit = &(db.DBRef{})
	dbRef.SetUpDB()
	r := mux.NewRouter()

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Router setUp Done\n")
	return r
}