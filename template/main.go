package main

import (
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg/execute"

	"github.com/kindermoumoute/adventofcode/pkg"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	//list := parse(input)
	//intList := pkg.ParseIntList(input)
	// bufio.NewReader(os.Stdin).ReadBytes('\n')
	//strconv.Itoa()
	return "", ""
}

func parse(s string) {
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		pkg.MustScanf(line, "")
	}
	return
}

func main() {
	execute.Run(run, nil, puzzle, true)
}
