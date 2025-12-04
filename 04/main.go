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
	fmt.Println("test1:", part1(testInput)) // 13
	fmt.Println("prod1:", part1(input))     // 1445

	fmt.Println("test2:", part2(testInput)) // 43
	fmt.Println("prod2:", part2(input))     // 8317
}

type Point struct {
	x, y int
}

var neighbours = []Point{
	{-1, 1}, {0, 1}, {1, 1},
	{-1, 0} /*{0, 0}*/, {1, 0},
	{-1, -1}, {0, -1}, {1, -1},
}

func part1(input string) int {
	input = strings.TrimSpace(input)
	grid := make(map[Point]rune)
	for y, row := range strings.Split(input, "\n") {
		row = strings.TrimSpace(row)
		if row == "" {
			continue
		}
		for x, ch := range row {
			p := Point{x, y}
			grid[p] = ch
		}
	}

	var total int
	for p, ch := range grid {
		if ch != '@' {
			continue
		}
		var neighbour int
		for _, n := range neighbours {
			if grid[Point{
				x: p.x + n.x,
				y: p.y + n.y,
			}] == '@' {
				neighbour++
			}
			if neighbour >= 4 {
				break
			}
		}
		if neighbour < 4 {
			total++
		}
	}
	return total
}

func part2(input string) int {
	input = strings.TrimSpace(input)
	grid := make(map[Point]rune)
	for y, row := range strings.Split(input, "\n") {
		row = strings.TrimSpace(row)
		if row == "" {
			continue
		}
		for x, ch := range row {
			p := Point{x, y}
			grid[p] = ch
		}
	}
	size := len(grid)

	for {
		var total int
		for p, ch := range grid {
			if ch != '@' {
				continue
			}
			var neighbour int
			for _, n := range neighbours {
				np := Point{
					x: p.x + n.x,
					y: p.y + n.y,
				}
				if grid[np] == '@' {
					neighbour++
				}
				if neighbour >= 4 {
					break
				}
			}
			if neighbour < 4 {
				total++
				delete(grid, p)
			}
		}
		if total == 0 {
			break
		}
	}

	return size - len(grid)
}
