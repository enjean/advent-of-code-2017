package main

import (
	"fmt"
	"github.com/enjean/advent-of-code-2017/adventutil"
	"strconv"
	"strings"
)

type Layer struct {
	depth        int
	scannerRange int
}

type scannerPosition struct {
	position    int
	maxRange    int
	isGoingDown bool
}

func TripSeverity(firewall []Layer) int {
	severity := 0
	maxDepth := firewall[len(firewall)-1].depth

	scanners := make(map[int]*scannerPosition)
	for _, l := range firewall {
		scanners[l.depth] = &scannerPosition{
			position:    0,
			maxRange:    l.scannerRange,
			isGoingDown: true,
		}
	}

	for depth := 0; depth <= maxDepth; depth++ {
		scannerAtDepth, ok := scanners[depth]
		if ok && scannerAtDepth.position == 0 {
			//fmt.Printf("Intersected at depth %d\n", depth)
			severity += depth * scanners[depth].maxRange
		}
		//fmt.Printf("Second %d\n", depth)
		for _, sc := range scanners {
			//fmt.Printf("  %d Before: %d %v\n", d, sc.position, sc.isGoingDown)
			if sc.position == 0 {
				sc.isGoingDown = true
			} else if sc.position == sc.maxRange-1 {
				sc.isGoingDown = false
			}
			if sc.isGoingDown {
				sc.position++
			} else {
				sc.position--
			}
			//fmt.Printf("  %d After: %d %v\n", d, sc.position, sc.isGoingDown)
		}
	}
	return severity
}

func main() {
	var firewall []Layer
	lines := adventutil.Parse(13)
	for _, line := range lines {
		lineParts := strings.Split(line, ": ")
		depth, _ := strconv.Atoi(lineParts[0])
		scannerRange, _ := strconv.Atoi(lineParts[1])
		firewall = append(firewall, Layer{
			depth:        depth,
			scannerRange: scannerRange,
		})
	}

	fmt.Printf("Part 1: severity %d\n", TripSeverity(firewall))
}
