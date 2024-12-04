package main

import (
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

var tests = execute.TestCases{
	{
		`#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`,
		`405`,
		`400`,
	},
	{
		puzzle,
		`27202`,
		`41566`,
	},
}

var puzzle = `.####..#.#.#.##..
........#..##....
..##..#.....#..##
......##.##.#####
######.#.####....
..##....#..##.#..
.#..#..#####.#...
..##...#..#...#.#
#######.#....####

#.###..####
..#####....
#.##...####
##..#.##..#
##.#.##.##.
##.#.##.##.
##..#.##..#
#.##...####
..###.#....
#.###..####
.##....#..#

.##...#...#####
...######..#.#.
..#.#..#.####..
#..#...##.##..#
#..#...##.##..#
..#.#..#.####..
#..######..#.#.
.##...#...#####
#..#..#.#...#.#
####.#####..#..
....#.##..##...
#..######..#.##
##.#...#..##.#.
##.#...#..##.#.
#..######..#.##
....#.##..##...
####.#####..#..

#####..#####.##
.............##
#..######..##..
.##.#..#.##.#.#
....####....##.
#..##..##..###.
#.########.#.#.
.##..##..##...#
#..##..##..##..
####....####..#
....####....#..

...#....##...
.#.##....#...
#..##.####.#.
#.#####.##.#.
##...#.##.#.#
##...#.##.#.#
#.#####.##.#.
#..##.####.#.
.#.##....#...
...#....##...
...#....##...
.#.##....#...
#..##.####.#.
#.###.#.##.#.
##...#.##.#.#

####.##.#..#.#...
....##.####.#.#..
....##.####.#.#..
####.##.#..#.#...
..#####...####.##
#####..##.###...#
.##...##.....###.
.##...##.....###.
#####..##.###...#
..#####...####.##
#.##.##.#..#.#...

#.##########.
##.#..##..#.#
..#.##..##.#.
#...#....#...
##.#......#.#
#.#.#....#.#.
###...##...##
###...##...##
#.#.#....#.#.
##.#......#.#
....#....#...

....##....##.....
.####.####.####..
###....##....###.
###....##....###.
.####.####.####..
....##....##.....
..##..#..#..##..#
.###############.
.##.#..##..#.##..

####.#..#.####..#
####..##..#######
#.#.#.##.#.#.####
##.#..##..#.##..#
.#####...####....
.###.#..#.###....
.##..#..#..##....
..#.######.#..##.
#..#......#..#..#
######..######..#
#.##..##..##.####
..##.####.##..##.
...##.##.##...##.
..#.#....#.#..##.
.##..#..#..##.##.
#####....########
##..##..##..#####

..###.#..#.##
#####.####.##
...##.#..#.##
...##.####.##
###.#......#.
###..######..
###..........
..##..##.#..#
...###.##.###

##...####...###
##...####...###
#....#..#....##
##..........###
#.##########.#.
..#..#..#.##...
.#...####...#.#
#..#.####.#..#.
.##..####..##.#
###..####..###.
#..##.##.##..#.
#.#........#.#.
#.....##.....#.

#..#####.#.
.#.#...####
.##.#####.#
#...###.#.#
#..####.#.#
.##.#####.#
.#.#...####
#..#####.#.
####..#..##
####....##.
####....##.
####..#..##
#..#####.#.
.#.#...####
.##.#####.#

.######.####...
..#..#....##...
..#..#....##...
.######.####...
###..######.##.
.........#..###
..#####.#.#####
#.#..#.##.....#
.##..##....####
...##...##..#..
.#....#.#.#..##
...##....#####.
.#.##.#.###.###

###.#...#.#.##.#.
#.#...####...#..#
#.#...####...#..#
###.#...#.#.##.#.
..####..#####....
.....###...#..###
.....###...#..###
..####..#####....
###.#...#.#.##.##

##.#..##.##.#
.#...........
.#####...##..
##.#...#.##.#
...#.###.##.#
.##..#..#..#.
#.##..#......
#..###..#..#.
#.#.##.......
...#.###.##.#
.#.###.#....#
##..#....##..
#..##########
##.##########
##..#....##..
.#.###.#....#
...#.###.##.#

##.######.##.
...#....#..##
####....#####
...#.##.#....
###..##..###.
#..######..##
..#.#..#.#..#
#.##.##.##.##
...##..##....
##...##...###
##...##...###
...##..##....
#.##.##.##.##
..#.#..#.#..#
#..######..##
###..##..###.
...#.##.#....

####..#.#####
#..#...##.##.
.##.##.######
#..#...#.....
#..#.##..####
###.#..#.....
#..###..#....
#..#......##.
####..##..##.

###......
..#.##..#
...######
...###...
###.#####
#..##....
#..##....
###.#####
...###...

..##..#########
####.....#.#.#.
###.#..####.###
.#..##..####.#.
#.#.##.##......
....####..#..##
###.##..##.....
.#..##....###.#
.#..##....###.#
###.##..##.....
....####..#..##
#.#.##.##......
.#..##..####.#.
###.#..####.###
####....##.#.#.
..##..#########
..##..#########

......#.##.#.
#....#......#
##..###....##
.#..#..#..#..
#....########
.#..#...##...
.#..#.##..##.
##..#########
..##........#
######......#
.####...##...
.#..#...##...
#....#......#
..##..#....#.
##..###....##

.###.###.
#.#..####
###.##.#.
###.##.##
#.#..####
.###.###.
.###.###.
#.#..####
###.##.##
###.##.#.
#.#..####
.###.###.
#...#.#.#

.#....#..
.##...#..
.###.##..
#.##.#...
.####..##
#.#.#.###
..#...#..
##...#...
##..####.
##..####.
##...#...
..#...#..
#.#.#.###
.####..##
#.##.#...
.###.##..
.##...#..

#..##....##..
.##.#.##....#
#..#.#..#..#.
#####....##..
....##.##...#
#..#...######
#..####.#..#.
.##.###..##..
....###......
.##..#..####.
#..###.#.##.#
.......######
.....####..##

####.#..##.##
.##.########.
.##.#.#..##.#
....#.#.#.###
.##..#.##.###
#..##..###.##
#####.###....
######.#.#..#
.##....###...
.........##.#
#####.##.#...
#..###.##.##.
....###.#..##
######.#..#.#
#..#.#.#..###
#..#.#.#..#.#
######.#..#.#

.#.#.####
.###.#..#
#..#.....
..#...##.
..#.##..#
.#.#.....
...#.####
###......
...###..#
...###..#
.##......

#....####..
.##.##..###
.....####..
.#..#...#.#
..#...#..#.
####.#.##..
####.#.##..
..#...#..#.
.#..#...#.#
.....####..
.##.##..###
#....####..
.....####..

.#..#..#.
#.#.####.
.#.######
#..######
####.##.#
####.##.#
....#..#.
..##....#
###......
###......
..##....#
....#..#.
.###.##.#

.#.##...#..##
#.#.#.#...#.#
.##.....##.#.
.##..##..####
...#.#####.##
#.#####..#...
#.#####..#...
.#.#.#####.##
.##..##..####
.##.....##.#.
#.#.#.#...#.#
.#.##...#..##
.#.##...#..##

####...##.#..##..
#######........##
..#..##.#....##.#
##....##..###.#.#
.####..###...###.
..#.###.#..#.#..#
..#.###.#..#.#..#
.####..###...###.
##....##..###.#.#
..#..##.#....##.#
#######........##
####...##.#..##..
#..#...#.##.##.#.
.###.#..#.###.#.#
.###.##.#.###.#.#
#..#...#.##.##.#.
####...##.#..##..

#...####.#.....
#..#.#.#.....#.
####.#.#....##.
#..#...#.#.####
.##......#.##..
.##...#####....
.##.#..#.......
.##.#..#.......
.##...#####....
.##......#.##..
#..#...#.#.####

.###.#.....
...##..#.#.
##....##...
..#..#.#...
##.##.##..#
##.##.##..#
..#..#.#...
##....##...
...##..#.#.
.###.#.....
..#.#.###.#
##.##..#.##
##.##..#.#.

.#....#.#.#.#....
#.####.#...####..
#..##..##.##..###
.#.##.#..#####.##
#..##..##.#....##
#.#..#.#.#.#.#...
.######....#.#...
#.####.#....#.###
...##...###..##..
..#..#..#.##..#..
#..##..#.#.##.#..
.#....#.#..##..##
##..#.###.####.##

.##..##.###
.######.#..
###..###...
#.#..#.###.
##....##...
##....###..
#..##..#.##
#########.#
.#.##.#..#.
..#..#...##
.#....#.##.
.#....#.##.
..#..#....#
.#.##.#..#.
#########.#
#..##..#.##
##....###..

#...#.#..#.
####.#..#.#
##.#..####.
#.###.#..#.
......####.
.###.######
#.#.##.##.#
...###....#
.....#....#
##..#..##..
##.#.######
##.#.######
##..#..##..
.....#....#
...###....#
#.#.##.##.#
.###.######

#.#####...#......
#.#####...#......
.##...###..#.#...
.#.####.####...#.
....##..#.##...##
###...#..#.##...#
....##.#.##.#.##.
##.#.#....#####.#
..##.#.#..###.#.#
##.###.####.#...#
#..###.####.#...#
..##.#.#..###.#.#
##.#.#....#####.#
....##.#.##.#.##.
###...#..#.##...#

##..###
.####..
#....##
..##...
#.##...
##..###
#.##.##

.#.#..#....##...#
.###..#....##...#
.#####.##..#.#..#
#.##.##.#........
#.##.##.#........
.#####.##..#.#..#
.###..#....##...#
.#.#..#....##...#
###..#..###.##.#.
##.....###.#..###
####...####..####

##..##.##.#
#...###..##
.#...######
#..##......
.##########
###........
###.#......
..#.#......
#.#.#######
.#..###..##
#.#.#......
.######..##
...#.######
..#.#..##..
..#####..##
#.#########
##.########

..##..#..
#.#..##..
#..#####.
.####...#
.####...#
#..####..
#.#..##..
..##..#..
##..#....
..#....##
..#....##
##..#....
..##..#..

.#..#....
#.##.###.
#....####
#....##.#
#.##.#..#
##..##.#.
#....#.#.
#....#..#
#....##.#
#....#.#.
##..##.#.
#.##.#..#
#....##.#
#....####
#.##.###.
.#..#....
##..###..

...####.##..#
#.....####..#
...#...#.#..#
###..#.###..#
.#.#.##.#.##.
...#.##.#.##.
###..#.###..#
...#...#.#..#
#.....####..#
...####.##..#
.....###.####

#.##...
.#.#...
##.#.##
.#.####
.#.##..
.##.#..
.###.##
#.#..##
..###..
.####..
.###...

.##..##.########.
.#.##.#...####...
.######.#.#..#.#.
.##..##.##....##.
.######.#.####.#.
########..#..#..#
..#..#..#.#..#.#.
..#..#..#.#..#.#.
.#......#.####.#.

#...###..##..##
.###..#####..##
.##.#.##.#.##.#
..#..###.##..##
#...#....##..##
##.#####..####.
#####...#.#..#.
###..#..###..##
###.###.#######
#...##...#....#
##.#.#.#.#.##.#
##.#.#.#.#.##.#
#...###..#....#

.#.##.#..##.#..
##..#.##.##..##
.#.#.#.....#...
#...###.####.##
#.#..#.##.##...
###..#.##.##...
#...###.####.##
.#.#.#.....#...
##..#.##.##..##
.#.##.#..##.#..
...#.#...#.##..
..###.#..####..
##..####.#.....
######.##...###
..#.....#.##.##

......#
......#
..#..#.
####.#.
..#.##.
##.##..
..#.###
...###.
##.###.
...###.
..#...#
##...#.
##...#.
####.##
##.#.##
###..#.
.#.##.#

..####..#.#...#..
.#....#.######..#
.#.##.#.....#....
.#.##.#.....#....
.#....#.######..#
..####..#.#...#..
..####..####.####
##....##..#.###..
.......##......#.

###..##..
..#.#..#.
###.####.
##..####.
...#....#
##.#.####
...######
...######
..#.#..#.
..#######
##.......
..#..##..
..#......
##...##..
#####..##

##........##...
...######......
..##....##..##.
###..##..######
##..#..#..#####
..###..###..###
.....##.....#.#
..#.####.#...#.
###..##..####..
...######.....#
..#..##..#..#..
#...####...#..#
##.#....#.##...
###......###.#.
..###..###..#..
###.#..#.####..
..#..##..#..#.#

#...##.#.
#...#..#.
...#.#...
.######..
.....#.#.
.#...####
..#..#...
..#..#...
.#...####
.....#.#.
.######..
...#.#...
#...#..#.
#...##.#.
#.#..#.##
#.#.##..#
.#.###.##

.######....#.#.
.........#...##
.#....#.#.#....
..#..#....#.###
#.####.##..#..#
...##...#..####
...##...#..####
#.####.##..#..#
..#..#....#.###
.#....#.#.#....
.........#..###

....###.#..
#...#...###
#...#...###
....###.#..
.......#.#.
#.##.##.###
##..#..#...
#####..#.##
##..#......
##...#.#.##
#.#..##.#..

..####...##.#
..####....###
##.##.##..###
...##...##.#.
.##..##.#####
##.##.#####..
..####...#.#.
........#...#
###..###..#.#
...##....#...
##....##..###
..####..##.##
##.##.###.##.

##...####...#
...#......##.
#####.##.####
##.#.#..#.#.#
###..#..#..##
....#....#...
...##.##.##..

.##.##...
.#...#...
###.#.###
#......##
.##.#####
###.#..##
.#....#..
##...#...
#.##.#.##
#.##.####
##...#...

####.#..#
...##.###
##.....#.
..#..###.
..#..###.
##.....#.
..###.###
####.#..#
....#.#.#
.##.#..#.
..#..##..
#.##.#.#.
#.##.#.#.
..#..##..
.##.#..#.

...#.#.
###.#..
......#
#####..
##..###
##.....
##.#...
##.#...
##.....
##..###
#####..
...#..#
###.#..

..#.#.#
####.##
######.
...#.#.
...#...
..##..#
...###.
###.##.
#####..
#####..
######.
...###.
..##..#
...#...
...#.#.
######.
####.##

.#..##.#.####
###.##....##.
..#####......
####.#.......
.##..###..##.
....#...#####
####..#.#####
#.##.##.#....
#######......
##.#.....####
#.##.###..##.
#.##.###..##.
##.#.....####
#######......
#.##.#..#....

....#.#...##.....
#..#..###...###.#
#..######..##.##.
#..#.#.#...#.#..#
.......#.##......
.##.##...####....
.....#....####..#

..####...
...##....
...##....
..####...
#.#..#.##
.#.##.#.#
..####...
########.
##...####
#..##..##
##....###
.######..
...##....

#..##.#..#..#.#
#..##.#..#..#.#
#.###...##.###.
....#.####.#.##
.##..#...#.#..#
###.#####..#.##
#.###...###.#.#
#..#.##....#.##
#..#.##....#.##
#.###...###.#.#
###.#####..#.##
.##..#...#.#..#
....#.####.#.##
#.###...##.#.#.
#..##.#..#..#.#

.##.#..#.
####....#
#####..##
.##.####.
.##.####.
..#.#..#.
####....#
#..#.##.#
....#..#.
####.##.#
.##.#..#.
####.##.#
.##......

....#..##..##
....#..##..##
..###.....##.
####.##..#.#.
#..#.#.#....#
.##.#.#..#.##
.#..##..#..##
#.#.....#.#..
###.....#.#..
.#..##..#..##
.##.#.#..#.##
#..#.#.#....#
####.##..#.#.
..###.....##.
....#..##..##

##.#.######.#
....##.##.##.
#####.#..#.##
##.##..##..##
...#.##..##.#
..##..####..#
..#....##....
##.###.##.#.#
####.#....#.#

.###.....###.##
####.#.##..#.#.
###..#.##..#.#.
.###.....###.##
#.#######.#..#.
#..#...#.#.#...
#...#.#####..#.
#...#.#####..#.
#..#...#.#.#...
#.#######.#..#.
.###.....###.##
###..#.##..#.#.
####.#.##..#.#.
.###.....###.##
.#....###.#.#..

..##.#....#
.######.##.
.######.##.
..##.#....#
.#.##.#....
###.##.####
##..#..###.
...#..#.#.#
###...#.#.#
###...#.#.#
...#..#.#.#
##..#..###.
.##.##.####
.#.##.#....
..##.#....#

#..#......###..
#..#......###..
#.####........#
#..##.....##..#
#..#.#######..#
.##..#...#...##
.....####..#...
#..##..##.#.#..
.##.###.#..#.##

###.###..
###.###..
..#.#.#.#
####.....
..#.....#
#....##..
###.##.##
#####..#.
#####..#.
###.#..##
#....##..

.###.#.##.#.###.#
##..#..##..#..###
#..##.#..#.##..##
##..#.#..#.#..##.
...###....###...#
.#...##..##...#..
..........#......
#....######....#.
..####.##.####...
.###.#.##.#.###.#
######.##.#######
##....#..#....###
##.#.#....#.#.##.
##.#.#....#.#.##.
##....#..#....###
######.##.#######
.###.#.##.#.###.#

.#.#..###..##
#.##.#####...
####.#..#.#..
.##...#.#####
#..#..###....
#..###.......
###..#.#.##..
..###.##.####
..#.#.##.####
###..#.#.##..
#..###.......
#..#..###....
.##...#.#####
####.#..#.#..
#.##.#####...
.#.#..###..##
##....#######

#.###...#.#####
######.#.####.#
#....##...##.#.
#....##...##.#.
######.#.####.#
#.###...#.#####
####..#.#..#..#
.##.###.##.##..
....########..#
##.##..###..#.#
##.##..###..#.#
....########..#
.##.###.##.##.#
####..#.#..#..#
#.###...#.#####

....##...
#######..
.......#.
....#.#..
####..#.#
#####...#
####.#..#
.....#.##
.##.#.#..
######.#.
.....#.##
.....###.
#####..#.
####.###.
######.##

.#...#.####..
.#...#.####..
..########...
.###..#.#.#..
#...#....#.##
.......#...#.
.#.######.#.#
###.#.#.####.
....###..##.#
...#.#.####..
...#.#.####..
....###..##.#
###.#.#.####.
.#.######.#.#
..#....#...#.
#...#....#.##
.###..#.#.#..

#.#..######
#.#.#######
.##.##....#
#..........
###.##.##.#
.##########
##.#...##..

....####..###
##.###.####.#
##..#........
##.#.#......#
..#.#.#.##.#.
....#.#....#.
..##.###..###
###.##..##..#
...##...##.#.
..#.....##...
###.##......#

.##..#.#..#..
#.....#######
#.#####.##...
#######..####
.####..####..
.####...###..
#######..####
#.#####.##...
#.....#######

.#.##....##.#....
.#.##....##.#....
####.####.#####..
#.##########.#.#.
.#.##....##.#..#.
####......#####.#
#...######......#
#.#.######.#.#..#
#.##.####.##.#.##
..#........#.....
#.#...##...#.#.##

####....#
##.#....#
.########
#..######
#.#......
#.#######
##..####.

###........##.##.
###...###.#..#...
#####.#.....#####
..#.##........#.#
..#.##........#.#
#####.#.....#####
###...###.#..##..
###........##.##.
######.###...#...
##..##.##.###.#.#
#####.#.##..#....
...#..#..#..#.##.
...##.####..###..
##.#....##.#....#
##.#.##.#####.###

#...#.##.
.###.#...
.#####...
#...#.##.
.......##
#.#...###
#.#...###

..#.##.###.##.#
.###.#.......##
#..#.#...#..#.#
########..####.
########..####.
#..#.#...#..#.#
.###.#.......##
..#.##.###.##.#
#...#...##.....
#.#..#.###.....
#..##.#.##.####
..#######.###.#
..###########.#

.#..##.
##.##..
##.##..
.#..##.
##.###.
#.#####
#..##.#
####..#
..##..#
..##..#
####...

##..##..##..#
....######...
######..#####
##.##.##.##.#
...########..
###.######.##
###.######.##
#####....####
..#........#.
###........##
##.########.#
..#.#....#.#.
##...####...#
######..#####
..#.##..##.##

#..####
....###
####.##
.##.###
####.#.
.##.#.#
.##.#..
....#.#
....#.#
.##.#..
.##.#.#
######.
.##.###

...######
.#..###.#
####.##.#
#.#.....#
#.#.....#
####.##.#
.#..###.#
...######
####.####
.#..##..#
##..#.##.
.....###.
.#..#.#..
.#..#.#..
.....##..

..#..#..#.#.....#
##.##.##.#...#.#.
##.##.##.#...#.#.
..#..#..#.#.....#
.#######.#..##.##
.........###..##.
.#....#..#.#..##.

......#..#.
..##...#.#.
#######.###
#....#.#..#
######....#
......#..#.
..##.......
##..######.
#.#####..#.
#....#.#.##
..##...#.##
..##......#
..##......#

...#.#..#
...######
..#......
..#######
...#..##.
...#.####
..#.#.##.
..##..##.
..#...##.
##.##.#..
##....##.
..#..#..#
###..#..#

.##.#..###.#...##
#..#.##....#.#..#
.........##.#.###
.##.....###...###
.##.##.#...###.#.
.##.##.#...###.#.
.##.....###...###
.........##.#.###
#..#.##.#..#.#..#

#.###.##.##
#..#......#
#.#.##..##.
.###.#..#.#
.###.#..#.#
#...##..##.
#..#......#
#.###.##.##
####......#

...##.##.#.#.#.
...##.##.#.#.#.
.##...#..##.#.#
.###.#.##..#.#.
....###..#.####
....###..#.####
.###.#.##..#.#.
.##...#..##.#.#
#..##.##.#.#.#.

#..##......##
#.#....##....
##.##.####.##
.#...##..##..
.#.####..####
.##...#..#...
.####.####.##
.#.##..##..##
#..##..##..##
#.#..#.##.#..
#..#.##..##.#
#.###.####.##
##.####..####
#.#....##....
..#..........

.#.##.#.#..##.#
#..##..##......
##....###......
.#....#.##..##.
#.#..#.#..#.##.
#......#.#.#..#
...##...#.#....
...##...#.#.##.
#########.#.##.
##....###......
#.####.#..#.##.

##..........##..#
..####..####..##.
..####..####.....
..##......##..##.
#..#..##..#..####
#.#........#.####
#####.###########

###########..
##.##.##...#.
#.#..#.##.#.#
###..######..
#..##..##....
#..##..##....
###..######..
#.#..#.##.#.#
##.##.##...#.
###########..
.#...##.#.#.#

..#.#.##..##.
....##.#.#.#.
#####......#.
#...#.#####..
##..##.....##
.#.##..###.#.
..###.##..#.#
..###.##..#.#
.#.##..###.#.
##...#.....##
#...#.#####..
#####......#.
....##.#.#.#.
..#.#.##..##.
..#.#.##..##.

..#.#.#..#...
..#.#.#..#...
....#...####.
..##.#.#.#..#
.#.#....#....
#####..#.#.##
.####.##.....
..######.#...
..##.#.#..###
..##...#..###
..######.#...
.####.##.....
#####..#.#.##
.#.#....#....
..##.#.#.#..#
....#...####.
..#.#.#..#...

.##.#.#.#..
.##.#.#.#..
#.####.#.#.
#.##.....##
...#.##.#..
.##.##.#..#
...##.##.#.
#.....#.#.#
#.....###.#
...##.##.#.
.##.##.#..#

.###..###
.##....##
#.#....#.
.########
##.#..#.#
##......#
##..###.#
#.#.##.#.
.#......#
..#....#.
####..###
...#..#..
...#..#..
####..###
..#....#.`