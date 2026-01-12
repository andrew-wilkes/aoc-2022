package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	bytes, _ := os.ReadFile("input.txt")
	data := string(bytes)
	lines := strings.Split(data, "\n")
	//fmt.Println(int('a'), int('Z'))

	var p1, p2 []int
	answer1 := 0
	rucksacks := [][]int{}
	for _, line := range lines {
		length := len(line) / 2
		left := line[0:length]
		right := line[length:]
		if len(right) != length {
			panic("Unequal lengths of item lists.")
		}
		// Convert chrs to ints
		p1 = make([]int, length)
		p2 = make([]int, length)
		for i, ch := range left {
			p1[i] = itemPriority(byte(ch))
			p2[i] = itemPriority(right[i])
		}
		for _, p := range p1 {
			if slices.Contains(p2, p) {
				answer1 += p
				break
			}
		}
		p1 = append(p1, p2...)
		rucksacks = append(rucksacks, p1)
	}
	fmt.Printf("Part 1 answer = %d\n", answer1)

	answer2 := 0
	for i := 0; i < len(rucksacks); i += 3 {
		for _, p := range rucksacks[i] {
			if slices.Contains(rucksacks[i+1], p) && slices.Contains(rucksacks[i+2], p) {
				answer2 += p
				break
			}
		}
	}
	fmt.Printf("Part 2 answer = %d\n", answer2)
}

func itemPriority(ch byte) int {
	p := int(ch)
	if p > 90 {
		p -= 96
	} else {
		p -= 38
	}
	return p
}
