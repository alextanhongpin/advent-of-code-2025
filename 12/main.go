// You can edit this code!
// Click here and start typing.
package main

import (
	"bufio"
	_ "embed"
	"fmt"
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

	fmt.Println("test1:", part1(testInput)) // 0
	fmt.Println("prod1:", part1(input))     // 0

	fmt.Println("test2:", part2(testInput)) // 0
	fmt.Println("prod2:", part2(input))     // 0
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
				txt := strings.TrimSpace(scanner.Text())
				if len(txt) == 0 {
					break
				}
				present += txt
			}
			presents = append(presents, present)
		} else {
			dim, arr, ok := strings.Cut(text, ": ")
			if !ok {
				panic("invalid input")
			}
			as := strings.Fields(arr)
			a := make([]int, len(as))
			for i, s := range as {
				a[i] = toInt(s)
			}
			ws, hs, ok := strings.Cut(dim, "x")
			if !ok {
				panic("invalid dimenstion")
			}
			w := toInt(ws)
			h := toInt(hs)

			if arrange(w, h, a, presents) {
				fmt.Println(w, h, a)
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
	if area < (width * height * 3 / 4) {
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
		var count int
		choices := getCombinations(h.choices[0])
		var ok bool
		for _, choice := range choices {
			for y := range height - 2 {
				for x := range width - 2 {
					var valid = true
					box := deepClone(h.box)
					bs := ""
					for _, r := range box {
						bs += sliceAsString(r)
					}
					if !strings.Contains(bs, "000") && len(h.choices) < 3 {
						return false
					}
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
					ok = true
					q = append([]state{{
						box:     box,
						choices: slices.Clone(h.choices[1:]),
					}}, q...)
				}
			}
		}
		if !ok {
			count++
			if count > 10 {
				break
			}
		}
	}

	return false
}

func sliceAsString(s []int) string {
	var t string
	for i := range s {
		t += strconv.Itoa(s[i])
	}
	return t
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
