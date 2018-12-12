package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
)

// returns part1 and part2
func run(input string) (string, string) {
	day := parse(input)
	//intList := pkg.ParseIntList(input)
	offset := 5
	offString := ""
	for i := 0; i < offset; i++ {
		offString += "."
	}
	currentState := offString + day.initial + offString
	nextState := []byte(currentState)
	for g := 1; g <= 20; g++ {
		for _, state := range day.states {
			currentIndex := strings.Index(currentState, state.chunk)
			indexSum := 0
			tmp := currentState
			for currentIndex >= 0 {
				indexSum += currentIndex + 2
				nextState[indexSum] = state.out
				fmt.Println(currentState, "MATCHES", state.chunk, "at", indexSum, "set", state.out)
				tmp = tmp[currentIndex+1:]
				currentIndex = strings.Index(tmp, state.chunk)
			}
		}
		currentState = string(nextState)

		fmt.Println(g, currentState)
	}
	part1 := 0
	for i, v := range nextState {
		switch v {
		case '#':
			part1 += i - offset
		}
	}

	return strconv.Itoa(part1), ""
}

type state struct {
	chunk string
	out   byte
}

type day12 struct {
	initial string
	states  []state
}

func parse(s string) day12 {
	lines := strings.Split(s, "\n")
	newDay := day12{}
	newDay.states = make([]state, len(lines)-2)
	n, err := fmt.Sscanf(lines[0], "initial state: %s", &newDay.initial)
	pkg.Check(err)
	if n != 1 {
		panic(fmt.Errorf("%d args expected in scanf, got %d", 0, n))
	}
	for i, line := range lines[2:] {
		n, err := fmt.Sscanf(line, "%s => %c", &newDay.states[i].chunk, &newDay.states[i].out)
		pkg.Check(err)
		if n != 2 {
			panic(fmt.Errorf("%d args expected in scanf, got %d", 0, n))
		}
	}
	return newDay
}

func main() {
	pkg.Execute(run, tests, puzzle, false)
}
