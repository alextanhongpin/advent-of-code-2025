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

var positions = []Point{
	{-1, 1},
	{-1, 0},
	{-1, -1},
	{0, -1},
	{0, 1},
	{1, 1},
	{1, 0},
	{1, -1},
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
		for _, n := range positions {
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

	var total int
	for {
		var local int
		for p, ch := range grid {
			if ch != '@' {
				continue
			}
			var neighbour int
			for _, n := range positions {
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
				local++
				total++
				delete(grid, p)
			}
		}
		if local == 0 {
			break
		}
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
