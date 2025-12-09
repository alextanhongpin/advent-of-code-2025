// You can edit this code!
// Click here and start typing.
package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed test.txt
var testInput string

//go:embed input.txt
var input string

func main() {
	testInput = strings.TrimSpace(testInput)
	input = strings.TrimSpace(input)

	fmt.Println("test1:", part1(testInput)) // 50
	fmt.Println("prod1:", part1(input))     // 4756718172

	fmt.Println("test2:", part2(testInput)) // 24
	fmt.Println("prod2:", part2(input))     // 1665679194
}

type Point struct {
	x, y int
}

func (p Point) Distance(o Point) int {
	x := (p.x - o.x)
	y := (p.y - o.y)
	return x*x + y*y
}

func part1(input string) int {
	var points []Point
	for row := range strings.SplitSeq(input, "\n") {
		a, b, ok := strings.Cut(row, ",")
		if !ok {
			panic("invalid input")
		}
		i, j := toInt(a), toInt(b)
		points = append(points, Point{i, j})
	}

	p := points
	var area int
	l := len(p)
	for i := range l {
		for j := range l {
			width := abs(p[i].x-p[j].x) + 1
			height := abs(p[i].y-p[j].y) + 1
			area = max(area, width*height)
		}
	}

	return area
}

func part2(input string) int {
	panic("not implemented")
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
