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

	fmt.Println("test1:", part1(testInput)) // 0
	fmt.Println("prod1:", part1(input))     // 0

	fmt.Println("test2:", part2(testInput)) // 0
	fmt.Println("prod2:", part2(input))     // 0
}

func part1(input string) int {
	for row := range strings.SplitSeq(input, "\n") {
		_ = row
	}
	return 0
}

func part2(input string) int {
	for row := range strings.SplitSeq(input, "\n") {
		_ = row
	}
	return 0
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
