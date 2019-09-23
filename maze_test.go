package main

import (
	"fmt"
	"testing"
)

func TestCoordLeft(t *testing.T) {
	a := Coord{1, 2}
	expected := Coord{1, 1}

	got := a.Left()

	if got != expected {
		t.Errorf("%#v.Left() = %#v, want %#v", a, got, expected)
	}
}

func TestCoordRight(t *testing.T) {
	a := Coord{6, 3}
	expected := Coord{6, 4}

	got := a.Right()

	if got != expected {
		t.Errorf("%#v.Right() = %#v, want %#v", a, got, expected)
	}
}

func TestCoordUp(t *testing.T) {
	a := Coord{0, 8}
	expected := Coord{-1, 8}

	got := a.Up()

	if got != expected {
		t.Errorf("%#v.Up() = %#v, want %#v", a, got, expected)
	}
}

func TestCoordDown(t *testing.T) {
	a := Coord{99, 6}
	expected := Coord{100, 6}

	got := a.Down()

	if got != expected {
		t.Errorf("%#v.Down() = %#v, want %#v", a, got, expected)
	}
}

func TestMazeContainsCell(t *testing.T) {
	var containstests = []struct {
		in  Coord
		out bool
	}{
		{Coord{0, 0}, true},
		{Coord{7, 2}, true},
		{Coord{15, 1}, false},
		{Coord{1, 12}, false},
		{Coord{3, -1}, false},
		{Coord{-999999, 0}, false},
	}

	m := NewMaze(9, 11)

	for _, tt := range containstests {
		t.Run(fmt.Sprintf("%#v", tt.in), func(t *testing.T) {
			v := m.ContainsCell(tt.in)
			if v != tt.out {
				t.Errorf("got %t, want %t", v, tt.out)
			}
		})
	}
}
