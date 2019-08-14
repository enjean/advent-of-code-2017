package main

import (
	"fmt"
	"github.com/enjean/advent-of-code-2017/adventutil"
	"strings"
)

type Coordinate struct {
	q, r int
}

func (c *Coordinate) move(dir string) {
	switch dir {
	case "nw":c.q--
	case "n": c.r--
	case "ne":
		c.q++
		c.r--
	case "se": c.q++
	case "s": c.r++
	case "sw":
		c.q--
		c.r++
	}
}

func (c Coordinate) dist(other Coordinate) int {
	return (abs(c.q - other.q) + abs(c.q + c.r - other.q - other.r)+ abs(c.r - other.r)) / 2
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func StepsAway(path string) int {
	steps := strings.Split(path, ",")
	coord := Coordinate{0,0}
	for _, step := range steps {
		coord.move(step)
	}
	return coord.dist(Coordinate{0,0})
}

func MaxStepsAway(path string) int {
	steps := strings.Split(path, ",")
	coord := Coordinate{0,0}
	maxStepsAway := 0
	for _, step := range steps {
		coord.move(step)
		stepsAway := coord.dist(Coordinate{0,0})
		if stepsAway > maxStepsAway {
			maxStepsAway = stepsAway
		}
	}
	return maxStepsAway
}

func main() {
	lines := adventutil.Parse(11)
	path := lines[0]

	fmt.Printf("Day 11 Part 1: %d steps away\n", StepsAway(path))
	fmt.Printf("Day 11 Part 2: at most %d steps away\n", MaxStepsAway(path))
}
