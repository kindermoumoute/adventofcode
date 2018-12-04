package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/kindermoumoute/adventofcode/pkg"
)

type guard struct {
	ID          int
	minutes     int
	minutChosen int
	allMinutes  map[int]int
	ts          time.Time
}

// returns part1 and part2
func run(input string) (string, string) {
	day4 := parse(input)
	guards := make(map[int]*guard)
	currentGuard := 0
	maxSleep := guard{}
	for _, list := range day4 {
		switch list.text {
		case "falls asleep":
			guards[currentGuard].ts = list.toTs()
			fmt.Println("guard", currentGuard, " feels asleep ", list.toTs())
		case "wakes up":
			guards[currentGuard].minutes += int(list.toTs().Sub(guards[currentGuard].ts).Minutes())
			for i := guards[currentGuard].ts.Minute(); i != list.toTs().Minute(); {
				fmt.Println("inc min ", i)
				guards[currentGuard].allMinutes[i]++
				i = (i + 1) % 60
			}
			if guards[currentGuard].minutes > maxSleep.minutes {
				maxSleep.ID = guards[currentGuard].ID
				maxSleep.minutes = guards[currentGuard].minutes
				maxSleep.minutChosen = list.toTs().Minute() - 1
				maxSleep.allMinutes = guards[currentGuard].allMinutes
				fmt.Println("guard", currentGuard, " minut chosen ", maxSleep.minutChosen)

			}
			fmt.Println("guard", currentGuard, " wakes up ", list.toTs(), ", ", guards[currentGuard].minutes, " minutes asleep")
		default:
			ID := 0
			_, err := fmt.Sscanf(list.text, "Guard #%d", &ID)
			pkg.Check(err)
			_, ok := guards[ID]
			if !ok {
				guards[ID] = &guard{ID: ID, allMinutes: make(map[int]int)}
			}
			currentGuard = ID
			fmt.Println("new guard:", currentGuard)
		}
	}
	max := 0
	minutChosen := 0
	for i, min := range maxSleep.allMinutes {
		if min > max {
			max = min
			minutChosen = i
		}
	}

	ID := 0
	max = 0
	minutChosenPart2 := 0
	for _, guard := range guards {
		for i, min := range guard.allMinutes {
			if min > max {
				ID = guard.ID
				max = min
				minutChosenPart2 = i
			}
		}
	}
	// fmt.Println("PASS")
	return strconv.Itoa(minutChosen * maxSleep.ID), strconv.Itoa(minutChosenPart2 * ID)
}

type ByDate []day4
type day4 struct {
	year, month, day, hour, minute int

	text string
}

func (a ByDate) Len() int           { return len(a) }
func (a ByDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDate) Less(i, j int) bool { return a[i].toTs().Unix() < a[j].toTs().Unix() }

func (d day4) toTs() time.Time {
	return time.Date(d.year, time.Month(d.month), d.day, d.hour, d.minute, 0, 0, time.UTC)
}

func parse(s string) []day4 {
	lines := strings.Split(s, "\n")
	newDay := make([]day4, len(lines))
	for i, line := range lines {
		tmp := ""
		n, err := fmt.Sscanf(line, "[%d-%d-%d %d:%d] %s %s", &newDay[i].year, &newDay[i].month, &newDay[i].day, &newDay[i].hour, &newDay[i].minute, &newDay[i].text, &tmp)
		newDay[i].text += " " + tmp
		pkg.Check(err)
		if n != 7 {
			panic("N parameters expected")
		}
	}
	sort.Sort(ByDate(newDay))
	return newDay
}

func main() {
	tests.Run(run, false)
	fmt.Println(run(puzzle))
}
