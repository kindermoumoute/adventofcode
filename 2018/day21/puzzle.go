package main

import "github.com/kindermoumoute/adventofcode/pkg"

var tests = pkg.TestCases{
	{
		puzzle,
		`11285115`,
		`2947113`,
	},
}

var puzzle = `#ip 2
seti 123 0 1
bani 1 456 1
eqri 1 72 1
addr 1 2 2
seti 0 0 2
seti 0 4 1
bori 1 65536 3
seti 10905776 4 1
bani 3 255 4
addr 1 4 1
bani 1 16777215 1
muli 1 65899 1
bani 1 16777215 1
gtir 256 3 4
addr 4 2 2
addi 2 1 2
seti 27 1 2
seti 0 6 4
addi 4 1 5
muli 5 256 5
gtrr 5 3 5
addr 5 2 2
addi 2 1 2
seti 25 1 2
addi 4 1 4
seti 17 9 2
setr 4 7 3
seti 7 4 2
eqrr 1 0 4
addr 4 2 2
seti 5 1 2`
