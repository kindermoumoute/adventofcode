package main

import (
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	ruckstacks := parse(input)

	part1, part2 := 0, 0
	groupItemsMasks := make(map[int]int)
	for elf, ruckstack := range ruckstacks {
		groupMember := elf % 3
		itemsMasks := make(map[int]int)
		for i := 0; i < len(ruckstack.Left); i++ {
			itemsMasks[ruckstack.Left[i]] |= 1
			itemsMasks[ruckstack.Right[i]] |= 2

			groupItemsMasks[ruckstack.Left[i]] |= 1 << groupMember
			groupItemsMasks[ruckstack.Right[i]] |= 1 << groupMember
		}
		for priority, mask := range itemsMasks {
			if mask == 3 {
				part1 += priority
			}
		}

		if groupMember == 2 {
			for priority, mask := range groupItemsMasks {
				if mask == 7 {
					part2 += priority
				}
			}
			groupItemsMasks = make(map[int]int)
		}
	}

	return part1, part2
}

type Rucksack struct {
	Left  []int
	Right []int
}

func parse(s string) []Rucksack {
	rucksacks := []Rucksack(nil)
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		rucksack := Rucksack{}
		for i, letter := range line {
			var priority int
			switch {
			case letter >= 'A' && letter <= 'Z':
				priority = int(letter - 'A' + 27)
			case letter >= 'a' && letter <= 'z':
				priority = int(letter - 'a' + 1)
			default:
				panic("unexpected")
			}

			if i < len(line)/2 {
				rucksack.Left = append(rucksack.Left, priority)
			} else {
				rucksack.Right = append(rucksack.Right, priority)
			}
		}
		rucksacks = append(rucksacks, rucksack)
	}
	return rucksacks
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
