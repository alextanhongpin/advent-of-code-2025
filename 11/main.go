// You can edit this code!
// Click here and start typing.
package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

//go:embed test.txt
var testInput string

//go:embed test2.txt
var testInput2 string

//go:embed input.txt
var input string

func main() {
	testInput2 = strings.TrimSpace(testInput2)
	testInput = strings.TrimSpace(testInput)
	input = strings.TrimSpace(input)

	fmt.Println("test1:", part1(testInput)) // 5
	fmt.Println("prod1:", part1(input))     // 571

	fmt.Println("test2:", part2(testInput2)) // 2
	fmt.Println("prod2:", part2(input))      // 511378159390560
}

type state struct {
	curr string
	seen []string
}

func part1(input string) int {
	grid := make(map[string][]string)
	for row := range strings.SplitSeq(input, "\n") {
		row = strings.ReplaceAll(row, ":", "")
		paths := strings.Fields(row)
		grid[paths[0]] = paths[1:]
	}

	// Alternative solution, just showing how both methods works.
	// return countPath(grid, "you", "out")
	return len(solvePath(grid, "you", "out"))
}

func part2(input string) int {
	grid := make(map[string][]string)
	for row := range strings.SplitSeq(input, "\n") {
		row = strings.ReplaceAll(row, ":", "")
		paths := strings.Fields(row)
		grid[paths[0]] = paths[1:]
	}

	a := countPath(grid, "fft", "dac")
	b := countPath(grid, "dac", "out")
	c := countPath(grid, "svr", "fft")

	d := countPath(grid, "svr", "fft")
	e := countPath(grid, "dac", "fft")
	f := countPath(grid, "fft", "out")
	return a*b*c + d*e*f
}

func countPath(grid map[string][]string, in, out string) int {
	cache := make(map[string]int)
	var solve func(curr string, seen []string) int
	solve = func(curr string, seen []string) int {
		if curr == out {
			return 1
		}
		if n, ok := cache[curr]; ok {
			return n
		}

		var total int
		for _, o := range grid[curr] {
			if slices.Contains(seen, o) {
				continue
			}
			if _, ok := cache[o]; !ok {
				cache[o] = solve(o, append(slices.Clone(seen), o))
			}
			total += cache[o]
		}

		cache[curr] = total
		return total
	}
	return solve(in, []string{in})
}

func solvePath(grid map[string][]string, in, out string) [][]string {
	q := []state{
		{curr: in, seen: []string{in}},
	}

	var res [][]string
	for len(q) > 0 {
		var h state
		h, q = q[0], q[1:]
		if h.curr == out {
			res = append(res, h.seen)
			continue
		}
		for _, o := range grid[h.curr] {
			if slices.Contains(h.seen, o) {
				continue
			}
			s := state{
				curr: o,
				seen: append(h.seen, o),
			}
			q = append(q, s)
		}
	}
	return res
}
