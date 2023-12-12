package main

import (
	"bytes"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	arrangements := parse(input)

	part1, part2 := 0, 0

	for _, a := range arrangements {
		part1 += a.Count()
		duplicate := [][]byte(nil)
		checksums := []int(nil)
		for i := 0; i < 5; i++ {
			duplicate = append(duplicate, a.seq)
			checksums = append(checksums, a.checkSum...)
		}
		a.seq = bytes.Join(duplicate, []byte{'?'})
		a.checkSum = checksums

		part2 += a.Count()
	}

	return part1, part2
}

type Arrangment struct {
	seq      []byte
	checkSum []int
}

func (a Arrangment) Count() int {
	return Arrangment{seq: append(a.seq, '.'), checkSum: a.checkSum}.count(0, 0, 0, make(map[[3]int]int))
}

func (a Arrangment) count(i, sum, groupID int, m map[[3]int]int) (c int) {
	key := [3]int{i, sum, groupID}
	v, exist := m[key]
	if exist {
		return v
	}

	defer func() {
		if !exist {
			m[key] = c
		}
	}()

	switch {
	case i == len(a.seq) && groupID == len(a.checkSum):
		return 1
	case i == len(a.seq) && groupID != len(a.checkSum):
		return 0
	case a.seq[i] == '#':
		return a.count(i+1, sum+1, groupID, m)
	case a.seq[i] == '.' || groupID == len(a.checkSum):
		if groupID < len(a.checkSum) && sum == a.checkSum[groupID] {
			return a.count(i+1, 0, groupID+1, m)
		} else if sum == 0 {
			return a.count(i+1, 0, groupID, m)
		} else {
			return 0
		}
	default:
		damagedCount := a.count(i+1, sum+1, groupID, m)
		operationalCount := 0
		if sum == a.checkSum[groupID] {
			operationalCount = a.count(i+1, 0, groupID+1, m)
		} else if sum == 0 {
			operationalCount = a.count(i+1, 0, groupID, m)
		}

		return damagedCount + operationalCount
	}
}

func parse(s string) []Arrangment {
	lines := strings.Split(s, "\n")
	ret := []Arrangment(nil)
	for _, line := range lines {
		tmp := strings.Split(line, " ")
		ret = append(ret, Arrangment{
			seq:      []byte(tmp[0]),
			checkSum: pkg.ParseIntList(tmp[1], ","),
		})
	}
	return ret
}

func main() {
	execute.Run(run, nil, puzzle, true)
}
