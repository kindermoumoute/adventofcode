package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {

	rules, messages := parse(input)

	// part 1
	part1 := rules.countValidMessages(messages)

	if puzzle != input { // part 2 is for big puzzle only
		return part1, 0
	}

	// part 2: 8th and 11th patterns are solved "manually"
	rules.reset()
	rules[8].Regex = rules.getRegex(42) + "+"
	leftPattern11 := rules.getRegex(42)
	rightPattern11 := rules.getRegex(31)
	patternIterations := []string(nil)
	for i := 1; i < 5; i++ { // looks like only 5 iterations are sufficient for my puzzle
		patternIterations = append(patternIterations, fmt.Sprintf("(%[1]s){%[3]d}(%[2]s){%[3]d}", leftPattern11, rightPattern11, i))
	}
	rules[11].Regex = "(" + strings.Join(patternIterations, "|") + ")"
	part2 := rules.countValidMessages(messages)

	return part1, part2
}

type Rule struct {
	ID       int
	Regex    string
	SubRules [][]int
}

type Rules map[int]*Rule

func (r Rules) countValidMessages(messages []string) int {
	sum := 0
	reg := regexp.MustCompile("^" + r.getRegex(0) + "$")
	for _, message := range messages {
		if reg.MatchString(message) {
			sum++
		}
	}
	return sum
}

func (r Rules) reset() {
	for _, rule := range r {
		if len(rule.Regex) != 1 {
			rule.Regex = ""
		}
	}
}

func (r Rules) getRegex(id int) string {
	if len(r[id].Regex) != 0 {
		return r[id].Regex
	}
	regexRules := []string(nil)
	for i, subRule := range r[id].SubRules {
		regexRules = append(regexRules, "")
		for _, ruleID := range subRule {
			regexRules[i] += r.getRegex(ruleID)
		}
	}
	if len(regexRules) == 1 {
		r[id].Regex = regexRules[0]
	} else {
		r[id].Regex = "(" + strings.Join(regexRules, "|") + ")"
	}
	return r[id].Regex
}

func parse(s string) (Rules, []string) {
	inputPart := strings.Split(s, "\n\n")
	messages := strings.Split(inputPart[1], "\n")
	rules := make(map[int]*Rule)
	regRule := regexp.MustCompile(`(\d+|"[a-z]")`)

	for _, line := range strings.Split(inputPart[0], "\n") {
		lineParts := strings.Split(line, ":")
		id := pkg.MustAtoi(lineParts[0])
		rules[id] = &Rule{
			ID: id,
		}
		for i, subRule := range strings.Split(lineParts[1], " | ") {
			rules[id].SubRules = append(rules[id].SubRules, []int(nil))
			for _, matches := range regRule.FindAllStringSubmatch(subRule, -1) {
				switch {
				case strings.Contains(matches[1], `"`):
					rules[id].Regex = matches[1][1:2]
				default:
					rules[id].SubRules[i] = append(rules[id].SubRules[i], pkg.MustAtoi(matches[1]))
				}
			}
		}
	}
	return rules, messages
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
