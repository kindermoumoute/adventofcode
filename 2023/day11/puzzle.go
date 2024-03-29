package main

import (
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

var tests = execute.TestCases{
	{
		`...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`,
		`374`,
		`1030`,
	},
	{
		puzzle,
		`9805264`,
		`779032247216`,
	},
}

var puzzle = `.......#.....#.....................#.............#.................#.........................#..............................................
#.....................#.......#.................................................................................#.......#...................
.......................................................................#.........................#..................................#.......
................................................................#..............#..............................................#.............
.........................................#............................................................#....................................#
.................#..................................................#.......................................................................
...........................#...............................................................................#................................
....#.......#...........................................#..........................#.............................#..........................
....................#.............#........#................................................................................................
...................................................#...................................................#...................#................
.........................#........................................#..................................................#......................
......#..........#...........................................#...........#............#.....................................................
......................................#..........................................#.........................#.....................#..........
.#...................#..................................#.......................................#...........................................
............#.................................#.....................#.......................................................................
............................................................................................................................................
..............................#.....................................................................................#..................#....
....................................................................................#........#......#.......................................
.........................#.........................#............#...........................................#...............................
.#......#.....#.......................#.................#........................................................#.....#....................
...........................................................................#.................................................#..............
...............................................#.......................................#...............#..........................#.........
......................#..................#..................................................................................................
..........................................................................................................................#.................
......#...........................#.........................................................#........................#......................
..............................................................................................................................#.............
..#.......................#.................#............................................................................................#..
..................#....................................................#................#...................................................
....................................................#.....#....................#............................................................
...............................#............................................................................................................
..............................................#...............#..........................................................#...........#......
.#.........#........................#..................................................................#....................................
..........................................................................#.......................#........................................#
...................#......................................................................#.................................................
.............................................................................................................................#..............
.....#......................................................................................................................................
........................................................#........................#..........................................................
.............................#...............#.....................#.....#............#..................#..............................#...
............................................................................................#......#...........#............................
.........................#........#...................................................................................#.....................
...........................................................................................................................................#
.....................................................#.....#.................................................................#..............
#.........................................................................................................#.......................#.........
................#...........................................................#.....................#.........................................
..............................#..............#...................................#.......#...........................#......................
........#................#.........#............................#............................................#.............#.........#......
............................................................................................................................................
.......................................................#...............#.............................#...........................#..........
.............#.................................#............................................................................................
............................................................................................................................................
..#.............................................................................#...........................................#...........#...
............................................................................................................................................
......................................#..........................................................#................#.........................
............................#................#......#......................#...........#....................................................
............................................................................................................................................
....................................................................................................#................#..........#...........
.#............#...................#...............................#........................................#.............................#..
.........#...............................................................#..................................................................
.......................#....................#........#...........................................#..........................................
............................................................#.....................................................#.........................
..............................#.......................................................................#...................#.......#.........
............#.......................#...................#............................#......................................................
...................................................#......................................#..............................................#..
.#................#...................................................#........................#............................................
......#....................................................................#................................#...............................
...............................#....................................................................................#.......................
.........................#.....................#......#.............................................#.................................#.....
............................................................................................................................................
...............#.....................#....................................................#.................................................
..................................................................................................................#.........................
....#......#.....................................#..............................#............................#............#.....#...........
....................#.............................................#....................#.........#.......................................#..
.........................................................#..............#.............................#.....................................
............................#...........#...........................................................................#.......................
............................................................................................................................................
#..................................#...............................................................#..........#.............................
.......#......................................................................................#.................................#.......#...
.....................#....................#...............................#.....#......................................#....................
................................#..........................#......#....................#..........................#.........................
..............#...................................#........................................................................#...............#
....#....................................................................................................#..................................
.......................#....................................................................................................................
......................................#................................#..................#.................................................
.....................................................#............................................#...............................#.........
............................................................................................................................................
............................................................................................................................................
....#.................#..................................#.........................#..........................#.............................
.............#........................................................................................#..................#..................
......................................#............................................................................#........................
.........#....................................................#.........#............................................................#......
.............................................................................................#..............................................
.............................#...........#............................................................................#.....................
#..................#................................................#....................#.......................................#.........#
...................................................................................................#......#.................................
...........#.............................................#...........................#......................................................
....................................#.......................................................................................................
..................................................#..........#.........#......................#.....................#.......................
.#......#....................................#..............................................................#...............................
.............................#.............................................#................................................................
.......................#...............................................................................#....................#...............
....................................................#.....................................#...........................................#.....
................#.....................#..................#.................................................................................#
...#.....#..........................................................#................#......................................................
.............................................#...............#................................#......#......................................
...............................#...............................................#............................................................
.............................................................................................................................#.........#....
.....................#................................................#.................................................#...................
...............#.........................#.......................#..................#.......................................................
.....#......................................................................................................#.......................#.......
....................................#.................#.........................#..................#........................................
........................................................................................................................................#...
.........#..............#.....................................#........................................#.....................#..............
...............................................#...................#........................................................................
......................................#.............#..................................#.....#.......................#......................
............#.....#...............................................................................#.........................................
...#...................................................................................................................................#....
.......................#.........................#..........................................................................................
.........................................................#...........#.........................#........#..........#.............#..........
............................#..............#....................#........................................................#..................
.......#.................................................................#................#.................................................
.................#..................................................................................#........................#...........#..
...................................................#.......#................................................................................
.........................................#......................................#......................................#....................
....................................#...........................................................................#...........................
..............................#.............................................................#...........#.......................#......#....
.........#...............................................#..............#...................................................................
........................#....................#.....................................#........................................................
#...........................................................................................................................................
............................#......................................................................#.......#................................
......................................#.....................................................................................................
.....#............#.........................................#..........................#..........................................#.........
...........................................#......................#..............................................#..........................
..............#..............................................................................#............................#.............#...
..#......#...........................................#..............................#..................#.....#..............................
....................#..........#.....#...............................#......................................................................
............................................................................................................................................
.................................................#..........................................................................................
.............#..........#...............................#...............#.............#..........#.........#........#.......................
............................................#................#.................#............................................................
....................................#....................................................................................#.............#....`
