// You can edit this code!
// Click here and start typing.
package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"slices"
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

	fmt.Println("test1:", part1(testInput)) // 4277556
	fmt.Println("prod1:", part1(input))     // 6371789547734

	fmt.Println("test2:", part2(testInput)) // 3263827
	fmt.Println("prod2:", part2(input))     // 11419862653216
}

var re = regexp.MustCompile(`[\*\+]`)

func part1(input string) int {
	rows := strings.Split(input, "\n")
	indices := re.FindAllStringIndex(rows[len(rows)-1], -1)
	slices.Reverse(indices)

	var total int
	for _, pos := range indices {
		p := pos[0]
		var parts []string
		for i, row := range rows {
			rows[i] = row[:p]
			parts = append(parts, strings.TrimSpace(row[p:]))
		}
		total += compute(parts)
	}

	return total
}

func part2(input string) int {
	rows := strings.Split(input, "\n")
	var cols int
	for _, row := range rows {
		cols = max(cols, len(row))
	}
	for i, row := range rows {
		rows[i] += strings.Repeat(" ", cols-len(row))
	}
	var s string
	for i := cols - 1; i > -1; i-- {
		for _, row := range rows {
			s += string(row[i])
		}
	}

	var re = regexp.MustCompile(`(\d+|\+|\*)`)
	matches := re.FindAllString(s, -1)

	var total int
	for {
		i := slices.IndexFunc(matches, func(s string) bool {
			return s == "*" || s == "+"
		})
		if i == -1 {
			break
		}
		total += compute(matches[:i+1])
		matches = matches[i+1:]
	}

	return total
}

func compute(s []string) int {
	s, t := s[:len(s)-1], s[len(s)-1]
	switch t {
	case "*":
		var n = 1
		for _, v := range s {
			n *= toInt(v)
		}
		return n
	case "+":
		var n int
		for _, v := range s {
			n += toInt(v)
		}
		return n
	default:
		panic("invalid operator: " + t)
	}
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
