package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/kindermoumoute/adventofcode/pkg"
)

// returns part1 and part2
func run(input string) (string, string) {
	records := parse(input).reposeRecords()
	return records.findPart1().result(), records.findPart2().result()
}

func parse(s string) Records {
	lines := strings.Split(s, "\n")
	records := make(Records, len(lines))
	for i, line := range lines {
		year, month, day, hour, minute, textPart1, textPart2 := 0, 0, 0, 0, 0, "", ""
		n, err := fmt.Sscanf(line, "[%d-%d-%d %d:%d] %s %s", &year, &month, &day, &hour, &minute, &textPart1, &textPart2)
		records[i].TS = time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.UTC)
		records[i].text = textPart1 + " " + textPart2
		pkg.Check(err)
		if n != 7 {
			panic("N parameters expected")
		}
	}
	sort.Sort(records)
	return records
}

type Records []Record

type Record struct {
	TS   time.Time
	text string
}

func (r Records) Len() int           { return len(r) }
func (r Records) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r Records) Less(i, j int) bool { return r[i].TS.Unix() < r[j].TS.Unix() }

func (r Records) reposeRecords() ReposeRecords {
	reposeRecords := make(ReposeRecords)
	currentGuardID := 0
	var napStartTime time.Time
	for _, record := range r {
		switch record.text {
		case "falls asleep":
			napStartTime = record.TS
		case "wakes up":
			reposeRecords[currentGuardID].cumulatedMinutes += int(record.TS.Sub(napStartTime).Minutes())
			for i := napStartTime.Minute(); i != record.TS.Minute(); i = (i + 1) % 60 {
				reposeRecords[currentGuardID].minutes[i]++
			}
		default:
			_, err := fmt.Sscanf(record.text, "Guard #%d", &currentGuardID)
			pkg.Check(err)
			if _, ok := reposeRecords[currentGuardID]; !ok {
				reposeRecords[currentGuardID] = &ReposeRecord{GuardID: currentGuardID}
			}
		}
	}
	return reposeRecords
}

type ReposeRecords map[int]*ReposeRecord

type ReposeRecord struct {
	GuardID          int
	cumulatedMinutes int
	minutes          [60]int
}

// find guard with longest cumulated minutes
// then find longest minute from this guard
func (g ReposeRecords) findPart1() result {
	part1 := result{}
	for currentGuard, guard := range g {
		if guard.cumulatedMinutes > part1.max {
			part1.ID = g[currentGuard].GuardID
			part1.max = g[currentGuard].cumulatedMinutes
		}
	}
	part1.max = 0
	for minute, minuteCount := range g[part1.ID].minutes {
		if minuteCount > part1.max {
			part1.max = minuteCount
			part1.minute = minute
		}
	}
	return part1
}

// then find longest minute among all guards
func (g ReposeRecords) findPart2() result {
	part2 := result{}
	for _, guard := range g {
		for minute, minuteCount := range guard.minutes {
			if minuteCount > part2.max {
				part2.max = minuteCount
				part2.ID = guard.GuardID
				part2.minute = minute
			}
		}
	}
	return part2
}

type result struct {
	ID     int
	minute int

	max int // temp variable used for finding max
}

func (r result) result() string {
	return strconv.Itoa(r.minute * r.ID)
}

func main() {
	tests.Run(run, false)
	fmt.Println(run(puzzle))
}
