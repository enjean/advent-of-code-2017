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

func TripSeverity(firewall []Layer) int {

	severity := 0
	for _, l := range firewall {
		frequency := (l.scannerRange - 1) * 2
		if l.depth%frequency == 0 {
			severity += l.depth * l.scannerRange
		}
	}

	return severity
}

func FindMinDelayForCleanTrip(firewall []Layer) int {
	frequencies := make(map[int]int)
	for _, l := range firewall {
		frequencies[l.depth] = (l.scannerRange - 1) * 2
	}
	delay := 0
	for {
		if isRunClean(frequencies, delay) {
			return delay
		}
		delay++
	}
}

func isRunClean(frequencies map[int]int, delay int) bool {
	for depth, frequency := range frequencies {
		secondHittingDepth := depth + delay
		if secondHittingDepth%frequency == 0 {
			return false
		}
	}
	return true
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

	severity := TripSeverity(firewall)
	fmt.Printf("Part 1: severity %d\n", severity)

	fmt.Printf("Part 2: delay %d\n", FindMinDelayForCleanTrip(firewall))
}
