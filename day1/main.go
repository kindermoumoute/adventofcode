package main

import "fmt"

func main() {
	part1 := 0
	part2 := 0
	puzzleLenght := len(puzzle)
	for i, r := range puzzle {
		if puzzle[i] == puzzle[(i+1)%puzzleLenght] {
			part1 += int(r - '0')
		}
		if puzzle[i] == puzzle[(puzzleLenght/2+i)%puzzleLenght] {
			part2 += int(r - '0')
		}
	}
	fmt.Println(part1)
	fmt.Println(part2)
}
