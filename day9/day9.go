package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Score(stream string) (int, int) {
	score, garbageCount := 0, 0
	currentLevel := 0
	ignoreNext, inGarbage := false, false
	for _, runeVal := range stream {
		if ignoreNext {
			ignoreNext = false
			continue
		}
		if inGarbage {
			switch runeVal {
			case '!':
				ignoreNext = true
			case '>':
				inGarbage = false
			default:
				garbageCount++
			}
		} else {
			switch runeVal {
			case '{':
				currentLevel++
			case '}':
				score += currentLevel
				currentLevel--
			case '<':
				inGarbage = true
			}
		}
	}
	return score, garbageCount
}

func main() {
	file, err := os.Open("day9/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	stream := scanner.Text()

	score, garbageCount := Score(stream)
	fmt.Printf("Part 1: %d\n", score)
	fmt.Printf("Part 2: %d\n", garbageCount)
}
