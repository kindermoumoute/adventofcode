package main

import (
	"fmt"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg/execute"

	"github.com/kindermoumoute/adventofcode/pkg"
)

const (
	left = iota
	straight
	right
)

var part1 = ""

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	track, carts := parse(input)
	alive := len(carts)
	for {
		positions := make(lookupMap)
		for _, cart := range carts {
			if cart == nil {
				continue
			}
			switch track[cart.pos.Y][cart.pos.X] {
			case '\\':
				if cart.pos.CurrentDirection == pkg.DOWN || cart.pos.CurrentDirection == pkg.UP {
					cart.pos.TurnRight()
				} else {
					cart.pos.TurnLeft()
				}
			case '/':
				if cart.pos.CurrentDirection == pkg.DOWN || cart.pos.CurrentDirection == pkg.UP {
					cart.pos.TurnLeft()
				} else {
					cart.pos.TurnRight()
				}

			case '+':
				switch cart.lastMove {
				case right:
					cart.pos.TurnLeft()
				case left:
					cart.pos.TurnRight()
				}
				cart.lastMove = (cart.lastMove + 1) % 3
			case '-':
			case '|':
			default:
				panic(fmt.Errorf("%v no track", cart))
			}

			positions.AddCartPosition(cart)
			cart.pos.Move(1)
			positions.AddCartPosition(cart)
		}
		dead := positions.CheckCrash(carts)
		alive -= dead
		if alive == 1 {
			for _, remainingCart := range carts {
				if remainingCart == nil {
					continue
				}
				return part1, fmt.Sprintf("%d,%d", remainingCart.pos.X, remainingCart.pos.Y)
			}
		}
	}
}

type Cart struct {
	id       int
	pos      pkg.P
	lastMove int
}

type lookupMap map[pkg.P][]*Cart

func (l lookupMap) CheckCrash(carts []*Cart) int {
	dead := 0
	for _, pos := range l {
		if len(pos) > 1 {
			if part1 == "" {
				part1 = fmt.Sprintf("%d,%d", pos[0].pos.X, pos[0].pos.Y)
			}
			for _, cart := range pos {
				if carts[cart.id] != nil {
					carts[cart.id] = nil
					dead++
				}
			}
		}
	}
	return dead
}
func (l lookupMap) AddCartPosition(c *Cart) {
	p := pkg.P{X: c.pos.X, Y: c.pos.Y}
	_, exist := l[p]
	if !exist {
		l[p] = []*Cart{c}
	} else {
		l[p] = append(l[p], c)
	}
}

func parse(s string) ([][]byte, []*Cart) {
	lines := strings.Split(s, "\n")
	track := make([][]byte, len(lines))
	carts := make([]*Cart, 0)
	for j, line := range lines {
		track[j] = []byte(line)
		for i := range track[j] {
			id := len(carts)
			dir := 0
			switch track[j][i] {
			case 'v':
				dir = pkg.UP
				track[j][i] = '|'
			case '<':
				dir = pkg.RIGHT
				track[j][i] = '-'
			case '^':
				dir = pkg.DOWN
				track[j][i] = '|'
			case '>':
				dir = pkg.LEFT
				track[j][i] = '-'
			default:
				continue
			}
			carts = append(carts, &Cart{
				id:  id,
				pos: pkg.P{X: i, Y: j, CurrentDirection: dir},
			})
		}
	}
	return track, carts
}

func main() {
	execute.Run(run, tests, puzzle, false)
}
