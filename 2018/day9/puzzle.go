package main

import "github.com/kindermoumoute/adventofcode/pkg"

var tests = pkg.TestCases{
	{
		`10 players; last marble is worth 25 points`,
		`32`,
		`19708`,
	},
	{
		`10 players; last marble is worth 1618 points`,
		`8317`,
		`74765078`,
	},
	{
		`13 players; last marble is worth 7999 points`,
		`146373`,
		`1406506154`,
	},
	{
		`17 players; last marble is worth 1104 points`,
		`2764`,
		`20548882`,
	},
	{
		puzzle,
		`422748`,
		`3412522480`,
	},
}

var puzzle = `430 players; last marble is worth 71588 points`
