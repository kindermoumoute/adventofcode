package main

import (
	"strconv"

	"github.com/kindermoumoute/adventofcode/pkg/execute"

	"github.com/kindermoumoute/adventofcode/pkg/font"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {

	bs := parse(input)
	minCount, minI := 99999999, 0
	layers := make(map[int]map[byte]int)
	render := make(map[int]byte)
	for i, b := range bs {
		layers[i] = make(map[byte]int)
		for j, char := range b {
			layers[i][char]++
			val, exist := render[j]
			if !exist || val == 2 {
				render[j] = char
			}
		}
		if layers[i][0] < minCount {
			minCount = layers[i][0]
			minI = i
		}

	}

	screen := []interface{}(nil)
	for j := range bs[0] {
		screen = append(screen, int(render[j]))
	}

	word := font.NewWordFromScreen(screen, 25, 1)

	return strconv.Itoa(layers[minI][1] * layers[minI][2]), word.String()
}

func parse(s string) [][]byte {
	b := []byte(s)
	for i := range b {
		b[i] -= '0'
	}

	bs := [][]byte(nil)
	size := 25 * 6
	for i := 0; i < len(b)/size; i++ {
		bs = append(bs, b[i*size:(i+1)*size])
	}
	return bs
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
