package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
)

const toUp = 'a' - 'A'
const toDown = 'A' - 'a'

// returns part1 and part2
func run(input string) (string, string) {
	part1 := []byte(input)
	//list := parse(input)
	//intList := pkg.ParseIntList(input)
	iterator := 0
	for {
		if iterator+1 > len(part1)-1 {
			break
		}
		if iterator < 0 {
			iterator++
		}
		diff := part1[iterator+1] - part1[iterator]
		if diff == 32 || diff == 224 {
			part1 = append(part1[:iterator], part1[iterator+2:]...)
			iterator -= 2
		}
		iterator++
	}
	all := make(map[byte][]byte)
	minI := byte(0)
	min := 9999999999999
	for i := byte('a'); i < 'z'; i++ {
		tmp := string([]byte{i})
		tmpUP := string([]byte{i - 32})
		fmt.Println("LETTERS ", tmp, tmpUP)
		all[i] = []byte(strings.Replace(strings.Replace(input, tmpUP, "", -1), tmp, "", -1))
		fmt.Println(string(all[i]))
		iterator := 0
		for {
			if iterator+1 > len(all[i])-1 {
				break
			}
			if iterator < 0 {
				iterator++
			}
			diff := all[i][iterator+1] - all[i][iterator]
			if diff == 32 || diff == 224 {
				all[i] = append(all[i][:iterator], all[i][iterator+2:]...)
				iterator -= 2
			}
			iterator++
		}

		fmt.Println(string(all[i]))
		if len(all[i]) < min {
			min = len(all[i])
			minI = i
		}
	}
	return strconv.Itoa(len(part1)), strconv.Itoa(len(all[minI]))
}

func parse(s string) {
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		n, err := fmt.Sscanf(line, "")
		pkg.Check(err)
		if n != 0 {
			panic(fmt.Errorf("%d args expected in scanf, got %d", 0, n))
		}
	}
	return
}

func main() {
	tests.Run(run, false)
	fmt.Println(run(puzzle))
}
