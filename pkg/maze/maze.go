package maze

import (
	"fmt"
	"math/rand"
	"strings"
)

type Coord struct {
	Row int
	Col int
}

func (a Coord) Left() Coord {
	return Coord{Row: a.Row, Col: a.Col - 1}
}

func (a Coord) Right() Coord {
	return Coord{Row: a.Row, Col: a.Col + 1}
}

func (a Coord) Up() Coord {
	return Coord{Row: a.Row - 1, Col: a.Col}
}

func (a Coord) Down() Coord {
	return Coord{Row: a.Row + 1, Col: a.Col}
}

type Cell struct {
	Up    bool
	Down  bool
	Left  bool
	Right bool
}

type Maze struct {
	Height int
	Width  int
	Cells  [][]Cell
}

func NewMaze(height, width int) Maze {
	m := Maze{Height: height, Width: width}
	m.Cells = make([][]Cell, height)
	for i := range m.Cells {
		m.Cells[i] = make([]Cell, width)
	}

	return m
}

func (m Maze) ContainsCell(a Coord) bool {
	return a.Row >= 0 && a.Row < m.Height && a.Col >= 0 && a.Col < m.Width
}

func (m Maze) OpenLeft(a Coord) {
	if m.ContainsCell(a) {
		m.Cells[a.Row][a.Col].Left = true
	}
	b := a.Left()
	if m.ContainsCell(b) {
		m.Cells[b.Row][b.Col].Right = true
	}
}

func (m Maze) OpenRight(a Coord) {
	m.OpenLeft(a.Right())
}

func (m Maze) OpenUp(a Coord) {
	if m.ContainsCell(a) {
		m.Cells[a.Row][a.Col].Up = true
	}
	b := a.Up()
	if m.ContainsCell(b) {
		m.Cells[b.Row][b.Col].Down = true
	}
}

func (m Maze) OpenDown(a Coord) {
	m.OpenUp(a.Down())
}

func (m Maze) ConnectAdjacent(a Coord, b Coord) {
	if a.Left() == b {
		m.OpenLeft(a)
	} else if a.Right() == b {
		m.OpenRight(a)
	} else if a.Up() == b {
		m.OpenUp(a)
	} else if a.Down() == b {
		m.OpenDown(a)
	} // TODO Error?
}

func (m Maze) String() string {
	if m.Height == 0 || m.Width == 0 {
		// Nothing to print
		return ""
	}

	var b strings.Builder

	// Print the top row
	fmt.Fprintf(&b, "+")
	for col := 0; col < m.Width; col++ {
		if m.Cells[0][col].Up {
			fmt.Fprintf(&b, "  ")
		} else {
			fmt.Fprintf(&b, "--")
		}
		if col < m.Width-1 {
			if m.Cells[0][col].Right {
				fmt.Fprintf(&b, "-")
			} else {
				fmt.Fprintf(&b, "+")
			}
		}
	}
	fmt.Fprintf(&b, "+\n")

	for row := 0; row < m.Height; row++ {
		if m.Cells[row][0].Left {
			fmt.Fprintf(&b, " ")
		} else {
			fmt.Fprintf(&b, "|")
		}
		for col := 0; col < m.Width; col++ {
			fmt.Fprintf(&b, "  ")
			if m.Cells[row][col].Right {
				fmt.Fprintf(&b, " ")
			} else {
				fmt.Fprintf(&b, "|")
			}
		}
		fmt.Fprintf(&b, "\n")
		// Now print the second part
		if m.Cells[row][0].Down {
			fmt.Fprintf(&b, "|")
		} else {
			fmt.Fprintf(&b, "+")
		}
		for col := 0; col < m.Width; col++ {
			if m.Cells[row][col].Down {
				fmt.Fprintf(&b, "  ")
			} else {
				fmt.Fprintf(&b, "--")
			}
			if m.Cells[row][col].Down && col != m.Width-1 && m.Cells[row][col+1].Down {
				fmt.Fprintf(&b, "|")
			} else if m.Cells[row][col].Down && col == m.Width-1 {
				fmt.Fprintf(&b, "|")
			} else if m.Cells[row][col].Right && row != m.Height-1 && m.Cells[row+1][col].Right {
				fmt.Fprintf(&b, "-")
			} else if m.Cells[row][col].Right && row == m.Height-1 {
				fmt.Fprintf(&b, "-")
			} else {
				fmt.Fprintf(&b, "+")
			}
		}
		fmt.Fprintf(&b, "\n")
	}

	return b.String()
}

type RandomWalkSolver struct {
	// The maze to solve
	m *Maze
	// next is a list of next coordinates to try. Note that the inner array, if any is assumed to be of length 2
	next []Coord

	visited map[Coord]bool
}

func NewRandomWalkSolver(m *Maze) RandomWalkSolver {
	// TODO Open a start AND an end (maybe label em?)
	var first Coord
	// Choose top or left for start bottom, right
	wallChoice := rand.Intn(2)
	switch wallChoice {
	case 0:
		// Top
		first.Row = 0
		first.Col = rand.Intn(m.Width-2) + 1
		m.OpenUp(first)
	case 1:
		// Left
		first.Row = rand.Intn(m.Height-2) + 1
		first.Col = 0
		m.OpenLeft(first)
	}

	// Now choose bottom or right for end
	wallChoice = rand.Intn(2)
	switch wallChoice {
	case 0:
		// Right
		first.Row = rand.Intn(m.Height-2) + 1
		first.Col = m.Width - 1
		m.OpenRight(first)
	case 1:
		// Bottom
		first.Row = m.Height - 1
		first.Col = rand.Intn(m.Width-2) + 1
		m.OpenDown(first)
	}

	return RandomWalkSolver{
		m:       m,
		next:    []Coord{first},
		visited: make(map[Coord]bool)}
}

func (r *RandomWalkSolver) Step() bool {
	if len(r.next) == 0 {
		return true
	}
	cur := r.next[len(r.next)-1]
	var choices []Coord
	// Make a list of choices assuming they're valid
	// TODO change to directions?
	for _, a := range []Coord{cur.Up(), cur.Right(), cur.Down(), cur.Left()} {
		if r.m.ContainsCell(a) && !r.visited[a] {
			choices = append(choices, a)
		}
	}

	if len(choices) == 0 {
		r.next = r.next[:len(r.next)-1]
	} else {
		// chooce random, connect, and save
		upcoming := choices[rand.Intn(len(choices))]
		r.m.ConnectAdjacent(cur, upcoming)
		r.next = append(r.next, upcoming)
	}
	r.visited[cur] = true

	return false
}
