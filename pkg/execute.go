package pkg

import (
	"fmt"
	"time"
)

func Execute(run func(string) (string, string), tests TestCases, puzzle string, verbose bool) {
	if tests != nil {
		tests.Run(run, !verbose)
	}

	start := time.Now()
	part1, part2 := run(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("PART1: %s\nPART2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
