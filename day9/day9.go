package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Score(stream string) int {
	score := 0
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
	return score
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

	score := Score(stream)
	fmt.Printf("Part 1: %d\n", score)
}
