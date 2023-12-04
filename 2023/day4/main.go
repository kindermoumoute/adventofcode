package main

import (
	"fmt"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	cards := parse(input)

	part1, part2 := 0, 0
	cardMultipliers := make(map[int]int)
	for _, card := range cards {
		cardMultipliers[card.ID]++
		winningNumbers := 0
		for n := range card.Hand {
			if card.Wins[n] {
				winningNumbers++
			}
		}
		part1Score := 0
		for i := 0; i < winningNumbers; i++ {
			cardMultipliers[card.ID+i+1] += cardMultipliers[card.ID]
			if part1Score == 0 {
				part1Score = 1
			} else {
				part1Score *= 2
			}
		}
		part1 += part1Score
		part2 += cardMultipliers[card.ID]
	}

	return part1, part2
}

type Card struct {
	ID   int
	Wins map[int]bool
	Hand map[int]bool
}

func parse(s string) []Card {
	cards := []Card(nil)
	lines := strings.Split(s, "\n")
	for id, line := range lines {
		c := Card{
			ID:   id + 1,
			Wins: make(map[int]bool),
			Hand: make(map[int]bool),
		}
		line = strings.Join(strings.Fields(line), " ")
		line, _ = strings.CutPrefix(line, fmt.Sprintf("Card %d: ", c.ID))
		tmp := strings.Split(line, " | ")
		for _, n := range pkg.ParseIntList(tmp[0], " ") {
			if c.Wins[n] {
				panic("change strat wins")
			}
			c.Wins[n] = true
		}
		for _, n := range pkg.ParseIntList(tmp[1], " ") {
			if c.Hand[n] {
				panic("change strat wins 2")
			}
			c.Hand[n] = true
		}

		cards = append(cards, c)
	}
	return cards
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
