package main

import (
	"container/list"
	"fmt"
	"sort"
	"strconv"

	"github.com/kindermoumoute/adventofcode/pkg"
)

// returns part1 and part2
func run(input string) (string, string) {
	players, lastMarble := parse(input)
	scores := make([]int, players)
	game := list.New()
	currentMarble := game.PushFront(0)
	currentPlayer := 0
	actualLastMarble := lastMarble * 100
	part1 := 0
	for i := 1; i <= actualLastMarble; i++ {
		if i%23 == 0 && i > 0 {
			scores[currentPlayer] += i
			for j := 0; j < 7; j++ {
				currentMarble = currentMarble.Prev()
				if currentMarble == nil {
					currentMarble = game.Back()
				}
			}
			toRemove := currentMarble
			scores[currentPlayer] += currentMarble.Value.(int)
			currentMarble = currentMarble.Next()
			if currentMarble == nil {
				currentMarble = game.Front()
			}
			game.Remove(toRemove)

		} else {
			currentMarble = currentMarble.Next()
			if currentMarble == nil {
				currentMarble = game.Front()
			}
			currentMarble = game.InsertAfter(i, currentMarble)
		}
		if i%lastMarble == 0 && i/lastMarble == 1 {
			tmpScores := make([]int, players)
			copy(tmpScores, scores)
			sort.Ints(tmpScores)
			part1 = tmpScores[len(tmpScores)-1]
		}
		currentPlayer = (currentPlayer + 1) % players

	}
	sort.Ints(scores)
	return strconv.Itoa(part1), strconv.Itoa(scores[len(scores)-1])
}

func parse(s string) (int, int) {
	var players, lastMarble int
	n, err := fmt.Sscanf(s, "%d players; last marble is worth %d points", &players, &lastMarble)
	pkg.Check(err)
	if n != 2 {
		panic("err")
	}
	return players, lastMarble
}

func main() {
	pkg.Execute(run, tests, puzzle, true)
}
