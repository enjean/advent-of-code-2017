package main

import (
	"fmt"
	"github.com/enjean/advent-of-code-2017/day10/knothash"
	"strconv"
)

func UsedSquares(key string) int {
	grid := buildGrid(key)
	used := 0
	for row := 0; row < 128; row++ {
		for col := 0; col < 128; col++ {
			if grid[row][col] {
				used++
			}
		}
	}
	return used
}

type point struct {
	row, col int
}

func (p point) neighbors() []point {
	var neighbors []point
	if p.row != 0 {
		neighbors = append(neighbors, point{p.row - 1, p.col})
	}
	if p.col != 0 {
		neighbors = append(neighbors, point{p.row, p.col - 1})
	}
	if p.row != 127 {
		neighbors = append(neighbors, point{p.row + 1, p.col})
	}
	if p.col != 127 {
		neighbors = append(neighbors, point{p.row, p.col + 1})
	}
	return neighbors
}

func CountRegions(key string) int {
	grid := buildGrid(key)
	visited := make(map[point]bool)
	regions := 0
	for row := 0; row < 128; row++ {
		for col := 0; col < 128; col++ {
			if !grid[row][col] {
				continue
			}
			gridPoint := point{row, col}
			if visited[gridPoint] {
				continue
			}
			regions++
			toVisit := []point{gridPoint}
			for len(toVisit) > 0 {
				pointToEvaluate := toVisit[0]
				toVisit = toVisit[1:]
				visited[pointToEvaluate] = true
				for _, neighbor := range pointToEvaluate.neighbors() {
					if visited[neighbor] || !grid[neighbor.row][neighbor.col] {
						continue
					}
					toVisit = append(toVisit, neighbor)
				}
			}
		}
	}
	return regions
}

func buildGrid(key string) [128][128]bool {
	var grid [128][128]bool
	for row := 0; row < 128; row++ {
		keyForRow := fmt.Sprintf("%s-%d", key, row)
		knothashVal := knothash.CreateSparseHash(keyForRow).ToDenseHash()
		for charIndex, r := range knothashVal {
			val, _ := strconv.ParseInt(string(r), 16, 32)
			asBinary := fmt.Sprintf("%04b", val)
			for digit, char := range asBinary {
				grid[row][4*charIndex+digit] = char == '1'
			}
		}
	}
	return grid
}

func main() {
	key := "ugkiagan"
	part1 := UsedSquares(key)
	fmt.Printf("Part 1: %d\n", part1)

	fmt.Printf("Part 2: %d\n", CountRegions(key))
}
