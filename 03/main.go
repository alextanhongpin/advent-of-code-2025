// You can edit this code!
// Click here and start typing.
package main

import (
	_ "embed"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

//go:embed test.txt
var testInput string

//go:embed input.txt
var input string

func main() {
	fmt.Println("test1:", part1(testInput)) // 0
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

		n := 0
		for i := 0; i < len(row); i++ {
			for j := i + 1; j < len(row); j++ {
				m := toInt(string(row[i]))*10 + toInt(string(row[j]))
				if m > n {
					n = m
				}
			}
		}
		total += n
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

		// how do we find the largest number sequence?
		// 98765111119999
		total += maxJoltage(row)
	}
	return total
}

func maxJoltage(s string) int {
	if len(s) < 12 {
		return 0
	}

	if len(s) == 12 {
		return toInt(s)
	}

	var m string
	for i := 0; i < 13; i++ {
		a := new(big.Int)
		a.SetString(m, 10)

		n := s[:i] + s[i+1:]

		b := new(big.Int)
		b.SetString(n, 10)
		if a.Cmp(b) == -1 {
			m = n
		}
	}

	return maxJoltage(m)
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
