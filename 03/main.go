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
	fmt.Println("test1:", part1(testInput)) // 357
	fmt.Println("prod1:", part1(input))     // 17324

	fmt.Println("test2:", part2(testInput)) // 3121910778619
	fmt.Println("prod2:", part2(input))     // 171846613143331
}

func part1(input string) int {
	total := 0
	input = strings.TrimSpace(input)
	for row := range strings.SplitSeq(input, "\n") {
		row = strings.TrimSpace(row)
		if row == "" {
			continue
		}

		total += maxJoltage(row, 2)
	}
	return total
}

func part2(input string) int {
	var total int
	input = strings.TrimSpace(input)
	for row := range strings.SplitSeq(input, "\n") {
		if row == "" {
			continue
		}

		total += maxJoltage(row, 12)
	}
	return total
}

func maxJoltage(s string, l int) int {
	if len(s) < l {
		return 0
	}

	if len(s) == l {
		return toInt(s)
	}

	var m string
	for i := 0; i < l+1; i++ {
		n := s[:i] + s[i+1:]
		nl := min(len(n), l)
		ml := min(len(m), l)
		if toInt(m[:ml]) < toInt(n[:nl]) {
			m = n
		}
	}

	return maxJoltage(m, l)
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		//panic(err)
	}
	return n
}
