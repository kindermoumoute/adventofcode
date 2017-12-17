package main

type testCase struct {
	input, expectedPart1, expectedPart2 string
}

var tests = []testCase{
	testCase{
		`3`,
		`638`,
		`1222153`,
	},
}

var puzzle = `349`
