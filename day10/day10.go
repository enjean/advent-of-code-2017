package main

import (
	"fmt"
	"github.com/enjean/advent-of-code-2017/adventutil"
	"github.com/enjean/advent-of-code-2017/day10/knothash"
	"strconv"
	"strings"
)

func main() {
	lines := adventutil.Parse(10)
	var lengths []int
	for _, lengthString := range strings.Split(lines[0], ",") {
		val, _ := strconv.Atoi(lengthString)
		lengths = append(lengths, val)
	}

	s := knothash.New(256)
	s.TieKnots(lengths, 0, 0)

	fmt.Printf("Part 1: %d\n", s[0]*s[1])

	sparseHash := knothash.CreateSparseHash(lines[0])
	knotHash := sparseHash.ToDenseHash()
	fmt.Printf("Part 2: %s\n", knotHash)
}
