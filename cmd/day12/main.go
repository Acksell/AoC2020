package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/Acksell/aoc2020/util"
)

type Action int

const (
	North Action = iota
	South
	East
	West
	Forward
	Left
	Right
)

type Instruction struct {
	action Action
	value  int
}

type Instructions []Instruction

func (i *Instructions) Load(s string) error {
	var action Action
	value, err := strconv.Atoi(s[1:])
	if err != nil {
		return err
	}
	switch s[0] {
	case 'N':
		action = North
	case 'S':
		action = South
	case 'E':
		action = East
	case 'W':
		action = West
	case 'F':
		action = Forward
	case 'L':
		action = Left
	case 'R':
		action = Right
	}
	*i = append(*i, Instruction{action, value})
	return nil
}

type Ship struct {
	x, y  int
	angle int
}

func (s *Ship) Execute(i Instruction) {
	action := i.action
	if action == Forward {
		if angle := s.angle % 360; angle == 0 {
			action = East
		} else if angle == 90 || angle == -270 {
			action = North
		} else if angle == 180 || angle == -180 {
			action = West
		} else if angle == 270 || angle == -90 {
			action = South
		}
	}
	switch action {
	case North:
		s.y += i.value
	case South:
		s.y -= i.value
	case East:
		s.x += i.value
	case West:
		s.x -= i.value
	case Left:
		s.angle += i.value
	case Right:
		s.angle -= i.value
	}
}

type Waypoint struct {
	x, y int
}

func (w *Waypoint) Rotate(deg int) {
	if deg == 0 {
		return
	} else if deg == 90 || deg == -270 {
		w.x, w.y = -w.y, w.x
	} else if deg == 180 || deg == -180 {
		w.x, w.y = -w.x, -w.y
	} else if deg == 270 || deg == -90 {
		w.x, w.y = w.y, -w.x
	}
}

type WaypointShip struct {
	x, y int
	w    Waypoint
}

func (ws *WaypointShip) Execute(i Instruction) {
	switch i.action {
	case Forward:
		ws.x += ws.w.x * i.value
		ws.y += ws.w.y * i.value
	case North:
		ws.w.y += i.value
	case South:
		ws.w.y -= i.value
	case East:
		ws.w.x += i.value
	case West:
		ws.w.x -= i.value
	case Left:
		ws.w.Rotate(+i.value)
	case Right:
		ws.w.Rotate(-i.value)
	}
}

const inputFilePath = "../../inputs/navigations.txt"

func main() {
	instr := make(Instructions, 0)
	util.ReadLines(inputFilePath, &instr)
	// part 1
	ship := Ship{}
	for _, i := range instr {
		ship.Execute(i)
	}
	fmt.Println(math.Abs(float64(ship.x)) + math.Abs(float64(ship.y)))
	// part 2
	w := Waypoint{10, 1}
	wship := WaypointShip{0, 0, w}
	for _, i := range instr {
		wship.Execute(i)
	}
	fmt.Println(math.Abs(float64(wship.x)) + math.Abs(float64(wship.y)))
}
