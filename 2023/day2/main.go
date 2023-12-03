package main

import (
	"fmt"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	games := parse(input)

	part1, part2 := 0, 0

	for _, game := range games {
		maxRed, maxGreen, maxBlue := 0, 0, 0
		part1Possible := true
		for _, round := range game.Rounds {
			if round.Red > 12 || round.Green > 13 || round.Blue > 14 {
				part1Possible = false
			}
			if round.Red > maxRed {
				maxRed = round.Red
			}
			if round.Green > maxGreen {
				maxGreen = round.Green
			}
			if round.Blue > maxBlue {
				maxBlue = round.Blue
			}
		}
		if part1Possible {
			part1 += game.ID
		}
		part2 += maxGreen * maxBlue * maxRed
	}

	return part1, part2
}

type Game struct {
	ID     int
	Rounds []Round
}

type Round struct {
	Green, Blue, Red int
}

func parse(s string) []Game {
	games := []Game(nil)
	lines := strings.Split(s, "\n")
	for i, line := range lines {

		gameID := i + 1
		game := Game{ID: gameID}
		rightPart, _ := strings.CutPrefix(line, fmt.Sprintf("Game %d: ", gameID))
		for _, round := range strings.Split(rightPart, "; ") {
			r := Round{}
			for _, colorAction := range strings.Split(round, ", ") {
				tmp := strings.Split(colorAction, " ")
				switch tmp[1] {
				case "red":
					r.Red = pkg.MustAtoi(tmp[0])
				case "green":
					r.Green = pkg.MustAtoi(tmp[0])
				case "blue":
					r.Blue = pkg.MustAtoi(tmp[0])
				default:
					panic(tmp[1])
				}
			}
			game.Rounds = append(game.Rounds, r)
		}

		games = append(games, game)
	}
	return games
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
