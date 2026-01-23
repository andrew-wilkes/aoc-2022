package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type aline struct {
	a, b, x, y int
}

type point struct {
	x, y int
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	hlines, vlines := []aline{}, []aline{}
	maxDepth := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		points := strings.Split(line, " -> ")
		var x, y, a, b int
		for i, p := range points {
			if i > 0 {
				a = x
				b = y
			}
			xy := strings.Split(p, ",")
			x, _ = strconv.Atoi(xy[0])
			y, _ = strconv.Atoi(xy[1])
			if y > maxDepth {
				maxDepth = y
			}
			if i > 0 {
				if a == x {
					if b > y {
						vlines = append(vlines, aline{a, y, x, b})
					} else {
						vlines = append(vlines, aline{a, b, x, y})
					}

				} else {
					if a > x {
						hlines = append(hlines, aline{x, b, a, y})
					} else {
						hlines = append(hlines, aline{a, b, x, y})
					}

				}
			}
		}
	}

	numUnits := 0
	sand := map[point]bool{}
	filling := true
	for filling {
		pos := point{500, 0}
		for {
			pos.y++
			if pos.y > maxDepth {
				filling = false
				break
			}
			if !sand[pos] && !isPointOnLine(pos.x, pos.y, hlines, vlines) {
				continue
			}
			pos.x--
			if !sand[pos] && !isPointOnLine(pos.x, pos.y, hlines, vlines) {
				continue
			}
			pos.x += 2
			if !sand[pos] && !isPointOnLine(pos.x, pos.y, hlines, vlines) {
				continue
			}
			pos.x--
			pos.y--
			if pos.y > 0 {
				numUnits++
				sand[pos] = true
				break
			}
		}
	}

	fmt.Printf("Part 1 answer = %d\n", numUnits)

	// Part 2

	numUnits = 0
	sand = map[point]bool{}
	filling = true
	maxDepth += 2
	for filling {
		pos := point{500, 0}
		for {
			pos.y++
			if pos.y < maxDepth {
				if !sand[pos] && !isPointOnLine(pos.x, pos.y, hlines, vlines) {
					continue
				}
				pos.x--
				if !sand[pos] && !isPointOnLine(pos.x, pos.y, hlines, vlines) {
					continue
				}
				pos.x += 2
				if !sand[pos] && !isPointOnLine(pos.x, pos.y, hlines, vlines) {
					continue
				}
				pos.x--
			}
			pos.y--
			numUnits++
			sand[pos] = true
			if pos.y == 0 {
				filling = false
			}
			break
		}
	}
	fmt.Printf("Part 2 answer = %d\n", numUnits)
}

func isPointOnLine(x, y int, hlines, vlines []aline) bool {
	for _, h := range hlines {
		if h.y == y {
			if x >= h.a && x <= h.x {
				return true
			}
		}
	}
	for _, v := range vlines {
		if v.x == x {
			if y >= v.b && y <= v.y {
				return true
			}
		}
	}
	return false
}
