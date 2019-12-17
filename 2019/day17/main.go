package main

import (
	"image/color"
	"strconv"

	"github.com/kindermoumoute/adventofcode/pkg"

	"golang.org/x/image/colornames"

	"github.com/kindermoumoute/adventofcode/pkg/intcode"
	"github.com/kindermoumoute/adventofcode/pkg/twod"

	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	// graphic rendering config
	twod.RenderingMap = map[interface{}]color.Color{
		'.': colornames.Yellow,
		'#': colornames.Black,
		'^': colornames.Brown,
		'>': colornames.Brown,
		'<': colornames.Brown,
		'v': colornames.Brown,
	}

	c := intcode.New(input, 0)
	c.Run()

	// todo: refacto this?
	s := ""
	for _, output := range c.Output.Buff {
		s += string(uint8(output))
	}
	scaffold := twod.NewMapFromInput(s)

	tmp := scaffold.Find('<', '>', 'v', '^')
	startPos := tmp[0]
	dirs := map[int32]twod.Vector{
		'<': twod.LEFT, '>': twod.RIGHT, 'v': twod.DOWN, '^': twod.UP,
	}
	revertDirs := map[twod.Vector]int32{
		twod.LEFT: '<', twod.RIGHT: '>', twod.DOWN: 'v', twod.UP: '^',
	}
	scaffold = scaffold.Filter('<', '>', 'v', '^', '#')

	robot := twod.NewPoint(startPos, dirs[scaffold[startPos].(int32)])
	intersect := make(map[twod.Vector]int)
	twod.SetFPS(60)
	movementInstruction := ""
	for len(intersect) != len(scaffold) { // todo remove last step
		// render with point
		scaffold[robot.Pos] = revertDirs[robot.Speed]
		scaffold.Render()
		scaffold[robot.Pos] = '#'

		// point appearance counter
		intersect[robot.Pos]++

		// move
		movementInstruction += move(scaffold, robot)
	}
	if robot.Steps != 0 {
		movementInstruction += strconv.Itoa(robot.Steps) + "\n"
	}

	part1 := intersectSum(intersect, scaffold.FindTopLeft())

	scaffold[robot.Pos] = revertDirs[robot.Speed]
	scaffold.Render()

	instructions := "A,B,A,C,B,C,B,C,A,B\n"
	instructions += "L,6,L,4,R,8\n"
	instructions += "R,8,L,6,L,4,L,10,R,8\n"
	instructions += "L,4,R,4,L,4,R,8\n"
	instructions += "n\n"
	inputBuff := []int(nil)
	for _, ins := range instructions {
		inputBuff = append(inputBuff, int(ins))
	}

	c = intcode.New(input, 0, inputBuff...)
	c.Memory[0] = 2
	part2 := c.Run()

	return part1, part2
}

func move(m twod.Map, p *twod.P) string {
	movementInstruction := ""
	p.Move(1)
	v, exist := m[p.Pos]
	if !exist || v != '#' {
		p.Move(-1)
		if p.Steps != 0 {
			movementInstruction += strconv.Itoa(p.Steps) + ","
			p.Steps = 0
		}

		p.TurnRight()
		move := "R"
		p.Move(1)
		v, exist = m[p.Pos]
		if !exist || v != '#' {
			p.Move(-1)
			p.TurnRight()
			p.TurnRight()
			p.Move(1)
			move = "L"
		}

		movementInstruction += move + ","
	}
	return movementInstruction
}

func intersectSum(intersect map[twod.Vector]int, topL twod.Vector) int {
	sum := 0
	for k, v := range intersect {
		if v > 1 {
			dist := topL - k
			sum += pkg.Abs(dist.X() * dist.Y())
		}
	}
	return sum
}

func main() {
	execute.RunWithPixel(run, nil, puzzle, true)
}
