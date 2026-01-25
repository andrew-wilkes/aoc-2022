package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
)

type point struct {
	x, y int
}

type sensor struct {
	pos, beacon point
}

type line struct {
	a, b point
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	sensors := []sensor{}
	r, _ := regexp.Compile("Sensor at x=([-0-9]+), y=([-0-9]+): closest beacon is at x=([-0-9]+), y=([-0-9]+)")
	for scanner.Scan() {
		line := scanner.Text()
		m := r.FindStringSubmatch(line)
		sensors = append(sensors, sensor{point{num(m[1]), num(m[2])}, point{num(m[3]), num(m[4])}})
	}

	// Part 1
	lines := []line{}
	beaconsOnLine := map[point]bool{}
	const loi = 2000000
	for _, s := range sensors {
		if s.beacon.y == loi {
			beaconsOnLine[s.beacon] = true
		}
		line, ok := coverage(s.pos, mdist(s.pos, s.beacon), loi)
		if ok {
			lines = append(lines, line)
		}
	}
	slices.SortFunc(lines, func(l, m line) int {
		return l.a.x - m.a.x
	})
	sum := -len(beaconsOnLine)
	xp := 0
	// Merge lines
	for i, l := range lines {
		if i > 0 {
			if l.a.x > xp {
				sum += l.b.x - l.a.x + 1
				xp = l.b.x
			} else if l.b.x > xp {
				sum += l.b.x - xp
				xp = l.b.x
			}
		} else {
			sum += l.b.x - l.a.x + 1
			xp = l.b.x
		}
	}
	fmt.Printf("Part 1 answer = %d\n", sum)

	// Part 2
	const limit = 4000000
	rows := map[int][]line{}
	for _, s := range sensors {
		dist := mdist(s.pos, s.beacon)
		for x := range dist + 1 {
			y := s.pos.y + x
			if y >= 0 && y <= limit {
				rows[y] = append(rows[y], line{point{s.pos.x - dist + x, y}, point{s.pos.x + dist - x, y}})
			}
			y = s.pos.y - x
			if y >= 0 && y <= limit && x > 0 {
				rows[y] = append(rows[y], line{point{s.pos.x - dist + x, y}, point{s.pos.x + dist - x, y}})
			}
		}
	}
	var db point
outer:
	for i := range rows {
		slices.SortFunc(rows[i], func(l, m line) int {
			return l.a.x - m.a.x
		})
		for j, l := range rows[i] {
			if j > 0 {
				if l.a.x > xp {
					if l.a.x-xp > 1 {
						if xp < 4000000 && xp > -1 && i <= 4000000 && i > -1 {
							db = point{xp + 1, i}
							break outer
						} else {
							break
						}
					}
					xp = l.b.x
				} else if l.b.x > xp {
					xp = l.b.x
				}
			} else {
				xp = l.b.x
			}
		}
	}
	fmt.Printf("Part 2 answer = %d\n", db.x*4000000+db.y)
}

func coverage(p point, dist, y int) (line line, ok bool) {
	yd := p.y - y
	if yd < 0 {
		yd = -yd
	}
	if yd <= dist {
		ok = true
		xd := dist - yd
		line.a = point{p.x - xd, y}
		line.b = point{p.x + xd, y}
	}
	return
}

func mdist(a, b point) int {
	x := a.x - b.x
	if x < 0 {
		x = -x
	}
	y := a.y - b.y
	if y < 0 {
		y = -y
	}
	return x + y
}

func num(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}
