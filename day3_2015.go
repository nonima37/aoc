package main

import (
	"aoc/functions"
	"fmt"
)

type cords struct {
	x int
	y int
}

var m = map[string]cords{
	"v": cords{x: 0, y: -1},
	"^": cords{x: 0, y: 1},
	">": cords{x: 1, y: 0},
	"<": cords{x: -1, y: 0},
}

func main() {
	input := functions.Read()
	visitedCords := map[cords]bool{}
	santaCord := cords{x: 0, y: 0}
	robotCord := cords{x: 0, y: 0}

	curCords := &santaCord
	for i, v := range input {
		fmt.Println(santaCord)
		if i%2 == 0 {
			curCords = &santaCord
		} else {
			curCords = &robotCord
		}

		curCords.x += m[string(v)].x
		curCords.y += m[string(v)].y
		_, ok := visitedCords[*curCords]

		if !ok {
			visitedCords[*curCords] = true
		}
	}
	fmt.Println(len(visitedCords))
}
