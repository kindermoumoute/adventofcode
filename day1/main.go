package main

import "fmt"

func main() {
	part1 := 0
	part2 := 0
	puzzleLength := len(puzzle)
	for i, r := range puzzle {
		if puzzle[i] == puzzle[(i+1)%puzzleLength] {
			part1 += int(r - '0')
		}
		if puzzle[i] == puzzle[(puzzleLenght/2+i)%puzzleLength] {
			part2 += int(r - '0')
		}
	}
	fmt.Println(part1)
	fmt.Println(part2)
}
