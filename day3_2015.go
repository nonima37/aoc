package main

import (
	"fmt"
)

var m = map[string]int{
	"v": -1,
	"^": 1,
	">": 1,
	"<": -1,
}

type cords struct {
	x int
	y int
}

func main() {
	//input := functions.Read()
	input := "^>v<"
	curCordsSanta := cords{x: 0, y: 0}
	curCordsRobo := cords{x: 0, y: 0}

	santaCords := map[cords]bool{}
	roboCords := map[cords]bool{}
	curMovement := "santa"

	for i, v := range input {
		if i%2 == 0 {
			curMovement = "santa"
		} else {
			curMovement = "robot"
		}

		if string(v) == "v" || string(v) == "^" {
			if curMovement == "santa" {
				curCordsSanta.y += m[string(v)]
			} else {
				curCordsRobo.y += m[string(v)]
			}
		} else {
			if curMovement == "santa" {
				curCordsSanta.x += m[string(v)]
			} else {
				curCordsRobo.x += m[string(v)]
			}
		}

		if curMovement == "santa" {
			_, ok := santaCords[curCordsSanta]
			_, ok2 := roboCords[curCordsSanta]
			if !ok && !ok2 {
				santaCords[curCordsSanta] = true
			}
		} else {
			_, ok := roboCords[curCordsRobo]
			_, ok2 := santaCords[curCordsRobo]
			if !ok && !ok2 {
				roboCords[curCordsRobo] = true
			}
		}
	}
	fmt.Println(len(santaCords) + len(roboCords))
}
