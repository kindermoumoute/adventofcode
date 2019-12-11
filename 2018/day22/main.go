package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
)

const (
	rocky = iota
	wet
	narrow
)

var tile = []string{".", "=", "|"}

type Region struct {
	GelogicalIndex int64
	Erosion        int64
}

func (r *Region) Type() int {
	return int(r.Erosion % 3)
}

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	depth, target := parse(input)
	cave := make(map[pkg.P]*Region)
	part1 := 0
	for j := 0; j <= target.Y; j++ {
		for i := 0; i <= target.X; i++ {
			gi := int64(0)
			if i == target.X && j == target.Y || i == 0 && j == 0 {

			} else if j == 0 {
				gi = (int64(i) * 16807) % 20183

			} else if i == 0 {
				gi = (int64(j) * 48271) % 20183
			} else {
				gi = ((cave[pkg.P{X: i - 1, Y: j}].GelogicalIndex + depth) * (cave[pkg.P{X: i, Y: j - 1}].GelogicalIndex + depth)) % 20183
			}
			p := pkg.P{X: i, Y: j}
			cave[p] = &Region{
				Erosion:        (gi + depth) % 20183,
				GelogicalIndex: gi,
			}
			fmt.Print(tile[cave[p].Type()])
			//fmt.Println(p, cave[p])
			part1 += cave[p].Type()
		}
		fmt.Println()
	}
	//intList := pkg.ParseIntList(input)
	// bufio.NewReader(os.Stdin).ReadBytes('\n')
	//strconv.Itoa()
	return strconv.Itoa(part1), ""
}

func parse(s string) (int64, pkg.P) {
	depth, target := 0, pkg.P{}
	lines := strings.Split(s, "\n")
	n, err := fmt.Sscanf(lines[0], "depth: %d", &depth)
	pkg.Check(err)
	if n != 1 {
		panic(fmt.Errorf("%d args expected in scanf, got %d", 0, n))
	}
	n, err = fmt.Sscanf(lines[1], "target: %d,%d", &target.X, &target.Y)
	pkg.Check(err)
	if n != 2 {
		panic(fmt.Errorf("%d args expected in scanf, got %d", 0, n))
	}
	return int64(depth), target
}

func main() {
	pkg.Execute(run, tests, puzzle, true)
}

// Get all prime factors of a given number n
func PrimeFactors(n int) (pfs []int) {
	// Get the number of 2s that divide n
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := 3; i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n > 2 {
		pfs = append(pfs, n)
	}

	return
}
