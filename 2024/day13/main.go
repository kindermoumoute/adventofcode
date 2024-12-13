package main

import (
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	games := parse(input)

	part1, part2 := 0, 0

	// A*AX+B*BX=TX
	// A*AY+B*BY=TY
	// A=(TX-B*BX)/AX=(TY-B*BY)/AY
	// B=(TX-A*AX)/BX=(TY-A*AY)/BY

	// B
	// (TX-B*BX)*AY=(TY-B*BY)*AX
	// TX*AY-B*BX*AY=TY*AX-B*BY*AX
	// B*BY*AX-B*BX*AY=TY*AX-TX*AY
	// B(BY*AX-BX*AY)=TY*AX-TX*AY
	// B=(TY*AX-TX*AY)/(BY*AX-BX*AY)
	//
	// A
	// (TX-A*AX)/BX=(TY-A*AY)/BY
	// (TX-A*AX)*BY=(TY-A*AY)*BX
	// TX*BY-A*AX*BY=TY*BX-A*AY*BX
	// A*AY*BX-A*AX*BY=TY*BX-TX*BY
	// A=(TY*BX-TX*BY)/(AY*BX-AX*BY)

	for _, game := range games {
		if a, b := game.solveEquation(); game.valid(a, b) {
			part1 += a*3 + b
		}

		game.PriceX += 10000000000000
		game.PriceY += 10000000000000
		if a, b := game.solveEquation(); game.valid(a, b) {
			part2 += a*3 + b
		}
	}

	return part1, part2
}

func (game Game) valid(a, b int) bool {
	return game.ButtonAX*a+game.ButtonBX*b == game.PriceX && game.ButtonAY*a+game.ButtonBY*b == game.PriceY
}

func (game Game) solveEquation() (int, int) {
	b := (game.PriceY*game.ButtonAX - game.PriceX*game.ButtonAY) / (game.ButtonBY*game.ButtonAX - game.ButtonAY*game.ButtonBX)
	a := (game.PriceY*game.ButtonBX - game.PriceX*game.ButtonBY) / (game.ButtonBX*game.ButtonAY - game.ButtonBY*game.ButtonAX)
	return a, b
}

type Game struct {
	ButtonAX, ButtonAY int
	ButtonBX, ButtonBY int
	PriceX, PriceY     int
}

func parse(s string) []Game {
	lines := strings.Split(s, "\n\n")
	games := []Game(nil)
	for _, line := range lines {
		tmp := strings.Split(line, "\n")
		game := Game{}
		pkg.MustScanf(tmp[0], "Button A: X+%d, Y+%d", &game.ButtonAX, &game.ButtonAY)
		pkg.MustScanf(tmp[1], "Button B: X+%d, Y+%d", &game.ButtonBX, &game.ButtonBY)
		pkg.MustScanf(tmp[2], "Prize: X=%d, Y=%d", &game.PriceX, &game.PriceY)
		games = append(games, game)
	}
	return games
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
