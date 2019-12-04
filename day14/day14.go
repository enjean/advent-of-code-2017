package main

import (
	"fmt"
	"github.com/enjean/advent-of-code-2017/day10/knothash"
	"strconv"
)

func UsedSquares(key string) int {
	used := 0
	for row := 0; row < 128; row++ {
		keyForRow := fmt.Sprintf("%s-%d", key, row)
		knothashVal := knothash.CreateSparseHash(keyForRow).ToDenseHash()
		for _, r := range knothashVal {
			val, _ := strconv.ParseInt(string(r), 16, 32)
			asBinary := fmt.Sprintf("%04b", val)
			for _, c := range asBinary {
				if c == '1' {
					used++
				}
			}
		}
	}
	return used
}

func main() {
	part1 := UsedSquares("ugkiagan")
	fmt.Printf("Part 1: %d\n", part1)
}
