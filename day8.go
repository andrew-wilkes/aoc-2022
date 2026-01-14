package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	trees := [][]int{}
	visibility := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))
		visibility = append(visibility, make([]int, len(line)))
		trees = append(trees, row)
		for i, ch := range line {
			n, _ := strconv.Atoi(string(ch))
			row[i] = n
		}
	}

	width := len(trees[0])
	height := len(trees)

	// Scan for visible trees
	leftRight, rightLeft, upDown, downUp := -1, -1, -1, -1
	for i := range width {
		for j := range height {
			th := trees[j][i]
			if th > upDown {
				visibility[j][i] = 1
				upDown = th
			}
			th = trees[height-j-1][i]
			if th > downUp {
				visibility[height-j-1][i] = 1
				downUp = th
			}
		}
		upDown, downUp = -1, -1
	}
	for i := range height {
		for j := range width {
			th := trees[i][j]
			if th > leftRight {
				visibility[i][j] = 1
				leftRight = th
			}
			th = trees[i][width-j-1]
			if th > rightLeft {
				visibility[i][width-j-1] = 1
				rightLeft = th
			}
		}
		leftRight, rightLeft = -1, -1
	}

	// Count visible trees
	answer := 0
	for i := range height {
		for j := range width {
			answer += visibility[i][j]
		}
	}
	fmt.Printf("Part 1 answer = %d\n", answer)

	// Part 2
	answer = 0
	for x := 1; x < width-1; x++ {
		for y := 1; y < height-1; y++ {
			ss := scenicScore(x, y, width, height, trees)
			if ss > answer {
				answer = ss
			}
		}
	}
	fmt.Printf("Part 2 answer = %d\n", answer)
}

func scenicScore(x, y, width, height int, trees [][]int) int {
	up, down, left, right := 0, 0, 0, 0
	th := trees[y][x]

	for i := x - 1; i >= 0; i-- {
		left++
		if trees[y][i] >= th {
			break
		}
	}

	for i := x + 1; i < width; i++ {
		right++
		if trees[y][i] >= th {
			break
		}
	}

	for j := y - 1; j >= 0; j-- {
		up++
		if trees[j][x] >= th {
			break
		}
	}

	for j := y + 1; j < height; j++ {
		down++
		if trees[j][x] >= th {
			break
		}
	}

	return left * right * up * down
}
