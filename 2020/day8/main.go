package main

import (
	"github.com/kindermoumoute/adventofcode/pkg/execute"
	"github.com/kindermoumoute/adventofcode/pkg/protocode"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	code := protocode.New(input)
	code.ExitCondition = func(c *protocode.ProtoCode) bool {
		return c.Memory[c.Cursor].Used >= 1
	}

	part1 := code.Run()
	part2 := 0
	for j := 0; j < len(code.Memory); j++ {
		if code.Memory[j].Opcode == protocode.JMP {
			code.Memory[j].Opcode = protocode.NOP
		} else if code.Memory[j].Opcode == protocode.NOP {
			code.Memory[j].Opcode = protocode.JMP
		} else {
			continue
		}
		code.Reset()
		part2 = code.Run()
		if code.Cursor == len(code.Memory) {
			break
		}
		if code.Memory[j].Opcode == protocode.NOP {
			code.Memory[j].Opcode = protocode.JMP
		} else if code.Memory[j].Opcode == protocode.JMP {
			code.Memory[j].Opcode = protocode.NOP
		}
	}

	return part1, part2
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
