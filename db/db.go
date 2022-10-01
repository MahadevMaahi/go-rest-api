package db

import(
	"fmt"
	"math/rand"
	"strconv"
	"time"
	model "github.com/MahadevMaahi/jwt-api-authentication/db/model"
)

var movieList []model.Movie

type DBKit interface {
	SetUpDB()
	GetMovieList() []model.Movie
	AddMovie(movie model.Movie)
	SetMovieList(list []model.Movie)
}

type DBRef struct {}

func (dbRef *DBRef) SetUpDB() {
	movieList = append(movieList, model.Movie{
		ID: "1",
		Isbn: strconv.Itoa(rand.Intn(1024)),
		Title: "Movie 1",
		Imdb_rating: (rand.Float32() * 10),
		Release_date: time.Now(),
		Director: &model.Director{
			FirstName: "Director - 1",
			LastName: "Movie - 1",
		},
	})
	movieList = append(movieList, model.Movie{
		ID: "2",
		Isbn: strconv.Itoa(rand.Intn(1024)),
		Title: "Movie 2",
		Imdb_rating: (rand.Float32() * 10),
		Release_date: time.Now(),
		Director: &model.Director{
			FirstName: "Director - 2",
			LastName: "Movie - 2",
		},
	})
	fmt.Println("DB Set up Done")
}

func (dbRef *DBRef) GetMovieList() []model.Movie {
	return movieList
}

func (dbRef *DBRef) SetMovieList(list []model.Movie) {
	movieList = list
}

func (dbRef *DBRef) AddMovie(movie model.Movie) {
	movieList = append(movieList, movie)
}