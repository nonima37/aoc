package functions

import (
	"fmt"
	"io/ioutil"
)

func Add(a int, b int) int {
	return a + b
}

func Read() string {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Err")
	}
	return string(content)
}
