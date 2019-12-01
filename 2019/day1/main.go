package main

import (
	"strconv"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
)

// returns part1 and part2
func run(input string) (string, string) {
	part1, part2 := parse(input)
	//intList := pkg.ParseIntList(input)
	// bufio.NewReader(os.Stdin).ReadBytes('\n')
	//strconv.Itoa()
	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func parse(s string) (int, int) {
	lines := strings.Split(s, "\n")
	part1 := 0
	part2 := 0
	for _, line := range lines {
		i, err := strconv.Atoi(line)
		pkg.Check(err)
		part1 += i/3 - 2
		part2 += sumRec(i)
	}
	return part1, part2
}
func sumRec(i int) int {
	i = i/3 - 2
	if i <= 0 {
		return 0
	}
	return (i) + sumRec(i)
}

func main() {
	pkg.Execute(run, nil, puzzle, true)
}
