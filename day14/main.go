package main

import (
	"fmt"
	"strconv"
)

func main() {
	salt := []int{17, 31, 73, 47, 23}
	//puzzle := []int{'f', 'l', 'q', 'r', 'g', 'n', 'k', 'x', '-'}
	puzzle := []int{'j', 'x', 'q', 'l', 'a', 's', 'b', 'h', '-'}
	sum := 0
	grid := make([][]bool, 128)
	for i := 0; i < 128; i++ {
		tmpstr := strconv.Itoa(i)
		tmpsli := []int{}
		for _, l := range tmpstr {
			tmpsli = append(tmpsli, int(l))
		}
		tmp := append(append(puzzle, tmpsli...), salt...)
		hash := actualKnowHash(tmp)
		grid[i] = []bool{}

		for _, j := range hash {
			for mask := 128; mask > 0; mask = mask >> 1 {
				grid[i] = append(grid[i], (mask&j) > 0)
				if (mask & j) > 0 {
					sum++
				}
			}
		}
	}
	regions := 0
	for i := 0; i < 128; i++ {
		for j := 0; j < 128; j++ {
			if !grid[i][j] {
				continue
			}
			explore(grid, i, j)
			regions++

		}
	}
	fmt.Println(sum, regions)

}

func explore(grid [][]bool, i, j int) {
	if i < 0 || j < 0 || j > 127 || i > 127 || !grid[i][j] {
		return
	}
	grid[i][j] = false
	explore(grid, i-1, j)
	explore(grid, i, j-1)
	explore(grid, i+1, j)
	explore(grid, i, j+1)
}

func actualKnowHash(s []int) []int {
	list := newList(256)
	pos, skipSize := 0, 0
	for i := 0; i < 64; i++ {
		pos, skipSize = knotHash(list, s, pos, skipSize)
	}
	hash := make([]int, 16)
	for i, v := range list {
		hash[i/len(hash)] ^= v
	}
	return hash
}

func knotHash(list, puzzle []int, pos, skipSize int) (int, int) {
	for _, length := range puzzle {
		for i := 0; i < length/2; i++ {
			swapStart := (pos + i) % len(list)
			swapEnd := (pos + length - i - 1) % len(list)
			list[swapStart], list[swapEnd] = list[swapEnd], list[swapStart]
		}
		pos = (pos + length + skipSize) % len(list)
		skipSize = (skipSize + 1) % len(list)
	}
	return pos, skipSize
}

func newList(ln int) []int {
	list := make([]int, ln)
	for i := range list {
		list[i] = i
	}
	return list
}
