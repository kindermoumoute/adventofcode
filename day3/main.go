package main

import (
	"fmt"
	"math"
)

func main() {
	puzzle := 312051

	//fmt.Println(part1(puzzle))
	fmt.Println(part2(puzzle))

}

type p struct {
	x, y int
}

func part2(puzzle int) int {
	// Part 2
	values := make(map[p]int)
	values[p{x: 0, y: 0}] = 1

	x := 0
	y := 0
	//xSign := 0
	//ySign := 1
	side := 1
	corner := 0
	for {
		side++
		corner = side / 2
		//value < puzzle
		for ; y < corner; y++ {
			if adjacentSum(x, y, values) > puzzle {
				goto found
			}
		}
		for ; x > -corner; x-- {
			if adjacentSum(x, y, values) > puzzle {
				goto found
			}
		}
		side++
		for ; y > -corner; y-- {
			if adjacentSum(x, y, values) > puzzle {
				goto found
			}
		}
		for ; x < corner; x++ {
			if adjacentSum(x, y, values) > puzzle {
				goto found
			}
		}
	}
found:

	return values[p{x, y}]
}

func adjacentSum(x, y int, v map[p]int) int {
	sum := 0
	fmt.Println(x, y)
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			value, exist := v[p{x + i, y + j}]
			if exist {
				sum += value
			}
		}
	}
	v[p{x, y}] = sum
	return sum
}

func sideSize(p int) int {
	tmpSqrt := math.Sqrt(float64(p))
	if tmpSqrt != float64(int(tmpSqrt)) {
		tmpSqrt++
	}
	return int(tmpSqrt)
}

func part1(puzzle int) int {
	side := sideSize(puzzle)
	x0 := side / 2
	y0 := x0
	stepsToNextCorner := side - 1
	x := 0
	y := 0
	if side%2 == 0 {
		topLeft := side * side
		topRight := side*side - stepsToNextCorner
		botRight := side*side - 2*stepsToNextCorner
		x0--
		y0--
		if topRight <= puzzle {
			x = topLeft - puzzle
			y = stepsToNextCorner
		} else {
			x = stepsToNextCorner
			y = puzzle - botRight
		}

	} else {
		botLeft := side*side - stepsToNextCorner
		if botLeft <= puzzle {
			x = puzzle - botLeft
			y = 0
		} else {
			x = 0
			y = botLeft - puzzle
		}
	}

	return int(math.Abs(float64(x0-x)) + math.Abs(float64(y0-y)))
}
