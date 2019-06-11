package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Memory struct {
	banks []int
}

func (mem Memory) Reallocate() {
	numBanks := len(mem.banks)

	bankToRedistribute, maxBlocks := -1, -1
	for i, blocks := range mem.banks {
		if blocks > maxBlocks {
			bankToRedistribute = i
			maxBlocks = blocks
		}
	}

	blocksToRedistribute := mem.banks[bankToRedistribute]
	mem.banks[bankToRedistribute] = 0
	bankI := bankToRedistribute
	for blocksToRedistribute > 0 {
		bankI = (bankI + 1) % numBanks
		mem.banks[bankI]++
		blocksToRedistribute--
	}
}

func equals(a, b []int) bool {
	for i, val := range a {
		if b[i] != val {
			return false
		}
	}
	return true
}

func CyclesBeforeLoop(m Memory) int {
	var pastConfigurations [][]int

	cycleCount := 0
	foundRepeat := false
	for !foundRepeat {
		old := make([]int, len(m.banks))
		copy(old, m.banks)
		pastConfigurations = append(pastConfigurations, old)
		cycleCount++
		m.Reallocate()
		for _, past := range pastConfigurations {
			if equals(m.banks, past) {
				foundRepeat = true
				break
			}
		}
	}
	return cycleCount
}

func main() {
	file, err := os.Open("day6/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	var banks []int
	for _, blocks := range strings.Split(line, "\t") {
		val, _ := strconv.Atoi(blocks)
		banks = append(banks, val)
	}

	m := Memory{banks:banks}

	part1 := CyclesBeforeLoop(m)
	fmt.Printf("Part 1: %d\n", part1)

	part2 := CyclesBeforeLoop(m)
	fmt.Printf("Part 2: %d\n", part2)
}
