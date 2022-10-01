package model

import(
	"time"
)

type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"` 
	Imdb_rating float32 `json:"imdbRating"`
	Release_date time.Time `json:"releaseDate"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}