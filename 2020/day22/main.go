package main

import (
	"container/list"
	"strconv"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	decks, winner := play(parse(input), false)
	part1 := winningScore(decks[winner])

	decks, winner = play(parse(input), true)
	part2 := winningScore(decks[winner])

	return part1, part2
}

func play(decks []*list.List, part2 bool) ([]*list.List, int) {
	winner := 0
	dejaVu := make(map[string]struct{})
	for decks[0].Len() != 0 && decks[1].Len() != 0 {
		if part2 {
			decksHash := hash(decks)
			_, exist := dejaVu[decksHash]
			if exist {
				return decks, 0
			}
			dejaVu[decksHash] = struct{}{}
		}

		cards := []int{
			decks[0].Remove(decks[0].Front()).(int),
			decks[1].Remove(decks[1].Front()).(int),
		}
		switch {
		case cards[0] <= decks[0].Len() && cards[1] <= decks[1].Len() && part2:
			_, winner = play([]*list.List{copyList(decks[0], cards[0]), copyList(decks[1], cards[1])}, part2)
		case cards[1] > cards[0]:
			winner = 1
		default:
			winner = 0
		}
		decks[winner].PushBack(cards[winner])
		decks[winner].PushBack(cards[1-winner])
	}
	return decks, winner
}

func hash(decks []*list.List) string {
	strs := []string(nil)
	for _, l := range decks {
		elements := []string(nil)
		for e := l.Front(); e != nil; e = e.Next() {
			elements = append(elements, strconv.Itoa(e.Value.(int)))
		}
		strs = append(strs, strings.Join(elements, ","))
	}
	return strings.Join(strs, "\n")
}

func winningScore(l *list.List) int {
	sum := 0
	i := 1
	for e := l.Back(); e != nil; e = e.Prev() {
		sum += e.Value.(int) * i
		i++
	}
	return sum
}

func copyList(l *list.List, n int) *list.List {
	newList := list.New()
	i := 0
	for e := l.Front(); e != nil && i < n; e = e.Next() {
		newList.PushBack(e.Value)
		i++
	}
	return newList
}

func parse(s string) []*list.List {
	lists := make([]*list.List, 2)
	for i, player := range strings.Split(s, "\n\n") {
		lines := strings.Split(player, "\n")
		lists[i] = list.New()
		for _, line := range lines[1:] {
			lists[i].PushBack(pkg.MustAtoi(line))
		}
	}
	return lists
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
