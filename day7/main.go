package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	tn, bottom := buildTree(puzzle)
	fmt.Println(tn[bottom])
}

type towerNode struct {
	name      string
	weight    int
	weightSum int

	parent *towerNode
	sub    []*towerNode
}

func buildTree(s string) (map[string]*towerNode, string) {
	bottomProgram := ""
	tempTN := make(map[string]*towerNode)
	tempChild := make(map[*towerNode][]string)
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		subPrograms := strings.Split(line, " -> ")
		prgmInfo := strings.Split(subPrograms[0], " (")
		w, _ := strconv.Atoi(prgmInfo[1][:len(prgmInfo[1])-1])
		tempTN[prgmInfo[0]] = &towerNode{name: prgmInfo[0], weight: w, weightSum: w}
		if len(subPrograms) > 1 {
			if bottomProgram == "" && strings.Count(s, prgmInfo[0]) == 1 {
				bottomProgram = prgmInfo[0]
			}
			tempChild[tempTN[prgmInfo[0]]] = strings.Split(subPrograms[1], ", ")
		}
	}
	for parent, children := range tempChild {
		parent.sub = make([]*towerNode, len(children))
		for i, child := range children {
			parent.sub[i] = tempTN[child]
		}
	}
	for _, parent := range tempTN {
		for _, child := range parent.sub {
			child.parent = parent
		}
	}
	for _, towerNode := range tempTN {
		for node := towerNode.parent; node != nil; node = node.parent {
			node.weightSum += towerNode.weight
		}
	}
	return tempTN, bottomProgram
}

func (tn *towerNode) String() string {
	str := tn.name + ":" + strconv.Itoa(tn.weight) + " ["
	for _, node := range tn.sub {
		str += node.String()
	}
	return str + "] "
}
