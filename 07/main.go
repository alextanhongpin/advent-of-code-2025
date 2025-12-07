// You can edit this code!
// Click here and start typing.
package main

import (
	"cmp"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed test.txt
var testInput string

//go:embed input.txt
var input string

func main() {
	testInput = strings.TrimSpace(testInput)
	input = strings.TrimSpace(input)

	fmt.Println("test1:", part1(testInput)) // 21
	fmt.Println("prod1:", part1(input))     // 1613

	fmt.Println("test2:", part2(testInput), part2_v1(testInput)) // 40
	fmt.Println("prod2:", part2(input), part2_v1(input))         // 48021610271997
}

type Point struct {
	x, y int
}

func (p Point) Move(o Point) Point {
	return Point{
		p.x + o.x,
		p.y + o.y,
	}
}

var (
	down  = Point{y: 1}
	left  = Point{x: -1}
	right = Point{x: 1}
)

func part1(input string) int {
	rows := strings.Split(input, "\n")
	grid := make(map[Point]rune)
	var start Point
	for r, row := range rows {
		for c, ch := range row {
			p := Point{
				x: c,
				y: r,
			}

			grid[p] = ch
			if ch == 'S' {
				start = p
			}
		}
	}

	var countSplits func(Point) int
	countSplits = func(p Point) int {
		if _, ok := grid[p]; !ok {
			return 0
		}
		d := p.Move(down)
		switch grid[d] {
		case '^':
			l := d.Move(left)
			r := d.Move(right)
			grid[l] = '|'
			grid[r] = '|'
			return 1 + countSplits(l) + countSplits(r)
		case '.':
			grid[d] = '|'
			return countSplits(d)
		default:
			return 0
		}
	}

	return countSplits(start)
}

func part2(input string) int {
	rows := strings.Split(input, "\n")

	var start Point
	grid := make(map[Point]rune)
	for r, row := range rows {
		for c, ch := range row {
			p := Point{
				x: c,
				y: r,
			}

			if ch == 'S' {
				start = p
			}
			grid[p] = ch
		}
	}

	seen := make(map[Point]int)
	var countBeams func(p Point) int
	countBeams = func(p Point) int {
		if _, ok := grid[p]; !ok {
			return 1
		}

		if seen, ok := seen[p]; ok {
			return seen
		}

		d := p.Move(down)
		switch grid[d] {
		case '^':
			l := d.Move(left)
			r := d.Move(right)
			seen[d] = countBeams(l) + countBeams(r)
			return seen[d]
		default:
			seen[d] = countBeams(d)
			return seen[d]
		}
	}

	return countBeams(start)
}

func part2_v1(input string) int {
	rows := strings.Split(input, "\n")
	cols := 0
	grid := make(map[Point]rune)
	var start Point
	for r, row := range rows {
		for c, ch := range row {
			p := Point{
				x: c,
				y: r,
			}

			grid[p] = ch
			if ch == 'S' {
				start = p
			}
		}
		cols = max(cols, len(row))
	}
	var q []Point = []Point{start}
	for len(q) > 0 {
		var h Point
		h, q = q[0], q[1:]
		d := h.Move(down)
		ch, ok := grid[d]
		if !ok {
			continue
		}
		switch ch {
		case '^':
			l := d.Move(left)
			r := d.Move(right)
			grid[l] = '|'
			grid[r] = '|'
			q = append(q, l, r)
		case '.':
			grid[d] = '|'
			q = append(q, d)
		}
	}

	scores := make(map[Point]int)
	for r := len(rows) - 1; r > -1; r-- {
		for c := range cols {
			p := Point{y: r, x: c}
			switch grid[p] {
			case 'S', '|':
				scores[p] += cmp.Or(scores[p.Move(down)], 1)
			case '^':
				scores[p] += scores[p.Move(left).Move(down)] + scores[p.Move(right).Move(down)]
			}
		}
	}

	return scores[start]
}

/*
.......S.......
.......|.......
......|^|......
......|.|......
.....|^|^|.....
.....|.|.|.....
....|^|^|^|....
....|.|.|.|....
...|^|^|||^|...
...|.|.|||.|...
..|^|^|||^|^|..
..|.|.|||.|.|..
.|^|||^||.||^|.
.|.|||.||.||.|.
|^|^|^|^|^|||^|
|.|.|.|.|.|||.|

.......S.......
.......|.......
......|^|......
......|.|......
.....|^|^|.....
.....|.|.|.....
....|^|^|^|....
....|.|.|.|....
...|^|^|||^|...
...|.|.|||.|...
..|^|^|||^|^|..
..|.|.|||.|.|..
.|^|||^||.||^|.
.|.|||.||.||.|.
|^|^|^|^|^|||^|
1.1.1.1.1.111.1

.......S.......
.......|.......
......|^|......
......|.|......
.....|^|^|.....
.....|.|.|.....
....|^|^|^|....
....|.|.|.|....
...|^|^|||^|...
...|.|.|||.|...
..|^|^|||^|^|..
..|.|.|||.|.|..
.|^|||^||.||^|.
.|.|||.||.||.|.
1^1^1^1^1^111^1
1.1.1.1.1.111.1

.......S.......
.......|.......
......|^|......
......|.|......
.....|^|^|.....
.....|.|.|.....
....|^|^|^|....
....|.|.|.|....
...|^|^|||^|...
...|.|.|||.|...
..|^|^|||^|^|..
..|.|.|||.|.|..
.|^|||^||.||^|.
.|.|||.||.||.|.
121212121211121
1.1.1.1.1.111.1

.......S.......
.......|.......
......|^|......
......|.|......
.....|^|^|.....
.....|.|.|.....
....|^|^|^|....
....|.|.|.|....
...|^|^|||^|...
...|.|.|||.|...
..|^|^|||^|^|..
..|.|.|||.|.|..
.2^2|2^2|.|2^2.
.2.2|2.2|.|2.2.
121212121211121
1.1.1.1.1.111.1
*/
