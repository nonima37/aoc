// comment
package main

import (
	"aoc/functions"
)

type cords struct {
	x int
	y int
}

var cordsMap = map[string]cords{
	"v": {x: 0, y: -1},
	"^": {x: 0, y: 1},
	">": {x: 1, y: 0},
	"<": {x: -1, y: 0},
}

func day3RedoMain() int {
	inputString := functions.Read()
	visited := map[cords]bool{}
	santaCurCords := cords{0, 0}
	robotCurCords := cords{0, 0}

	for i, v := range inputString {
		if i%2 == 0 {
			santaCurCords = cords{x: santaCurCords.x + cordsMap[string(v)].x, y: santaCurCords.y + cordsMap[string(v)].y}
			visited[santaCurCords] = true
		} else {
			robotCurCords = cords{x: robotCurCords.x + cordsMap[string(v)].x, y: robotCurCords.y + cordsMap[string(v)].y}
			visited[robotCurCords] = true
		}
	}
	return len(visited)
}
