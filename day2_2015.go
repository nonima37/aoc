package main

import (
	"aoc/functions"
	"fmt"
	"strconv"
	"strings"
)

func stringToInt(inputString string) []int {
	inputStringArr := strings.Split(inputString, "x")
	var inputInt []int
	for _, i := range inputStringArr {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		inputInt = append(inputInt, j)
	}
	return inputInt
}

func minSides(arr []int) []int {
	min := arr[0]
	min2 := arr[1]

	if min2 < min {
		min2 = min
		min = arr[1]
	}

	if arr[2] < min {
		min2 = min
		min = arr[2]
	} else if arr[2] < min2 {
		min2 = arr[2]
	}
	res := []int{min, min2}
	return res
}

func day2Main() {
	inputString := strings.Split(functions.Read(), "\r\n")
	var res1 int = 0
	var res2 int = 0

	var curDims []int
	for i := 0; i < len(inputString)-1; i++ {
		curDims = stringToInt(inputString[i])
		curMinSides := minSides(curDims)
		res1 += 2*curDims[0]*curDims[1] + 2*curDims[1]*curDims[2] + 2*curDims[0]*curDims[2] + curMinSides[0]*curMinSides[1]
		res2 += curMinSides[0]*2 + curMinSides[1]*2 + curDims[0]*curDims[1]*curDims[2]
	}
	fmt.Println(res2)
}
