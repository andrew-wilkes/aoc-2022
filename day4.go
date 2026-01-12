package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytes, _ := os.ReadFile("input.txt")
	data := string(bytes)
	lines := strings.Split(data, "\n")
	answer1 := 0
	answer2 := 0
	for _, line := range lines {
		pair := strings.Split(line, ",")
		r1 := strings.Split(pair[0], "-")
		r2 := strings.Split(pair[1], "-")
		a1, _ := strconv.Atoi(r1[0])
		a2, _ := strconv.Atoi(r1[1])
		b1, _ := strconv.Atoi(r2[0])
		b2, _ := strconv.Atoi(r2[1])
		if a1 >= b1 && a2 <= b2 || b1 >= a1 && b2 <= a2 {
			answer1++
		}
		if a1 >= b1 && a1 <= b2 || a2 >= b1 && a2 <= b2 || b1 >= a1 && b1 <= a2 || b2 >= a1 && b2 <= a2 {
			answer2++
		}
	}
	fmt.Printf("Part 1 answer = %d\n", answer1)
	fmt.Printf("Part 2 answer = %d\n", answer2)
}
