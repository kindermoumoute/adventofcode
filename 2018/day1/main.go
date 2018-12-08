package main

import (
	"fmt"
	"strconv"

	"github.com/kindermoumoute/adventofcode/pkg"
)

// returns part1 and part2
func run(input string) (string, string) {
	// part1
	list := pkg.ParseIntList(input, "\n")
	part1 := strconv.Itoa(pkg.Sum(list...))

	// part2
	frequencyMap := map[int]struct{}{0: {}}
	sum := 0
	for {
		for _, number := range list {
			sum += number
			_, exist := frequencyMap[sum]
			if exist {
				return part1, strconv.Itoa(sum)
			}
			frequencyMap[sum] = struct{}{}
		}
	}
}

func main() {
	tests.Run(run, false)
	fmt.Println(run(puzzle))
}
