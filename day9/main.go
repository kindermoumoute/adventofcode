package main

import "fmt"

func main() {
	stack := []uint8{}
	part1 := 0
	inc := 1
	part2 := 0
	for i := 0; i < len(puzzle); i++ {
		if len(stack) > 0 && stack[len(stack)-1] == '<' {
			part2++
		}
		switch puzzle[i] {
		case '!':
			i++
			part2--
		case '<':
			if len(stack) == 0 || stack[len(stack)-1] != '<' {
				stack = append(stack, '<')
			}
		case '{':
			if len(stack) == 0 || stack[len(stack)-1] != '<' {
				stack = append(stack, '{')
				part1 += inc
				inc++
			}
		case '}':
			if len(stack) > 0 && stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
				inc--
			}
		case '>':
			if len(stack) > 0 && stack[len(stack)-1] == '<' {
				stack = stack[:len(stack)-1]
				part2--
			}

		default:
		}
	}
	fmt.Println(part1, part2)
}
