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

	fmt.Println(solver(testInput)) // 24
	fmt.Println(solver(input))     // 1665679194
}

type Point struct {
	x, y int
}

func solver(input string) (int, int) {
	var points []Point
	for row := range strings.SplitSeq(input, "\n") {
		a, b, ok := strings.Cut(row, ",")
		if !ok {
			panic("invalid input")
		}
		p := Point{toInt(a), toInt(b)}
		points = append(points, p)
	}

	type data struct {
		minX, maxX int
		minY, maxY int
	}
	var pairs [][]Point
	var greens []data
	for i := range len(points) - 1 {
		a, b := points[i], points[i+1]
		greens = append(greens, data{
			minX: min(a.x, b.x),
			maxX: max(a.x, b.x),
			minY: min(a.y, b.y),
			maxY: max(a.y, b.y),
		})
		pairs = append(pairs, []Point{a, b})
	}

	var part1 int
	var part2 int
	for i, a := range points {
		for _, b := range points[i+1:] {
			d := data{
				minX: min(a.x, b.x),
				maxX: max(a.x, b.x),
				minY: min(a.y, b.y),
				maxY: max(a.y, b.y),
			}
			area := (d.maxX - d.minX + 1) * (d.maxY - d.minY + 1)
			part1 = max(part1, area)
			if area > part2 {
				var valid = true
				for _, g := range greens {
					/* Basically find all rextangles that does not overlap with any of the green segments.
                     d.minX| g.minX| |d.maxX |g.maxX
					 -- d.minY
					 -- g.minY
					 -- d.maxY
					 -- g.maxY
					 x
					*/
					if d.minX < g.maxX && d.minY < g.maxY && d.maxX > g.minX && d.maxY > g.minY {
						valid = false
						break
					}
				}
				if !valid {
					continue
				}
				part2 = max(part2, area)
			}
		}
	}

	return part1, part2
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
