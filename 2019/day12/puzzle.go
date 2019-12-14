package main

import (
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

var tests = execute.TestCases{
	{
		`<x=-1, y=0, z=2>
<x=2, y=-10, z=-7>
<x=4, y=-8, z=8>
<x=3, y=5, z=-1>`,
		``,
		2772,
	},
	{
		puzzle,
		11384,
		452582583272768,
	},
}

var puzzle = `<x=-7, y=-1, z=6>
<x=6, y=-9, z=-9>
<x=-12, y=2, z=-7>
<x=4, y=-17, z=-12>`
