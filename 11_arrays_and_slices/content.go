package arrays_and_slices

import "fmt"

var MOVIES_ON_DB = []string{
	"The Shawshank Redemption",
	"The Godfather",
	"The Dark Knight",
	"Pulp Fiction",
	"The Lord of the Rings: The Return of the King",
	"Schindler's List",
	"Fight Club",
	"Inception",
	"Forrest Gump",
	"The Matrix",
	"Goodfellas",
	"Star Wars: Episode IV - A New Hope",
	"The Silence of the Lambs",
	"The Usual Suspects",
	"Léon: The Professional",
	"Saving Private Ryan",
	"Gladiator",
}

func Slices() {
	// Performance with slices
	resultsFromApi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // think it as the ids from favorite movies

	// Pre allocate our slice of movies to have the same size from resultsFromApi, since we know the size of the results from the API will be 10
	movies := make([]string, 0, len(resultsFromApi)) // make(type, length, capacity), 0 is the initial length from our slice in this case

	for _, id := range resultsFromApi {
		movie := MOVIES_ON_DB[id]

		fmt.Println("Movies length", len(movies), "Movies capacity", cap(movies))

		movies = append(movies, movie)

		fmt.Println("Movies length", len(movies), "Movies capacity", cap(movies))
	}

	fmt.Println(movies)
}

func Matrix() {
	matrix := [][]int{}
	matrix3D := [][][]int{}

	fmt.Println(matrix)
	fmt.Println(matrix3D)
}
