package main

import (
	"strconv"

	"github.com/kindermoumoute/adventofcode/pkg"
)

// returns part1 and part2
func run(input string) (string, string) {
	// part1
	part1 := Polymer(input).react()

	// part2
	polymers, minLetter, min := make(map[byte]Polymer), byte(0), 9999999999999
	for letter := byte('a'); letter < 'z'; letter++ {
		polymers[letter] = removeLetterFromPolymer(Polymer(input), letter).react()
		if len(polymers[letter]) < min {
			min = len(polymers[letter])
			minLetter = letter
		}
	}
	return strconv.Itoa(len(part1)), strconv.Itoa(len(polymers[minLetter]))
}

type Polymer []byte

func (p Polymer) react() Polymer {
	for i := 0; i < len(p)-1; i++ {
		if p[i+1]^p[i] == 32 {
			p = append(p[:i], p[i+2:]...)
			if i -= 2; i < -1 {
				i = -1
			}
		}
	}
	return p
}

func removeLetterFromPolymer(p Polymer, letter byte) Polymer {
	for i := 0; i < len(p); i++ {
		if letter == p[i] || letter-32 == p[i] {
			p = append(p[:i], p[i+1:]...)
			i--
		}
	}
	return p
}

func main() {
	pkg.Execute(run, tests, puzzle, false)
}
