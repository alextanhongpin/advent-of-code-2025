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

	fmt.Println("test1:", part1(testInput)) // 1227775554
	fmt.Println("prod1:", part1(input))     // 28844599675

	fmt.Println("test2:", part2(testInput)) // 4174379265
	fmt.Println("prod2:", part2(input))     // 48778605167
}

func part1(input string) int {
	var total int
	for row := range strings.SplitSeq(input, ",") {
		a, b, ok := strings.Cut(row, "-")
		if !ok {
			panic("invalid string")
		}
		for i := toInt(a); i < toInt(b)+1; i++ {
			s := strconv.Itoa(i)
			if len(s)%2 != 0 {
				continue
			}
			m := len(s) / 2
			if s[:m] == s[m:] {
				total += i
			}
		}
	}
	return total
}

func part2(input string) int {
	var total int
	for row := range strings.SplitSeq(input, ",") {
		row = strings.TrimSpace(row)
		a, b, ok := strings.Cut(row, "-")
		if !ok {
			panic("invalid string")
		}
		for i := toInt(a); i < toInt(b)+1; i++ {
			s := strconv.Itoa(i)
			for j := 1; j < len(s); j++ {
				t := len(s) / j
				if t < 2 {
					break
				}
				if strings.Repeat(s[:j], t) == s {
					total += i
					break
				}
			}
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
