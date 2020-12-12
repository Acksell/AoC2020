package main

import (
	"errors"
	"fmt"

	"github.com/Acksell/aoc2020/util"
)

// Tile is an enum type used to represent the type of tile in the grid.
type Tile int

const (
	Floor Tile = iota
	Occupied
	Empty
)

type Grid [][]Tile

// Vec is an integer vector.
type Vec struct {
	i, j int
}

// Add two vectors
func (v Vec) Add(v2 Vec) Vec {
	return Vec{v.i + v2.i, v.j + v2.j}
}

// Mul -tiply with a scalar
func (v Vec) Mul(c int) Vec {
	return Vec{c * v.i, c * v.j}
}

func (p Tile) String() string {
	characters := []string{".", "#", "L"}
	return fmt.Sprint(characters[p])
}

func (g Grid) String() string {
	s := ""
	for _, row := range g {
		s += fmt.Sprintln(row)
	}
	return s
}

func (g *Grid) Load(s string) error {
	row := make([]Tile, 0)
	for _, p := range s {
		var point Tile
		switch p {
		case '.':
			point = Floor
		case '#':
			point = Occupied
		case 'L':
			point = Empty
		}
		row = append(row, point)
	}
	*g = append(*g, row)
	return nil
}

func (g Grid) Size() (N, M int) {
	N = len(g)
	M = len(g[0])
	return
}

// WithinBounds takes a number of coordinates and returns a slice of the same
// length where the corresponding elements are true if coordinate is within bounds.
func (g Grid) WithinBounds(c Vec) bool {
	N, M := g.Size()
	return 0 <= c.i && c.i < N && 0 <= c.j && c.j < M
}

func (g *Grid) Set(i, j int, p Tile) {
	(*g)[i][j] = p
}

func (g Grid) Get(i, j int) (Tile, error) {
	if !g.WithinBounds(Vec{i, j}) {
		return Floor, errors.New("Out of bounds")
	}
	return g[i][j], nil
}

func getAdjacent(i, j int) []Vec {
	return []Vec{Vec{i - 1, j - 1}, Vec{i - 1, j}, Vec{i - 1, j + 1}, Vec{i, j - 1}, Vec{i, j + 1}, Vec{i + 1, j - 1}, Vec{i + 1, j}, Vec{i + 1, j + 1}}
}

func (g Grid) AdjacentTiles(i, j int) []Tile {
	adj := getAdjacent(i, j)
	tiles := make([]Tile, 0)
	for _, v := range adj {
		tile, _ := g.Get(v.i, v.j)
		tiles = append(tiles, tile)
	}
	return tiles
}

func (g *Grid) ScanDir(i, j int, direction Vec) Tile {
	var foundTile Tile
	ipos, jpos := i, j
	for foundTile == Floor || (ipos == i && jpos == j) {
		ipos, jpos = ipos+direction.i, jpos+direction.j
		tile, err := g.Get(ipos, jpos)
		if err != nil {
			return foundTile // out of bounds, return floor tile.
		}
		foundTile = tile

	}
	return foundTile
}

func (g Grid) AdjScan(i, j int) []Tile {
	directions := make([]Vec, 0)
	adj := getAdjacent(i, j)
	pos := Vec{i, j}
	tiles := make([]Tile, len(adj))
	for _, v := range adj { // subtract the position from the adjacencies to get the set of directions
		dir := v.Add(pos.Mul(-1))
		directions = append(directions, dir)
	}
	for _, dir := range directions {
		tile := g.ScanDir(i, j, dir)
		tiles = append(tiles, tile)
	}
	return tiles
}

func CountOccupied(tiles []Tile) int {
	numOccupied := 0
	for _, tile := range tiles {
		if tile == Occupied {
			numOccupied++
		}
	}
	return numOccupied
}

// Next step in the automata. Returns true if anything changed, false otherwise.
func (g *Grid) Next1() bool {
	changed := false
	for i, row := range *g {
		for j, tile := range row {
			tiles := g.AdjacentTiles(i, j)
			occupied := CountOccupied(tiles)
			if tile == Empty && occupied == 0 {
				defer g.Set(i, j, Occupied) // defer because we don't want to change while still looping.
				changed = true
			} else if tile == Occupied && occupied >= 4 {
				defer g.Set(i, j, Empty) // defer because we don't want to change while still looping.
				changed = true
			}
		}
	}
	return changed
}

// Next step in the automata. Returns true if anything changed, false otherwise.
// can refactor with Next1 but whatever. Too much effort.
func (g *Grid) Next2() bool {
	changed := false
	for i, row := range *g {
		for j, tile := range row {
			tiles := g.AdjScan(i, j)
			occupied := CountOccupied(tiles)
			if tile == Empty && occupied == 0 {
				defer g.Set(i, j, Occupied) // defer because we don't want to change while still looping.
				changed = true
			} else if tile == Occupied && occupied >= 5 {
				defer g.Set(i, j, Empty) // defer because we don't want to change while still looping.
				changed = true
			}
		}
	}
	return changed
}

const inputFilePath = "../../inputs/ferry_seats.txt"

func main() {
	grid := make(Grid, 0)
	util.ReadLines(inputFilePath, &grid)
	for grid.Next1() {
		// step forward until the grid doesn't change.
	}
	occ := 0
	for _, row := range grid {
		occ += CountOccupied(row)
	}
	fmt.Println(occ)
	// Reload input into grid.
	grid = make(Grid, 0)
	util.ReadLines(inputFilePath, &grid)
	for grid.Next2() {
		// step forward until the grid doesn't change.
	}
	occ = 0
	for _, row := range grid {
		occ += CountOccupied(row)
	}
	fmt.Println(occ)
}
