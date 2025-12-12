// You can edit this code!
// Click here and start typing.
package main

import (
	_ "embed"

	"bufio"
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

var re = regexp.MustCompile(`\d+`)

func main() {
	testInput = strings.TrimSpace(testInput)
	input = strings.TrimSpace(input)

	fmt.Println("test1:", part1(testInput)) // 2
	fmt.Println("prod1:", part1(input))     // 425
}

func part1(input string) int {
	var total int
	scanner := bufio.NewScanner(strings.NewReader(input))
	var presents []string
	for scanner.Scan() {
		text := scanner.Text()
		if !strings.Contains(text, "x") {
			var present string
			for scanner.Scan() {
				p := strings.TrimSpace(scanner.Text())
				if len(p) == 0 {
					break
				}
				present += p
			}
			presents = append(presents, present)
		} else {
			m := toIntSlice(re.FindAllString(text, -1))
			if arrange(m[0], m[1], m[2:], presents) {
				total++
			}
		}
	}
	return total
}

type state struct {
	box     [][]int
	choices []string
}

func arrange(width, height int, arr []int, presents []string) bool {
	box := make([][]int, height)
	for i := range box {
		box[i] = make([]int, width)
	}
	var choices []string
	var area int
	for i, r := range arr {
		for range r {
			choices = append(choices, presents[i])
			area += strings.Count(presents[i], "#")
		}
	}
	if area+10*9 < width*height {
		return true
	}
	if area > width*height {
		return false
	}

	var q []state = []state{{box: box, choices: choices}}
	for len(q) > 0 {
		var h state
		h, q = q[0], q[1:]

		if len(h.choices) == 0 {
			return true
		}
		choices := getCombinations(h.choices[0])
		for _, choice := range choices {
			for y := range height - 2 {
				for x := range width - 2 {
					var valid = true
					box := deepClone(h.box)
					for dy := range 3 {
						for dx := range 3 {
							box[y+dy][x+dx] += toInt(string(choice[dy*3+dx]))
							if box[y+dy][x+dx] > 1 {
								valid = false
								break
							}
						}
					}
					if !valid {
						continue
					}
					q = append([]state{{
						box:     box,
						choices: slices.Clone(h.choices[1:]),
					}}, q...)
				}
			}
		}
	}

	return false
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func toIntSlice(ss []string) []int {
	res := make([]int, len(ss))
	for i, s := range ss {
		res[i] = toInt(s)
	}
	return res
}

// Rotate 90 degree, flip vertical and horizontal.
var patterns = `012345678
210543876
678345012
630741852
876543210
258147036`

func getCombinations(s string) []string {
	s = strings.ReplaceAll(s, "#", "1")
	s = strings.ReplaceAll(s, ".", "0")
	var res []string
	for r := range strings.SplitSeq(patterns, "\n") {
		r = strings.TrimSpace(r)
		indices := strings.Split(r, "")
		var o string
		for _, i := range indices {
			n, _ := strconv.Atoi(i)
			o += string(s[n])
		}
		if slices.Contains(res, o) {
			continue
		}
		res = append(res, o)
	}
	return res
}

func deepClone[T any](ss [][]T) [][]T {
	c := slices.Clone(ss)
	for i, r := range c {
		c[i] = slices.Clone(r)
	}

	return c
}
