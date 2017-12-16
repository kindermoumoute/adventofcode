package main

import (
	"fmt"
	"strconv"
	"strings"
)

const input = "abcdefghijklmnop"

func main() {
	fmt.Println(parse(puzzle, []byte(input)))
}

func parse(s string, programs []byte) (string, string) {
	part1 := ""
	lines := strings.Split(s, ",")
	programSlice := make([]string, 0, len(lines))
	programMap := make(map[string]int)

	stop := -1
	for i := 0; ; i++ {
		_, exist := programMap[string(programs)]
		if exist {
			stop = 1000000000 % i
			break
		}
		programSlice = append(programSlice, string(programs))
		programMap[string(programs)] = i
		for _, line := range lines {
			switch line[0] {
			case 's':
				a, _ := strconv.Atoi(line[1:])
				programs = spin(programs, a)
			case 'x':
				values := strings.Split(line[1:], "/")
				a, _ := strconv.Atoi(values[0])
				b, _ := strconv.Atoi(values[1])
				exchange(programs, a, b)
			case 'p':
				partner(programs, byte(line[1]), byte(line[3]))
			}

			if i == 0 {
				part1 = string(programs)
			}
		}
	}
	return part1, programSlice[stop]
}

func spin(s []byte, n int) []byte {
	t := make([]byte, len(s))
	for i := 0; i < 16; i++ {
		t[(i+n)%16] = s[i]
	}
	return t
}

func exchange(s []byte, a, b int) {
	s[a], s[b] = s[b], s[a]
}

func partner(s []byte, aName, bName byte) {
	a, b := -1, -1
	for i := range s {
		if s[i] == aName {
			a = i
		}
		if s[i] == bName {
			b = i
		}
	}
	exchange(s, a, b)
}
