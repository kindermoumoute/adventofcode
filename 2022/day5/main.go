package main

import (
	"container/list"
	"fmt"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	stacks, actions := parse(input)
	stacksPart2, _ := parse(input)

	for _, action := range actions {
		tmp := list.New()
		for i := 0; i < action.Amount; i++ {
			// Part 2
			tmp.PushBack(stacksPart2[action.From].Remove(stacksPart2[action.From].Front()))

			// Part 1
			stacks[action.To].PushFront(stacks[action.From].Remove(stacks[action.From].Front()))
		}
		stacksPart2[action.To].PushFrontList(tmp)

		//for i := 1; i <= len(stacks); i++ {
		//	Print(stacks[i])
		//}
	}

	word := []int32(nil)
	for i := 1; i <= len(stacks); i++ {
		word = append(word, stacks[i].Front().Value.(int32))
	}
	wordPart2 := []int32(nil)
	for i := 1; i <= len(stacksPart2); i++ {
		wordPart2 = append(wordPart2, stacksPart2[i].Front().Value.(int32))
	}

	return string(word), string(wordPart2)
}

type Action struct {
	Amount, From, To int
}

func Print(l *list.List) {
	v := l.Front()
	for i := 0; i < l.Len(); i++ {
		fmt.Print(string(v.Value.(int32)))
		v = v.Next()
	}
	fmt.Println()
}

func parse(s string) (map[int]*list.List, []Action) {
	lines := strings.Split(s, "\n")
	parseStacks := true
	stacks := make(map[int]*list.List)
	actions := []Action(nil)
	for i := 0; i < len(lines); i++ {
		if lines[i][1] == '1' {
			i += 2
			parseStacks = false
		}
		line := lines[i]
		if parseStacks {
			for j, c := range line {
				if c == ' ' || c == '[' || c == ']' {
					continue
				}
				stackIndex := j/4 + 1
				stack, exist := stacks[stackIndex]
				if !exist {
					stack = list.New()
					stacks[stackIndex] = stack
				}

				stack.PushBack(c)
			}
		} else {
			action := Action{}
			pkg.MustScanf(line, "move %d from %d to %d", &action.Amount, &action.From, &action.To)
			actions = append(actions, action)
		}
	}
	return stacks, actions
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
