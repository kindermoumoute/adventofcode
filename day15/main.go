package main

import "fmt"

const (
	aFactor = 16807
	bFactor = 48271

	divider = 2147483647
)

func main() {
	A, B := 634, 301
	part1 := 0
	for i := 0; i < 40000000; i++ {
		A = (A * aFactor) % divider
		B = (B * bFactor) % divider
		if A&0xffff == B&0xffff {
			part1++
		}
	}

	A, B = 634, 301
	part2 := 0
	for i := 0; i < 5000000; i++ {
		for {
			A = (A * aFactor) % divider
			if A%4 == 0 {
				break
			}
		}
		for {
			B = (B * bFactor) % divider
			if B%8 == 0 {
				break
			}
		}
		if A&0xffff == B&0xffff {
			part2++
		}
	}
	fmt.Println(part1, part2)
}
