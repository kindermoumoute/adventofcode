package main

import (
	"strconv"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	ts, busList := parse(input)

	part1, part2 := 0, 0
	mod := 1
	minID := 0
	for i, b := range busList {
		freq := b.ID - ts%b.ID
		if freq < busList[minID].ID-ts%busList[minID].ID {
			part1 = b.ID * freq
			minID = i
		}

		for (part2+b.Offset)%b.ID != 0 {
			part2 += mod
		}
		mod *= b.ID
	}
	return part1, part2
}

type bus struct {
	ID     int
	Offset int

	frequencyPart1 int
}

func parse(s string) (int, []*bus) {
	lines := strings.Split(s, "\n")
	ts := pkg.MustAtoi(lines[0])
	ids := []*bus(nil)
	for i, budID := range strings.Split(lines[1], ",") {
		ID, err := strconv.Atoi(budID)
		if err != nil {
			continue
		}
		ids = append(ids, &bus{
			ID:     ID,
			Offset: i,
		})
	}
	return ts, ids
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
