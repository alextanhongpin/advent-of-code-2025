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

	// Part 1.
	fmt.Println("test1:", part1(testInput)) // 3
	fmt.Println("prod1:", part1(input))     // 1158

	fmt.Println("test2:", part2(testInput)) // 6
	fmt.Println("prod2:", part2(input))     // 6860
}

func part1(input string) int {
	c := 0 // The number of times at 0.
	d := 50
	for row := range strings.SplitSeq(input, "\n") {
		dir, n := row[0], toInt(row[1:])
		switch dir {
		case 'L':
			d -= n
		case 'R':
			d += n
		}
		d %= 100
		d += 100
		d %= 100
		if d == 0 {
			c += 1
		}
	}

	return c
}

func part2(input string) int {
	c := 0 // The number of times it pass by 0.
	d := 50
	for row := range strings.SplitSeq(input, "\n") {
		dir, n := row[0], toInt(row[1:])
		c += n / 100
		for range n % 100 {
			switch dir {
			case 'L':
				d--
			case 'R':
				d++
			}
			d += 100
			d %= 100
			if d == 0 {
				c++
			}
		}
	}

	return c
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
