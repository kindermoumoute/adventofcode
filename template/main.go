package main

import (
	"fmt"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
)

// returns part1 and part2
func run(input string) (string, string) {
	//list := parse(input)
	//intList := pkg.ParseIntList(input)

	//strconv.Itoa()
	return "", ""
}

func parse(s string) {
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		n, err := fmt.Sscanf(line, "")
		pkg.Check(err)
		if n != 0 {
			panic(fmt.Errorf("%d args expected in scanf, got %d", 0, n))
		}
	}
	return
}

func main() {
	pkg.Execute(run, nil, puzzle, true)
}
