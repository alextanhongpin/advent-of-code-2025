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

	fmt.Println("test1:", part1(testInput)) // 3
	fmt.Println("prod1:", part1(input))     // 770

	fmt.Println("test2:", part2(testInput)) // 14
	fmt.Println("prod2:", part2(input))     // 357674099117260
}

type Interval struct {
	x, y int
}

func (i *Interval) In(v int) bool {
	return v >= i.x && v <= i.y
}

func (i Interval) Overlap(o Interval) bool {
	return max(i.x, o.x) <= min(i.y, o.y)
}

func part1(input string) int {
	ranges, ingredients, ok := strings.Cut(input, "\n\n")
	if !ok {
		panic("invalid input row")
	}
	var intervals []Interval
	for row := range strings.SplitSeq(ranges, "\n") {
		a, b, ok := strings.Cut(row, "-")
		if !ok {
			panic("invalid input")
		}
		intervals = append(intervals, Interval{
			x: toInt(a),
			y: toInt(b),
		})
	}

	var count int
	for row := range strings.SplitSeq(ingredients, "\n") {
		x := toInt(row)
		for _, i := range intervals {
			if i.In(x) {
				count++
				break
			}
		}
	}
	return count
}

func part2(input string) int {
	ranges, _, ok := strings.Cut(input, "\n\n")
	if !ok {
		panic("invalid input row")
	}

	var intervals []Interval
	for row := range strings.SplitSeq(ranges, "\n") {
		a, b, ok := strings.Cut(row, "-")
		if !ok {
			panic("invalid input")
		}
		np := Interval{
			x: toInt(a),
			y: toInt(b),
		}
		intervals = append(intervals, np)
	}

	for range len(intervals) {
		h := intervals[0]
		r := intervals[1:]

		var overlap bool
		for j, i := range r {
			if i.Overlap(h) {
				overlap = true
				r[j] = Interval{
					x: min(i.x, h.x),
					y: max(i.y, h.y),
				}
				intervals = r
				break
			}
		}
		if !overlap {
			intervals = append(r, h)
		}
	}

	var total int
	for _, i := range intervals {
		total += i.y - i.x + 1
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
