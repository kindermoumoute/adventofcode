package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(part(1, parse(puzzle)), part(2, parse(puzzle)))
}

func part(part int, tab []int) int {
	i := 0
	itmp := 0
	compteur := 0
	for i < len(tab) {
		itmp = i
		i += tab[i]
		if part == 2 && tab[itmp] >= 3 {
			tab[itmp]--
		} else {
			tab[itmp]++
		}
		compteur++
	}
	return compteur
}

func parse(s string) []int {
	tempTab := strings.Split(s, "\n")
	tab := make([]int, len(tempTab))
	for i, v := range tempTab {
		tab[i], _ = strconv.Atoi(v)
	}
	return tab
}
