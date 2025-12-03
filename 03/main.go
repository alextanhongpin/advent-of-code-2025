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
	fmt.Println("prod1:", part1(input))     // 17324
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

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
