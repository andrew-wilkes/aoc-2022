package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	positions := map[[2]int]bool{}

	// For part 2
	positions2 := map[[2]int]bool{}
	knots := [10][2]int{}
	for idx := range 10 {
		knots[idx] = [2]int{0, 0}
	}

	hpos := [2]int{0, 0}
	tpos := [2]int{0, 0}
	positions[tpos] = true

	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Split(line, " ")
		dir := data[0]
		steps, _ := strconv.Atoi(data[1])
		for range steps {
			switch dir {
			case "U":
				hpos[1]++
			case "D":
				hpos[1]--
			case "L":
				hpos[0]--
			case "R":
				hpos[0]++
			}

			// Move tpos
			tpos = move(hpos, tpos)
			positions[tpos] = true

			// Part 2
			knots[0] = hpos
			for i := range 9 {
				knots[i+1] = move(knots[i], knots[i+1])
			}
			positions2[knots[9]] = true
		}
	}
	fmt.Printf("Part 1 answer = %d\n", len(positions))

	fmt.Printf("Part 2 answer = %d\n", len(positions2))
}

func move(hpos, tpos [2]int) [2]int {
	dx := hpos[0] - tpos[0]
	dy := hpos[1] - tpos[1]
	if dx > 1 || dx < -1 {
		tpos[0] += dx / 2
		if dy > 0 {
			tpos[1] += 1
		}
		if dy < 0 {
			tpos[1] -= 1
		}
	} else if dy > 1 || dy < -1 {
		tpos[1] += dy / 2
		if dx > 0 {
			tpos[0] += 1
		}
		if dx < 0 {
			tpos[0] -= 1
		}
	}
	return tpos
}
