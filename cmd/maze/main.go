package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/selmanj/maze/pkg/maze"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	// Mazes are 25x25 for now
	m := maze.NewMaze(25, 50)
	s := maze.NewRandomWalkSolver(&m)
	for !s.Step() {
		// continue...
	}
	fmt.Printf(m.String())
}
