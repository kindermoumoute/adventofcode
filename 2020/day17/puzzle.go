package main

import (
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

var tests = execute.TestCases{
	{
		`.#.
..#
###`,
		112,
		848,
	},
	{
		puzzle,
		223,
		1884,
	},
}

var puzzle = `...#...#
..##.#.#
###..#..
........
...##.#.
.#.####.
...####.
..##...#`
