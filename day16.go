package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type valve struct {
	flowRate int
	tunnels  []string // IDs of connected nodes
	open     *bool
}

type cave struct {
	valveId          string
	minute           int
	pressRelief      int
	totalPressRelief int
}

func main() {
	file, _ := os.Open("example.txt")
	defer file.Close()

	valves := map[string]valve{}
	scanner := bufio.NewScanner(file)
	r, _ := regexp.Compile("Valve ([A-Z]+) has flow rate=([0-9]+); tunnels? leads? to valves? (.+)")
	for scanner.Scan() {
		line := scanner.Text()
		m := r.FindStringSubmatch(line)
		valves[m[1]] = valve{num(m[2]), strings.Split(m[3], ", "), new(bool)}
	}

	// Part 1

	// Build a distance map for distances between nodes that have working valves including AA
	distances := map[string]map[string]int{}

	for id, n := range valves {
		if n.flowRate > 0 || id == "AA" {
			addDistances(id, distances, valves)
		}
	}

	maxPressureRelief := dfsTunnels(distances, valves)

	//	testPath := []string{"AA", "DD", "CC", "BB", "AA", "II", "JJ", "II", "AA", "DD", "EE", "FF", "GG", "HH", "GG", "FF", "EE", "DD", "CC"}

	fmt.Printf("Part 1 answer = %d\n", maxPressureRelief)
}

func dfsTunnels(distances map[string]map[string]int, valves map[string]valve) int {
	stack := []cave{{valveId: "AA", minute: 1}}
	maxPressureRelief := 0
	for len(stack) > 0 {
		idx := len(stack) - 1
		c := stack[idx] // pop cave status off stack
		stack = stack[:idx]

		if !*valves[c.valveId].open {
			*valves[c.valveId].open = true
			c.pressRelief += valves[c.valveId].flowRate
		}
		atEnd := true
		if c.minute < 30 {
			for id, d := range distances[c.valveId] {
				if c.minute+d <= 30 {
					atEnd = false
					stack = append(stack, cave{id, c.minute + d, c.pressRelief, c.totalPressRelief + d*c.pressRelief})
				}
			}
		}
		if atEnd {
			remainingMinutes := 30 - c.minute
			c.totalPressRelief += c.pressRelief * remainingMinutes
			if c.totalPressRelief > maxPressureRelief {
				maxPressureRelief = c.totalPressRelief
				fmt.Println(maxPressureRelief)
			}
		}
	}
	return maxPressureRelief
}

func addDistances(id string, distances map[string]map[string]int, valves map[string]valve) {
	distances[id] = map[string]int{}
	for nid, valve := range valves {
		if nid != id && valve.flowRate > 0 {
			if distances[nid] != nil && distances[nid][id] != 0 {
				distances[id][nid] = distances[nid][id]
				continue
			}
			distances[id][nid] = getDistanceBFS(id, nid, valves)
		}
	}
}

func getDistanceBFS(a, b string, valves map[string]valve) int {
	q := []string{a}
	dist := map[string]int{a: 0}
	for len(q) > 0 {
		nid := q[0]
		d := dist[nid]
		if nid == b {
			return d
		}
		q = q[1:]
		d++
		for _, id := range valves[nid].tunnels {
			q = append(q, id)
			dist[id] = d
		}
	}
	return 0
}

// A simple utility function to simplify code above.
func num(s string) int {
	n, _ := strconv.Atoi(s) // Assume that there will be no error or else we should get a panic.
	return n
}
