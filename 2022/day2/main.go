package main

import (
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

var shapeCounter = map[string]string{
	"A": "B",
	"B": "C",
	"C": "A",
}

var shapeLoose = map[string]string{
	"A": "C",
	"B": "A",
	"C": "B",
}

var shapeScore = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
}

var part1ToShape = map[string]string{
	"X": "B",
	"Y": "C",
	"Z": "A",
}

var resultScore = map[string]int{
	"AA": 3,
	"BA": 0,
	"CA": 6,
	"AB": 6,
	"BB": 3,
	"CB": 0,
	"AC": 0,
	"BC": 6,
	"CC": 3,
}

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	matches := parse(input)

	part1, part2 := 0, 0
	for _, match := range matches {
		part1Shape := part1ToShape[match.Result]
		part1 += shapeScore[part1Shape] + resultScore[match.Opponent+part1Shape]

		var part2Shape string
		switch match.Result {
		case "X":
			part2Shape = shapeLoose[match.Opponent]
		case "Y":
			part2Shape = match.Opponent
		case "Z":
			part2Shape = shapeCounter[match.Opponent]
		}
		part2 += shapeScore[part2Shape] + resultScore[match.Opponent+part2Shape]
	}

	return part1, part2
}

type Match struct {
	Opponent string
	Result   string
}

func parse(s string) []Match {
	lines := strings.Split(s, "\n")
	allMatches := []Match(nil)
	for _, line := range lines {
		m := Match{}
		pkg.MustScanf(line, "%s %s", &m.Opponent, &m.Result)
		allMatches = append(allMatches, m)
	}
	return allMatches
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
