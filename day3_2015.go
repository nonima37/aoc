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
	curCord := cords{x: 0, y: 0}

	for i, v := range input {
		curCord = cords{x: m[string(v)].x, y: m[string(v)].y}
		ok := true
		if i%2 == 0 {
			santaCord.x += curCord.x
			santaCord.y += curCord.y
			_, ok = visitedCords[santaCord]
			if !ok {
				visitedCords[santaCord] = true
			}
		} else {
			robotCord.x += curCord.x
			robotCord.y += curCord.y
			_, ok = visitedCords[robotCord]
			if !ok {
				visitedCords[robotCord] = true
			}
		}
	}
	fmt.Println(len(visitedCords))
}
