package main

import (
	"regexp"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	ranges, attributes, myTicket, tickets := parse(input)

	part1, part2 := 0, 1
	filteredTickets := [][]int(nil)
	for _, ticket := range tickets {
		discarded := false
	value:
		for _, value := range ticket {
			for _, r := range ranges {
				for _, limit := range r {
					if limit[0] <= value && value <= limit[1] {
						continue value
					}
				}
			}
			part1 += value
			discarded = true
		}
		if !discarded {
			filteredTickets = append(filteredTickets, ticket)
		}
	}

	// For each attribute, map all field index that could work
	multiIndexMapping := make(map[int]map[int]struct{})
	for i := 0; i < len(ranges); i++ {
		multiIndexMapping[i] = make(map[int]struct{})
	nextIndex:
		for j := 0; j < len(ranges); j++ {
			for _, ticket := range filteredTickets {
				if !(ranges[i][0][0] <= ticket[j] && ticket[j] <= ranges[i][0][1] ||
					ranges[i][1][0] <= ticket[j] && ticket[j] <= ranges[i][1][1]) {
					continue nextIndex
				}
			}
			multiIndexMapping[i][j] = struct{}{}
		}
	}

	// Only one index per attribute can work
	mapping := make(map[int]int)
	for k := 0; k < len(ranges); k++ {
		for i, indices := range multiIndexMapping {
			if len(indices) == 1 {
				for index := range indices {
					mapping[i] = index
					for j := range multiIndexMapping {
						if j != i {
							if _, exist := multiIndexMapping[j]; exist {
								delete(multiIndexMapping[j], index)
							}
						}
					}
				}
			}
		}
	}

	for i := range ranges {
		if !strings.Contains(attributes[i], "departure") {
			continue
		}
		part2 *= myTicket[mapping[i]]
	}

	return part1, part2
}

func parse(s string) (ranges [][2][2]int, attributes []string, yourTicket []int, tickets [][]int) {
	tmp := strings.SplitN(s, "\n\nyour ticket:\n", 2)
	regAttr := regexp.MustCompile(`(.*): (\d+)-(\d+) or (\d+)-(\d+)`)
	for _, match := range regAttr.FindAllStringSubmatch(tmp[0], -1) {
		attributes = append(attributes, match[1])
		ranges = append(ranges, [2][2]int{
			{pkg.MustAtoi(match[2]), pkg.MustAtoi(match[3])},
			{pkg.MustAtoi(match[4]), pkg.MustAtoi(match[5])},
		})
	}

	tmp = strings.SplitN(tmp[1], "\n\nnearby tickets:\n", 2)
	yourTicket = pkg.ParseIntList(tmp[0], ",")
	for _, line := range strings.Split(tmp[1], "\n") {
		tickets = append(tickets, pkg.ParseIntList(line, ","))
	}
	return
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
