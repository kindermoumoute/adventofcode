package main

type testCase struct {
	input, expectedPart1, expectedPart2 string
}

var tests = []testCase{
	testCase{
		`3`,
		`638`,
		`1226`,
	},
}

var puzzle = `349`
