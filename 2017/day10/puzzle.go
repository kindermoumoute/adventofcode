package main

type testCase struct {
	input, expectedPart1, expectedPart2 string
}

var tests = []testCase{
	testCase{
		``,
		`0`,
		`a2582a3a0e66e6e86e3812dcb672a272`,
	},
	testCase{
		`AoC 2017`,
		`0`,
		`33efeb34ea91902bb2f59c9920caa6cd`,
	},
	testCase{
		`1,2,3`,
		`0`,
		`3efbe78a8d82f29979031a4aa0b16a9d`,
	},
	testCase{
		`1,2,4`,
		`0`,
		`63960835bcdc130f0b66d7ff4f6a5a8e`,
	},
}

var puzzle = `197,97,204,108,1,29,5,71,0,50,2,255,248,78,254,63`
