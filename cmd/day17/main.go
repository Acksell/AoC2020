package main

import (
	"errors"
	"fmt"
	"strings"
)

type Point interface {
	GetAdjacent() []Point
}

type Point3D struct {
	x, y, z int
}

func (p Point3D) GetAdjacent() []Point {
	adj := make([]Point, 0)
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			for k := -1; k < 2; k++ {
				if i == 0 && j == 0 && k == 0 {
					continue // don't include the point itself
				}
				adj = append(adj, Point3D{p.x + i, p.y + j, p.z + k})
			}
		}
	}
	return adj
}

type Point4D struct {
	x, y, z, w int
}

func (p Point4D) GetAdjacent() []Point {

	adj := make([]Point, 0)
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			for k := -1; k < 2; k++ {
				for l := -1; l < 2; l++ {
					if i == 0 && j == 0 && k == 0 && l == 0 {
						continue // don't include the point itself
					}
					adj = append(adj, Point4D{p.x + i, p.y + j, p.z + k, p.w + l})
				}
			}
		}
	}
	return adj
}

type State map[Point]bool

func (g State) ApplyRule(cube Point, nactive int) {
	if active, ok := g[cube]; ok && active {
		if !(nactive == 2 || nactive == 3) {
			g.SetCube(cube, false) // defer in order to toggle at end of this iteration, not while looping.
		}
	} else {
		if nactive == 3 { // defer in order to toggle at end of this iteration, not while looping
			g.SetCube(cube, true)
		}
	}
}

func (g State) CountActive(cubes []Point) int {
	nactive := 0
	for _, p := range cubes {
		if v, ok := g[p]; ok && v {
			nactive++
		}
	}
	return nactive
}

func (g State) Next() {
	for c := range g {
		adj := c.GetAdjacent()
		nactive := g.CountActive(adj)
		defer g.ApplyRule(c, nactive)
		for _, p := range adj { // also modify all neighbors.
			adj = p.GetAdjacent()
			nactive = g.CountActive(adj)
			defer g.ApplyRule(p, nactive)
		}
	}
}

func (g *State) SetCube(p Point, active bool) {
	(*g)[p] = active
}

func (g *State) Load(state string, dim int) {
	for i, row := range strings.Split(state, "\n") {
		for j, c := range row {
			var p Point
			if dim == 3 {
				p = Point3D{i, j, 0}
			} else if dim == 4 {
				p = Point4D{i, j, 0, 0}
			} else {
				panic(errors.New("Invalid dimension"))
			}
			if c == '#' {
				g.SetCube(p, true)
			}
		}
	}
}

const input = `...#..#.
.....##.
##..##.#
#.#.##..
#..#.###
...##.#.
#..##..#
.#.#..#.`

func main() {
	game := make(State)
	dim := 3
	game.Load(input, dim)
	for i := 0; i < 6; i++ {
		fmt.Println(i)
		game.Next()
	}
	sum := 0
	for _, active := range game {
		if active {
			sum++
		}
	}
	fmt.Println(sum)

	dim = 4
	game = make(State)
	game.Load(input, dim)

	for i := 0; i < 6; i++ {
		fmt.Println(i)
		game.Next()
	}
	sum = 0
	for _, active := range game {
		if active {
			sum++
		}
	}
	fmt.Println(sum)
}
