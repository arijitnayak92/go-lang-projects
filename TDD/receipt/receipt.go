package receipt

import (
	"github.com/arijitnayak92/taskAfford/TDD/movie"
)

type Receipt struct {
	movie []movie.Movie
}

func NewReceipt() *Receipt {
	return &Receipt{}
}
