package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// returns part1 and part2
func run(input string) (string, string) {
	particles := parse(input)

	part1 := -1
	dist := 99999999
	for i, part := range particles {
		tmp := part.a.distance()
		if tmp < dist {
			part1 = i
			dist = tmp
		}

	}

	for i := 0; i <= 50; i++ {
		particles = particles.detectCollisions()
		particles.update()

	}
	return strconv.Itoa(part1), strconv.Itoa(len(particles))
}

type particle struct {
	p, v, a vector
}

type vector struct {
	x, y, z int
}

func (v vector) distance() int {
	return int(math.Abs(float64(v.x)) + math.Abs(float64(v.y)) + math.Abs(float64(v.z)))
}

type particles []particle

func (p particles) detectCollisions() particles {
	for j := 0; j < len(p); j++ {
		del := false
		for k := 1; k < len(p); k++ {
			if j == k {
				continue
			}
			if p[j].p.x == p[k].p.x && p[j].p.y == p[k].p.y && p[j].p.z == p[k].p.z {
				p = append(p[:k], p[k+1:]...)
				del = true
				k--

			}
		}
		if del {
			p = append(p[:j], p[j+1:]...)
			j--
		}
	}
	return p
}

func (p particles) update() {
	for i := range p {
		p[i].v.x += p[i].a.x
		p[i].v.y += p[i].a.y
		p[i].v.z += p[i].a.z

		p[i].p.x += p[i].v.x
		p[i].p.y += p[i].v.y
		p[i].p.z += p[i].v.z
	}
}

func parse(s string) particles {
	lines := strings.Split(s, "\n")
	list := make([]particle, len(lines))
	for j, line := range lines {
		vectors := strings.Split(line, ", ")
		tempCoords := make([]int, len(vectors)*3)
		for i, vector := range vectors {
			coord := strings.Split(vector[3:len(vector)-1], ",")
			nb, err := strconv.Atoi(coord[0])
			check(err)

			tempCoords[i*3] = nb
			nb, err = strconv.Atoi(coord[1])
			check(err)

			tempCoords[i*3+1] = nb
			nb, err = strconv.Atoi(coord[2])
			check(err)

			tempCoords[i*3+2] = nb
		}

		list[j].p.x = tempCoords[0]
		list[j].p.y = tempCoords[1]
		list[j].p.z = tempCoords[2]

		list[j].v.x = tempCoords[3]
		list[j].v.y = tempCoords[4]
		list[j].v.z = tempCoords[5]

		list[j].a.x = tempCoords[6]
		list[j].a.y = tempCoords[7]
		list[j].a.z = tempCoords[8]
	}
	return list
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	for _, test := range tests {
		part1, part2 := run(test.input)
		if part1 != test.expectedPart1 {
			fmt.Println("Input ", test.input, " - PART1: ", part1, " but expected ", test.expectedPart1)
			return
		}
		if part2 != test.expectedPart2 {
			fmt.Println("Input ", test.input, " - PART2: ", part2, " but expected ", test.expectedPart2)
			return
		}
	}
	fmt.Println(run(puzzle))
}
