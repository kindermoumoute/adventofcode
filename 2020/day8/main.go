package main

import (
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	m := parse(input)
	part1 := 0
	for i := 0; m[i].usage < 1; i++ {
		switch m[i].verb {
		case "jmp":
			i += m[i].value - 1
		case "acc":
			part1 += m[i].value
		case "nop":
		}
		m[i].usage++
	}
	part2 := 0

	for j := 0; j < len(m); j++ {
		m = parse(input)
		if m[j].verb == "jmp" {
			m[j].verb = "nop"
		} else if m[j].verb == "nop" {
			m[j].verb = "jmp"
		} else {
			continue
		}
		part2 = 0
		i := 0
		for i = 0; i < len(m) && m[i].usage < 1; i++ {
			switch m[i].verb {
			case "jmp":
				i += m[i].value - 1
			case "acc":
				part2 += m[i].value
			case "nop":
			}

			_, exist := m[i]
			if !exist {
				break
			}
			m[i].usage++
		}
		if i == len(m) {
			break
		}
	}

	return part1, part2
}

type ins struct {
	verb     string
	value    int
	usage    int
	nopCount int
	jmpCount int
}

func parse(s string) map[int]*ins {
	lines := strings.Split(s, "\n")
	m := make(map[int]*ins)
	for i, line := range lines {
		line = strings.ReplaceAll(line, "+", "")
		verb, value := "", 0
		pkg.MustScanf(line, "%s %d", &verb, &value)
		m[i] = &ins{
			verb:  verb,
			value: value,
		}
	}
	return m
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
