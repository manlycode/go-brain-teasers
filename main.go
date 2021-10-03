package main

import (
	"fmt"
)

func main() {
	puzzles := map[string]func(){
		"puzzle1": func() {
			var π = 22 / 7.0
			fmt.Println(π)
		},

		"puzzle2": func() {
		},
	}

	for key, element := range puzzles {
		fmt.Println("Puzzle: ", key)
		fmt.Println("------------------------------")
		element()
		fmt.Println("")
	}
}
