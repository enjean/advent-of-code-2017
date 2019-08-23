package main

import (
	"fmt"
	"github.com/enjean/advent-of-code-2017/adventutil"
	"strconv"
	"strings"
)

func ProgramsInGroup(group int, connections map[int][]int) int {
	visited := map[int]bool {group:true}
	toVisit := []int{group}
	for len(toVisit) > 0 {
		program := toVisit[0]
		toVisit = toVisit[1:]
		visited[program] = true
		for _, neighbor := range connections[program] {
			if visited[neighbor] {
				continue
			}
			toVisit = append(toVisit, neighbor)
		}
	}
	return len(visited)
}

func main() {
	lines := adventutil.Parse(12)
	input := make(map[int][]int)
	for _, line := range lines {
		splitLine := strings.Split(line, " <-> ")
		program, _ := strconv.Atoi(splitLine[0])
		neighborStrings := splitLine[1]
		var neighbors []int
		for _, neighborStr := range strings.Split(neighborStrings, ", ") {
			neighbor, _ := strconv.Atoi(neighborStr)
			neighbors = append(neighbors, neighbor)
		}
		input[program] = neighbors
	}

	groupSize := ProgramsInGroup(0, input)
	fmt.Printf("Part 1: %d in group\n", groupSize)
}
