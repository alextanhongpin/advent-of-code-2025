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
	fmt.Println("test1:", part1(testInput)) // 0
	fmt.Println("prod1:", part1(input))     // 0

	fmt.Println("test2:", part2(testInput)) // 0
	fmt.Println("prod2:", part2(input))     // 0
}

type Point struct {
	x, y int
}

func (p *Point) In(v int) bool {
	return v >= p.x && v <= p.y
}

func (p Point) Overlap(o Point) bool {
	return max(p.x, o.x) <= min(p.y, o.y)
}

func part1(input string) int {
	input = strings.TrimSpace(input)
	ranges, ingredients, ok := strings.Cut(input, "\n\n")
	if !ok {
		panic("invalid input row")
	}
	var points []Point
	for row := range strings.SplitSeq(ranges, "\n") {
		if strings.TrimSpace(row) == "" {
			continue
		}
		a, b, ok := strings.Cut(row, "-")
		if !ok {
			panic("invalid input")
		}
		points = append(points, Point{
			x: toInt(a),
			y: toInt(b),
		})
	}

	var count int
	for row := range strings.SplitSeq(ingredients, "\n") {
		row = strings.TrimSpace(row)
		if row == "" {
			continue
		}
		x := toInt(row)
		for _, p := range points {
			if p.In(x) {
				count++
				break
			}
		}
	}
	return count
}

func part2(input string) int {
	input = strings.TrimSpace(input)
	ranges, _, ok := strings.Cut(input, "\n\n")
	if !ok {
		panic("invalid input row")
	}

	var points []Point
	for row := range strings.SplitSeq(ranges, "\n") {
		if strings.TrimSpace(row) == "" {
			continue
		}
		a, b, ok := strings.Cut(row, "-")
		if !ok {
			panic("invalid input")
		}
		np := Point{
			x: toInt(a),
			y: toInt(b),
		}
		points = append(points, np)
	}

	for range len(points) {
		h := points[0]
		r := points[1:]
		var overlap bool
		for j, p := range r {
			if p.Overlap(h) || h.Overlap(p) {
				overlap = true
				r[j] = Point{
					x: min(p.x, h.x),
					y: max(p.y, h.y),
				}
				points = r
				break
			}
		}
		if !overlap {
			points = append(r, h)
		}
	}
	// Merge all points.
	var total int
	for _, p := range points {
		total += p.y - p.x + 1
	}
	return total
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
