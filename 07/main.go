// You can edit this code!
// Click here and start typing.
package main

import (
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

	fmt.Println("test2:", part2(testInput)) // 40
	fmt.Println("prod2:", part2(input))     // 48021610271997
}

type Point struct {
	x, y int
}

func (p Point) Add(o Point) Point {
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
	var splits int

	for len(q) > 0 {
		var h Point
		h, q = q[0], q[1:]
		d := h.Add(down)
		ch, ok := grid[d]
		if !ok {
			continue
		}
		switch ch {
		case '^':
			l := d.Add(left)
			r := d.Add(right)
			grid[l] = '|'
			grid[r] = '|'
			q = append(q, l, r)
			splits++
		case '.':
			grid[d] = '|'
			q = append(q, d)
		}
	}
	/*
		for r := range rows {
			for c := range cols {
				fmt.Print(string(grid[Point{c, r}]))
			}
			fmt.Println()
		}
	*/

	return splits
}

func part2(input string) int {
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
		d := h.Add(down)
		ch, ok := grid[d]
		if !ok {
			continue
		}
		switch ch {
		case '^':
			l := d.Add(left)
			r := d.Add(right)
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
				score, ok := scores[p.Add(down)]
				if !ok {
					score = 1
				}
				scores[p] += score
			case '^':
				scores[p] += scores[p.Add(left).Add(down)] + scores[p.Add(right).Add(down)]
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
