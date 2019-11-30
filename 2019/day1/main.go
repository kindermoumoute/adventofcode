package main

import (
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
)

// returns part1 and part2
func run(input string) (string, string) {
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
	pkg.Execute(run, nil, puzzle, true)
}
