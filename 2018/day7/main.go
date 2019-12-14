package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/kindermoumoute/adventofcode/pkg/execute"

	"github.com/kindermoumoute/adventofcode/pkg"
)

type Node struct {
	ID       string
	children Nodes
	parents  Nodes

	done     bool
	taken    bool
	workTime int
}

type Nodes []*Node

func (r Nodes) Len() int           { return len(r) }
func (r Nodes) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r Nodes) Less(i, j int) bool { return r[i].ID < r[j].ID }

func (r Nodes) String() string {
	if r == nil {
		return ""
	}
	str := ""
	for _, node := range r {
		str += node.ID

	}
	return str
}

func (r *Node) String() string {
	if r == nil {
		return ""
	}
	return r.ID + " - child " + r.children.String() + " - parents " + r.parents.String()
}

func (r *Node) SetDone(value bool) {
	r.done = value
}

func (r *Node) IsDone() bool {
	return r.done
}

func (r *Node) Time() int {
	return int(r.ID[0] - 'A' + 61)
}

func (r *Node) Work() {
	fmt.Printf("work %s waiting for availability\n", r.ID)
	for _, node := range r.parents {
		for !node.IsDone() {
		}
	}
	fmt.Printf("work %s ready\n", r.ID)
	time.Sleep(refTime * time.Duration(r.workTime))
	r.SetDone(true)
	fmt.Printf("work %s done\n", r.ID)
	return
}

func (r *Node) Available() bool {
	if r == nil {
		return false
	}
	for _, node := range r.parents {
		if !node.IsDone() {
			return false
		}
	}
	return true
}

var refTime = time.Millisecond * 5

func Part2(wlen int, allJobs map[string]*Node, jobs []string) string {
	var workers = make([]*Node, wlen)

	workStack := []*Node{}
	for _, job := range jobs {
		allJobs[job].SetDone(false)
		allJobs[job].workTime = int(allJobs[job].ID[0]-'A') + 61
		if wlen == 2 {
			allJobs[job].workTime -= 60
		}
		workStack = append(workStack, allJobs[job])
	}

	seconds := 0
	for done := ""; len(done) != len(jobs); {
		defers := []func(){}
		for i := range workers {
			if workers[i] != nil {
				workers[i].workTime--
				if workers[i].workTime <= 0 {
					done += workers[i].ID
					node := workers[i]
					defers = append(defers, func() {
						node.SetDone(true)
					})
					workers[i] = nil
				}
			}
			if workers[i] == nil {
				for j := 0; j < len(workStack); j++ {
					if workStack[j].Available() && !workStack[j].taken {
						workStack[j].workTime--
						workStack[j].taken = true
						workers[i] = workStack[j]
						break
					}
				}

			}
		}
		for _, f := range defers {
			f()
		}

		seconds++
	}
	return strconv.Itoa(seconds)
}

func (r *Node) Part1() []string {
	part1 := []string{}
	r.done = true
	for i := 0; i < len(r.children); i++ {
		node := r.children[i]
		if !node.done && node.Available() {
			node.done = true

			r.children = append(r.children[:i], r.children[i+1:]...)
			r.children = append(node.children, r.children...)
			sort.Sort(r.children)
			part1 = append(part1, node.ID)
			i = -1
		}
	}

	return part1
}

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	connections := parse(input)

	nodeIDs := make(map[string]*Node)
	for _, link := range connections {
		_, exist := nodeIDs[link.step1]
		if !exist {
			nodeIDs[link.step1] = &Node{ID: link.step1, children: make([]*Node, 0), parents: make([]*Node, 0)}
		}
		_, exist = nodeIDs[link.step2]
		if !exist {
			nodeIDs[link.step2] = &Node{ID: link.step2, children: make([]*Node, 0), parents: make([]*Node, 0)}
		}
	}
	for _, link := range connections {
		nodeIDs[link.step1].children = append(nodeIDs[link.step1].children, nodeIDs[link.step2])
		nodeIDs[link.step2].parents = append(nodeIDs[link.step2].parents, nodeIDs[link.step1])
	}
	root := &Node{children: make([]*Node, 0)}
	for _, node := range nodeIDs {
		if len(node.parents) == 0 {
			root.children = append(root.children, node)
			node.parents = append(node.parents, root)
		}
		sort.Sort(node.children)
	}
	if len(root.children) == 0 {
		panic("no ROOTS")
	}
	sort.Sort(root.children)

	wlen := 2
	if puzzle == input {
		wlen = 5
	}
	part1 := root.Part1()
	return strings.Join(part1, ""), Part2(wlen, nodeIDs, part1)
}

type steps struct {
	step1, step2 string
}

func parse(s string) []steps {
	lines := strings.Split(s, "\n")
	steps := make([]steps, len(lines))
	for i, line := range lines {
		n, err := fmt.Sscanf(line, "Step %s must be finished before step %s can begin.", &steps[i].step1, &steps[i].step2)
		pkg.Check(err)
		if n != 2 {
			panic(fmt.Errorf("%d args expected in scanf, got %d", 0, n))
		}
	}
	return steps
}

func main() {
	execute.Run(run, tests, puzzle, false)
}
