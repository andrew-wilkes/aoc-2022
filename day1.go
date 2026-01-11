package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	bytes, _ := os.ReadFile("input.txt")
	data := string(bytes)
	loads := strings.Split(data, "\n\n")
	counts := []int{}
	for _, load := range loads {
		cals := strings.Split(load, "\n")
		ccount := 0
		for _, cal := range cals {
			n, _ := strconv.Atoi(cal)
			ccount += n
		}
		counts = append(counts, ccount)
	}
	slices.Sort(counts)
	slices.Reverse(counts)
	fmt.Printf("Part 1 answer = %d\n", counts[0])

	top3Cals := counts[0] + counts[1] + counts[2]
	fmt.Printf("Part 2 answer = %d\n", top3Cals)
}
