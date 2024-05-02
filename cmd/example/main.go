package main

import (
	"fmt"

	"github.com/swkwon/go-astar"
)

func main() {
	fmt.Println("start a* algorithm")
	// first case
	if m1, err := astar.MakeDefaultMap("map-1", 10, 10, astar.Coordinate{X: 3, Y: 2}, astar.Coordinate{X: 6, Y: 9},
		[]astar.Coordinate{
			{X: 2, Y: 4},
			{X: 3, Y: 4},
			{X: 4, Y: 4},
			{X: 5, Y: 4},
			{X: 5, Y: 8},
			{X: 6, Y: 8},
			{X: 7, Y: 8},
			{X: 8, Y: 8},
		}); err == nil {
		m1.Find()
		m1.Print()
	} else {
		panic(err)
	}

	// second case
	if m2, err := astar.MakeDefaultMap("map-2", 10, 10, astar.Coordinate{X: 3, Y: 2}, astar.Coordinate{X: 6, Y: 9},
		[]astar.Coordinate{
			{X: 2, Y: 4},
			{X: 3, Y: 4},
			{X: 4, Y: 4},
			{X: 5, Y: 4},
			{X: 5, Y: 8},
			{X: 6, Y: 8},
			{X: 7, Y: 8},
			{X: 8, Y: 8},
			{X: 5, Y: 9},
		}); err == nil {
		m2.Find()
		m2.Print()
	} else {
		panic(err)
	}

	fmt.Println("end a* algorithm")
}
