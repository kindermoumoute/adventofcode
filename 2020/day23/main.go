package main

import (
	"container/ring"
	"fmt"

	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {

	part1 := ""

	// part 1
	current, index := parse(input)
	initLen := 9
	for i := 0; i < 100; i++ {
		current = moveCups(current, index, initLen)
	}
	index[1].Next().Do(func(i interface{}) {
		if i == 1 {
			return
		}
		part1 += fmt.Sprintf("%d", i)
	})

	// part 2
	current, index = parse(input)
	additionalRings := ring.New(1000000 - 9)
	for i := 10; i <= 1000000; i++ {
		additionalRings.Value = i
		index[i] = additionalRings
		additionalRings = additionalRings.Next()
	}
	current.Prev().Link(additionalRings)
	for i := 0; i < 10000000; i++ {
		current = moveCups(current, index, 1000000)
	}
	part2 := index[1].Next().Value.(int) * index[1].Next().Next().Value.(int)

	return part1, part2
}

func moveCups(current *ring.Ring, index map[int]*ring.Ring, ringLen int) *ring.Ring {
	excluded := map[interface{}]struct{}{current.Value: {}}
	tmp := current.Unlink(3)
	tmp.Do(func(i interface{}) {
		excluded[i] = struct{}{}
	})

	nextValue := current.Value.(int)
	for _, exist := excluded[nextValue]; exist; _, exist = excluded[nextValue] {
		if nextValue--; nextValue < 1 {
			nextValue = ringLen
		}
	}
	index[nextValue].Link(tmp)

	return current.Next()
}

func parse(s string) (*ring.Ring, map[int]*ring.Ring) {
	r := ring.New(len(s))
	index := make(map[int]*ring.Ring)
	for _, cup := range s {
		r.Value = int(cup - '0')
		index[r.Value.(int)] = r
		r = r.Next()
	}
	return r, index
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
