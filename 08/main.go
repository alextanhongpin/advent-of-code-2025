// You can edit this code!
// Click here and start typing.
package main

import (
	_ "embed"
	"fmt"
	"maps"
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

	fmt.Println("test1:", part1(testInput, 10)) // 40
	fmt.Println("prod1:", part1(input, 1_000))  // 123420

	fmt.Println("test2:", part2(testInput)) // 25272
	fmt.Println("prod2:", part2(input))     // 673096646
}

func part1(input string, connections int) int {
	rows := strings.Split(input, "\n")
	us := NewUnionSet[string]()

	var parts [][]string
	for i, r := range rows {
		for _, s := range rows[i+1:] {
			parts = append(parts, []string{r, s})
		}
		_ = us.Find(r)
	}
	slices.SortFunc(parts, func(a, b []string) int {
		return dist(a[0], a[1]) - dist(b[0], b[1])
	})

	for _, p := range parts[:connections] {
		l, r := p[0], p[1]
		if us.Find(l) == us.Find(r) {
			continue
		}
		us.Unite(l, r)
	}

	circuits := make(map[string]int)
	for _, p := range rows {
		circuits[us.Find(p)]++
	}
	sizes := slices.Collect(maps.Values(circuits))
	slices.Sort(sizes)
	slices.Reverse(sizes)

	var total int = 1
	for i := range 3 {
		total *= sizes[i]
	}

	return total
}

func part2(input string) int {
	rows := strings.Split(input, "\n")
	us := NewUnionSet[string]()

	var parts [][]string
	for i, r := range rows {
		for _, s := range rows[i+1:] {
			parts = append(parts, []string{r, s})
		}
		_ = us.Find(r)
	}
	slices.SortFunc(parts, func(a, b []string) int {
		return dist(a[0], a[1]) - dist(b[0], b[1])
	})

	done := func() bool {
		circuits := make(map[string]int)
		for _, p := range rows {
			circuits[us.Find(p)]++
		}
		sizes := slices.Collect(maps.Values(circuits))
		return len(sizes) == 1
	}

	var l, r string
	for _, p := range parts {
		l, r = p[0], p[1]
		if us.Find(l) == us.Find(r) {
			continue
		}

		us.Unite(l, r)
		if done() {
			lhs := strings.Split(l, ",")[0]
			rhs := strings.Split(r, ",")[0]
			return toInt(lhs) * toInt(rhs)
		}
	}

	panic("not possible :)")
}

func dist(a, b string) int {
	as := strings.Split(a, ",")
	bs := strings.Split(b, ",")
	if len(as) != len(bs) {
		panic("invalid len")
	}
	var dist int
	for i := range len(as) {
		n := toInt(as[i]) - toInt(bs[i])
		dist += n * n
	}
	return dist
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return n
}

type UnionSet[T comparable] struct {
	parents map[T]T
}

func NewUnionSet[T comparable]() *UnionSet[T] {
	return &UnionSet[T]{
		parents: make(map[T]T),
	}
}
func (u *UnionSet[T]) Find(t T) T {
	v, ok := u.parents[t]
	if !ok {
		u.parents[t] = t
		return t
	}
	if v == t {
		return v
	}
	return u.Find(u.parents[t])
}

func (u *UnionSet[T]) Unite(i, j T) {
	l := u.Find(i)
	r := u.Find(j)
	u.parents[l] = r
}
