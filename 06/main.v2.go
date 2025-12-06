// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var re = regexp.MustCompile(`[\d\*\+]+`)

func main() {
	input = strings.TrimSpace(input)
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(in string) int {
	var res []string
	for row := range strings.SplitSeq(in, "\n") {
		parts := re.FindAllString(row, -1)
		if len(res) == 0 {
			res = parts
		} else {
			for i := range parts {
				res[i] += " " + parts[i]
			}
		}
	}
	var total int
	for _, v := range res {
		total += op(v)
	}

	return total
}

func part2(in string) int {
	rows := strings.Split(in, "\n")
	cols := len(rows[0])

	for i, r := range rows {
		if len(r) != cols {
			rows[i] += strings.Repeat(" ", cols-len(r))
		}
	}
	var total int
	var s string
	for range cols {
		for i, r := range rows {
			h, r := pop(r)
			rows[i] = r
			s += h
		}
		if s[len(s)-1] == '*' || s[len(s)-1] == '+' {
			total += op(s)
			s = ""
		}
	}
	return total
}

func op(s string) int {
	op, r := pop(s)
	rest := strings.Fields(strings.TrimSpace(r))
	switch op {
	case "+":
		var n int
		for _, v := range rest {
			n += toInt(v)
		}
		return n
	case "*":
		var n = 1
		for _, v := range rest {
			n *= toInt(v)
		}
		return n
	default:
		panic("invalid")
	}
}

func pop(s string) (string, string) {
	return string(s[len(s)-1]), s[:len(s)-1]
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

var input = `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `
