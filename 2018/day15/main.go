package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/beefsack/go-astar"

	"github.com/kindermoumoute/adventofcode/pkg"
)

// returns part1 and part2
func run(input string) (string, string) {
	d15 := parse(input)

	//fmt.Println(d15.String())
	//fmt.Println()
	return strconv.Itoa(d15.part1()), ""
}

func (d *day15) part1() int {
	rounds := 0
	for {
		sort.SliceStable(d.warriors, func(i, j int) bool {
			return d.warriors[i].p.Score() < d.warriors[j].p.Score()
		})

		for i := 0; i < len(d.warriors); i++ {
			sum, winnerRace := d.warriors.HPs()
			if winnerRace != -1 {
				//fmt.Println(rounds, sum)
				return rounds * sum
			}

			w := d.warriors[i]
			attacked, canAttack := w.attack()
			if !canAttack {
				w.moveTowardClosestTarget()
				attacked, _ = w.attack()
			}
			if attacked != -1 {
				//fmt.Println(w, "attack", d.warriors[attacked])
				d.warriors[attacked].HP -= w.attackPower
				if d.warriors[attacked].HP <= 0 {
					//fmt.Println(d.warriors[attacked], "is dead")
					d.warriors = append(d.warriors[:attacked], d.warriors[attacked+1:]...)
					if attacked < i {
						i--
					}
				}
			}
		}
		rounds++

		//if d.rowLength == 32 {
		//	m := d.String()
		//	fmt.Println("round", rounds)
		//	fmt.Println()
		//	fmt.Println(m)
		//	fmt.Println()
		//}
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
	p           *pkg.Tile
	HP          int
	attackPower int

	data *day15
}

func (wrs Warriors) String() string {
	str := ""
	for i := 0; i < len(wrs); i++ {
		str += " " + wrs[i].String()
	}
	return str
}

func (w *warrior) String() string {
	return fmt.Sprintf("%v(%d)", w.race, w.HP)
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

	shortest := 9999999999.0
	var nextStep *pkg.Tile
	var closestPointToAttack *pkg.Tile

	for _, way := range w.p.P.FourWays() {
		p, exist := w.p.EmptyPoints[way]
		if exist {
			p.EmptyPoints = w.p.EmptyPoints
			for _, target := range w.data.warriors {
				if target.race != w.race {
					//fmt.Println(w.p.P, "==>", target.p.P)
					p.EmptyPoints[target.p.P] = target.p
					pathers, dist, found := astar.Path(p, target.p)
					delete(p.EmptyPoints, target.p.P)
					if !found || dist > shortest {
						continue
					}

					pointToAttack := pathers[1].(*pkg.Tile)
					if dist == shortest && pointToAttack.Score() > closestPointToAttack.Score() {
						continue
					}

					if dist == shortest && p.Score() > nextStep.Score() {
						continue
					}
					closestPointToAttack = pointToAttack
					nextStep = p
					shortest = dist
				}
			}
		}
	}
	if nextStep == nil {
		return
	}

	w.p = nextStep
}

func (w *warrior) BuildEmptyMap() {
	newMap := make(map[pkg.P]*pkg.Tile)
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
	emptyMap                map[pkg.P]*pkg.Tile
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
		emptyMap:     make(map[pkg.P]*pkg.Tile),
		columnLength: len(lines),
	}
	for j, line := range lines {
		if d15.rowLength == 0 {
			d15.rowLength = len(line)
		}
		for i, point := range line {

			p := pkg.P{X: i, Y: j}
			pA := &pkg.Tile{P: p, RowLenght: d15.rowLength}
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
