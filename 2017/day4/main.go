package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	fmt.Println(part(1), part(2))
}

func part(n int) int {
	phrases := strings.Split(puzzle, "\n")
	validWords := len(phrases)
	for _, p := range phrases {
		tempMap := make(map[string]int)
		for _, w := range strings.Split(p, " ") {
			if n == 1 {
				tempMap[w]++
			} else {
				if strings.Count(w, w[:1]) < len(w) {
					w = sortString(w)
					tempMap[w]++
				}
			}
			if tempMap[w] > 1 {
				validWords--
				break
			}
		}
	}
	return validWords
}

type ByLetter []rune

func (a ByLetter) Len() int           { return len(a) }
func (a ByLetter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLetter) Less(i, j int) bool { return a[i] < a[j] }

func sortString(s string) string {
	runes := []rune(s)
	sort.Sort(ByLetter(runes))
	return string(runes)
}
