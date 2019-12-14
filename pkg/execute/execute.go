package execute

import (
	"fmt"
	"time"

	"github.com/faiface/pixel/pixelgl"
	"github.com/kindermoumoute/adventofcode/pkg/twod"
)

func Run(run func(string) (interface{}, interface{}), tests TestCases, puzzle string, verbose bool) {
	if tests != nil {
		tests.Run(run, !verbose)
	}

	start := time.Now()
	part1, part2 := run(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("PART1: %v\nPART2: %v\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}

func RunWithPixel(run func(string) (interface{}, interface{}), tests TestCases, puzzle string, verbose bool) {
	twod.RenderingEnabled = true
	pixelgl.Run(func() {
		Run(run, tests, puzzle, verbose)
	})
}
