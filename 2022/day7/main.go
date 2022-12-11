package main

import (
	"path/filepath"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	tree := parse(input)

	part1, part2 := 0, 0

	dirSum := make(map[string]int)
	diskSize := 0
	for path, file := range tree {
		diskSize += file.Size
		for {
			path = filepath.Dir(path)
			if path == "/" {
				break
			}
			dirSum[path] += file.Size
		}
	}

	requiredSpace := 30000000 - (70000000 - diskSize)
	for _, s := range dirSum {
		if s >= requiredSpace && (part2 == 0 || s < part2) {
			part2 = s
		}

		if s > 100000 {
			continue
		}
		part1 += s
	}

	return part1, part2
}

type File struct {
	Name string
	Size int
}

func parse(s string) map[string]File {
	currentFolder := "/"
	tree := make(map[string]File)
	lines := strings.Split(s, "\n")

	for _, line := range lines {
		words := strings.Split(line, " ")
		switch words[0] {
		case "$": // command
			if words[1] == "cd" {
				switch words[2] {
				case "/":
				case "..":
					currentFolder = filepath.Dir(currentFolder)
				default:
					currentFolder = filepath.Join(currentFolder, words[2])
				}
			}
		case "dir": // show directory
		default: // show a file
			name := words[1]
			tree[filepath.Join(currentFolder, name)] = File{
				Name: name,
				Size: pkg.MustAtoi(words[0]),
			}
		}
	}
	return tree
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
