package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type vec struct {
	x, y int
}

type position struct {
	v    vec
	dist int
}

func main() {
	grid := [][]rune{}
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var y int
	var start, end position

	for scanner.Scan() {
		line := scanner.Text()
		row := make([]rune, len(line))
		for x, r := range line {
			if r == 'S' {
				start = position{vec{x, y}, 0}
				r = 'a'
			}
			if r == 'E' {
				end = position{vec{x, y}, 0}
				r = 'z'
			}
			row[x] = r
		}
		grid = append(grid, row)
		y++
	}
	xLimit := len(grid[0]) - 1
	yLimit := len(grid) - 1

	// Find the shortest distance from S to E
	// Use BFS

	queue := []position{start}
	dist := 0
	visited := map[vec]bool{}
	visited[start.v] = true
	for len(queue) > 0 {
		pos := queue[0]
		dist = pos.dist
		if pos.v == end.v {
			break
		}
		queue = queue[1:]
		elev := grid[pos.v.y][pos.v.x]
		// Explore around the current position
		if pos.v.x > 0 {
			addPos(pos.v.x-1, pos.v.y, dist, elev, &queue, visited, grid)
		}
		if pos.v.x < xLimit {
			addPos(pos.v.x+1, pos.v.y, dist, elev, &queue, visited, grid)
		}
		if pos.v.y > 0 {
			addPos(pos.v.x, pos.v.y-1, dist, elev, &queue, visited, grid)
		}
		if pos.v.y < yLimit {
			addPos(pos.v.x, pos.v.y+1, dist, elev, &queue, visited, grid)
		}
	}
	fmt.Printf("Part 1 answer = %d\n", dist)

	// Part 2

	// Find shortest distance from any 'a' level to E
	shortest := math.MaxInt
	for y := range len(grid) {
		for x := range len(grid[0]) {
			if grid[y][x] == 'a' {
				v := vec{x, y}
				visited = map[vec]bool{}
				queue := []position{{v, 0}}
				visited[v] = true
				found := false
				for len(queue) > 0 {
					pos := queue[0]
					dist = pos.dist
					if pos.v == end.v {
						found = true
						break
					}
					queue = queue[1:]
					elev := grid[pos.v.y][pos.v.x]
					// Explore around the current position
					if pos.v.x > 0 {
						addPos(pos.v.x-1, pos.v.y, dist, elev, &queue, visited, grid)
					}
					if pos.v.x < xLimit {
						addPos(pos.v.x+1, pos.v.y, dist, elev, &queue, visited, grid)
					}
					if pos.v.y > 0 {
						addPos(pos.v.x, pos.v.y-1, dist, elev, &queue, visited, grid)
					}
					if pos.v.y < yLimit {
						addPos(pos.v.x, pos.v.y+1, dist, elev, &queue, visited, grid)
					}
				}
				if found && dist < shortest {
					shortest = dist
				}
			}
		}
	}

	fmt.Printf("Part 2 answer = %d\n", shortest)
}

func addPos(x, y, dist int, elev rune, queue *[]position, visited map[vec]bool, grid [][]rune) {
	v := vec{x, y}
	if !visited[v] {
		el := grid[v.y][v.x]
		if el <= elev+1 {
			*queue = append(*queue, position{v, dist + 1})
			visited[v] = true
		}
	}
}
