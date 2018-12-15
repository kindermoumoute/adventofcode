package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/beefsack/go-astar"

	"github.com/kindermoumoute/adventofcode/pkg"
)

// returns part1 and part2
func run(input string) (string, string) {
	d15 := parse(input)

	fmt.Println(d15.String())
	fmt.Println()
	return strconv.Itoa(d15.part1()), ""

	return "", ""
}

var ExpectedSteps = []int{0, 5988, 11922, 17802, 23628, 29370, 35028, 40593, 46080, 51462, 56760, 61974, 67032, 71955, 76734, 81405, 85968, 90372, 94662, 98781, 102780, 106659, 110638, 114563, 118320, 121750, 125346, 128952, 132300, 135459, 138330, 140988, 143520, 145926, 148206, 150360, 152388, 154327, 156370, 158145, 159800, 161335, 162750, 164045, 165088, 166050, 167440, 168589, 169728, 170765, 171700, 172533, 173264, 173893, 174420, 174845, 175168, 175389, 175160, 175289, 175560, 176351, 177010, 178227, 179520, 180765, 181962, 183111, 184212, 185265, 186270, 187227, 188208, 189946, 191660, 193350, 195016, 196658}

func (d *day15) part1() int {
	rounds := 0
	for {
		sort.SliceStable(d.warriors, func(i, j int) bool {
			return d.warriors[i].p.Score() < d.warriors[j].p.Score()
		})

		for i := 0; i < len(d.warriors); i++ {
			sum, winnerRace := d.warriors.HPs()
			if winnerRace != -1 {
				fmt.Println(rounds, sum)
				return rounds * sum
			}
			if d.rowLength == 32 && i == 0 {
				fmt.Println(rounds, rounds*sum, ExpectedSteps[rounds])
				if rounds*sum != ExpectedSteps[rounds] {
					//if rounds == 14 {
					panic("unexpected")
				}
			}

			w := d.warriors[i]
			attacked, canAttack := w.attack()
			if !canAttack {
				w.moveTowardClosestTarget()
				attacked, _ = w.attack()
			}
			if attacked != -1 {
				fmt.Println(w, "attack", d.warriors[attacked])
				d.warriors[attacked].HP -= w.attackPower
				if d.warriors[attacked].HP <= 0 {
					fmt.Println(d.warriors[attacked], "is dead")
					d.warriors = append(d.warriors[:attacked], d.warriors[attacked+1:]...)
					if attacked < i {
						i--
					}
				}
			}
		}
		rounds++

		if d.rowLength == 32 {
			m := d.String()
			bufio.NewReader(os.Stdin).ReadBytes('\n')
			fmt.Println("round", rounds)
			fmt.Println()
			fmt.Println(m)
			fmt.Println()
		}
	}
}

func (wrs Warriors) HPs() (int, Race) {
	HPs := make(map[Race]int)
	for _, w := range wrs {
		HPs[w.race] += w.HP
	}
	if len(HPs) == 1 {
		for race, sum := range HPs {
			return sum, race
		}
	}

	return HPs[ELF] + HPs[GOBLIN], -1
}

type Warriors []*warrior
type Race int

func (r Race) String() string {
	if r == ELF {
		return "E"
	}
	return "G"
}

type warrior struct {
	race        Race
	p           *pkg.PAstar
	HP          int
	attackPower int

	data *day15
}

func (wrs Warriors) String() string {
	str := ""
	for i := 0; i < len(wrs); i++ {
		str += " |" + wrs[i].String()
	}
	return str
}

func (w *warrior) String() string {
	return fmt.Sprintf("%v %v %d", w.race, w.p.P, w.HP)
}

func (w *warrior) attack() (int, bool) {
	li := -1
	ennemies := make(map[pkg.P]int)
	for i, target := range w.data.warriors {
		if w.race != target.race {
			ennemies[target.p.P] = i
		}
	}
	for _, p := range w.p.P.FourWays() {
		ti, exist := ennemies[p]
		if exist {
			target := w.data.warriors[ti]
			if li == -1 || target.HP < w.data.warriors[li].HP {
				li = ti
				continue
			}
			if target.HP > w.data.warriors[li].HP {
				continue
			}
			if target.HP == w.data.warriors[li].HP && target.p.Score() < w.data.warriors[li].p.Score() {
				li = ti
			}
		}
	}

	if li != -1 {
		return li, true
	}

	return -1, false
}

func (w *warrior) moveTowardClosestTarget() {
	w.BuildEmptyMap()

	//fmt.Println(w.p.P, "==>", nextStep.P)

	shortest := 9999999999.0
	var nextStep *pkg.PAstar
	var closestPointToAttack *pkg.PAstar
	for _, target := range w.data.warriors {
		if target.race != w.race {
			//fmt.Println(w.p.P, "==>", target.p.P)
			w.p.EmptyPoints[target.p.P] = target.p
			pathers, dist, found := astar.Path(w.p, target.p)
			delete(w.p.EmptyPoints, target.p.P)
			if !found || dist > shortest {
				continue
			}

			pointToAttack := pathers[1].(*pkg.PAstar)
			if dist == shortest && pointToAttack.Score() > closestPointToAttack.Score() {
				continue
			}
			closestPointToAttack = pointToAttack
			nextStep = pathers[len(pathers)-2].(*pkg.PAstar)
			shortest = dist
		}
	}
	if nextStep == nil {
		return
	}

	//fmt.Println(w.p.P, "==>", nextStep.P)
	w.p = nextStep
}

func (w *warrior) BuildEmptyMap() {
	newMap := make(map[pkg.P]*pkg.PAstar)
	for k, v := range w.data.emptyMap {
		newMap[k] = v
	}

	for _, v := range w.data.warriors {
		if v.p.Score() != w.p.Score() {
			delete(newMap, v.p.P)
		}
	}

	w.p.EmptyPoints = newMap
}

type day15 struct {
	warriors Warriors

	rowLength, columnLength int
	emptyMap                map[pkg.P]*pkg.PAstar
}

func (d *day15) String() string {

	tmp := make(map[pkg.P]*warrior)
	for _, w := range d.warriors {
		tmp[w.p.P] = w
	}

	str := make([]string, d.columnLength)
	for j := 0; j < d.columnLength; j++ {
		tmpWarriors := Warriors{}
		for i := 0; i < d.rowLength; i++ {
			p := pkg.P{X: i, Y: j}
			_, exist := tmp[p]
			if exist {
				tmpWarriors = append(tmpWarriors, tmp[p])
				str[j] += tmp[p].race.String()
				continue
			}
			_, exist = d.emptyMap[p]
			if exist {
				str[j] += "."
			} else {
				str[j] += "#"
			}
		}
		str[j] += tmpWarriors.String()
	}

	return strings.Join(str, "\n")
}

const (
	ELF = iota
	GOBLIN
)

func parse(s string) *day15 {
	lines := strings.Split(s, "\n")
	d15 := &day15{
		warriors:     Warriors{},
		emptyMap:     make(map[pkg.P]*pkg.PAstar),
		columnLength: len(lines),
	}
	for j, line := range lines {
		if d15.rowLength == 0 {
			d15.rowLength = len(line)
		}
		for i, point := range line {

			p := pkg.P{X: i, Y: j}
			pA := &pkg.PAstar{P: p, RowLenght: d15.rowLength}
			e := &warrior{
				p:           pA,
				HP:          200,
				attackPower: 3,
				data:        d15,
			}
			switch point {
			case '.':
			case 'E':
				e.race = ELF
				d15.warriors = append(d15.warriors, e)
			case 'G':
				e.race = GOBLIN
				d15.warriors = append(d15.warriors, e)
			default:
				continue
			}
			d15.emptyMap[p] = pA
		}
	}
	return d15
}

func main() {
	pkg.Execute(run, tests, puzzle, false)
}
