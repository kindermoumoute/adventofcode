package main

import (
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
	"github.com/kindermoumoute/adventofcode/pkg/font"
	"github.com/kindermoumoute/adventofcode/pkg/twod"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	actions := parse(input)

	part1 := 0
	actionIndex := 0
	x := 1
	allCycles := make(map[int]int)
	draw := make(twod.Map)
	for i := 1; i <= 240; i++ {
		drawPoint(i, x, draw)

		allCycles[i] = x * i
		action := actions[actionIndex%len(actions)]
		if action.Name == "addx" {
			i++
			drawPoint(i, x, draw)
			allCycles[i] = x * i
			x += action.Value
		}
		actionIndex++
	}

	part1 = allCycles[20] + allCycles[60] + allCycles[100] + allCycles[140] + allCycles[180] + allCycles[220]
	part2, _ := font.FindWordInMap(draw)

	return part1, part2
}

func drawPoint(i, x int, draw twod.Map) {

	p := twod.NewVector((i-1)%40, (i-1)/40)
	if p.X() >= x-1 && p.X() <= x+1 {
		draw[p] = 'X'
	}
}

type Action struct {
	Name  string
	Value int
}

func parse(s string) []Action {
	lines := strings.Split(s, "\n")
	actions := []Action(nil)
	for _, line := range lines {
		lineDetails := strings.Split(line, " ")
		a := Action{
			Name: lineDetails[0],
		}
		switch lineDetails[0] {
		case "addx":
			a.Value = pkg.MustAtoi(lineDetails[1])
		case "noop":
		}
		actions = append(actions, a)
	}
	return actions
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
