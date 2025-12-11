// You can edit this code!
// Click here and start typing.
package main

import (
	"container/heap"
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

var re = regexp.MustCompile(`[\[\]\(\)\{\}]+`)

func main() {
	testInput = strings.TrimSpace(testInput)
	input = strings.TrimSpace(input)

	fmt.Println("test1:", part1(testInput)) // 7
	fmt.Println("prod1:", part1(input))     // 441

	fmt.Println("test2:", part2(testInput)) // 33
	fmt.Println("prod2:", part2(input))     // 18559
}

func part1(input string) int {
	var total int
	for row := range strings.SplitSeq(input, "\n") {
		row = re.ReplaceAllString(row, "")
		row = strings.ReplaceAll(row, ".", "0")
		row = strings.ReplaceAll(row, "#", "1")
		parts := strings.Fields(row)
		press := minPress(parts)
		total += press
	}
	return total
}

func part2(input string) int {
	panic("see python solution")
}

func minPress(parts []string) int {
	parts = parts[:len(parts)-1] // Remove joltage
	ind := parseBin(parts[0])
	buttons := make([]int, len(parts[1:]))
	for i, btn := range parts[1:] {
		buttons[i] = parseButton(btn, len(parts[0]))
	}
	item := &Item{
		ind:     ind,
		buttons: buttons,
	}
	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, item)

	// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		if item.ind == 0 {
			return item.priority
		}
		for i, btn := range item.buttons {
			ind := item.ind
			ind ^= btn
			btns := slices.Clone(item.buttons)
			btns = slices.Delete(btns, i, i+1)
			heap.Push(&pq, &Item{
				priority: item.priority + 1,
				buttons:  btns,
				ind:      ind,
			})
		}
	}

	panic("invalid")
}

func parseButton(s string, size int) int {
	b := []rune(strings.Repeat("0", size))
	for s := range strings.SplitSeq(s, ",") {
		b[toInt(s)] = '1'
	}
	return parseBin(string(b))
}

func parseBin(s string) int {
	n, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(n)
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func parseInd(s string) []bool {
	ind := make([]bool, len(s))
	for i, v := range s {
		ind[i] = v == '#'
	}
	return ind
}

// An Item is something we manage in a priority queue.
type Item struct {
	ind      int
	buttons  []int
	priority int // The priority of the item in the queue.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // don't stop the GC from reclaiming the item eventually
	*pq = old[0 : n-1]
	return item
}
