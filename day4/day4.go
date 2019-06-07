package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Validate(passphrase string) bool {
	words := strings.Split(passphrase, " ")
	for w1 := 0; w1 < len(words) - 1; w1++ {
		for w2 := w1 + 1; w2 < len(words); w2++ {
			if words[w1] == words[w2] {
				return false
			}
		}
	}
	return true
}

func main() {
	file, err := os.Open("day4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var passphrases[]string
	for scanner.Scan() {
		passphrases = append(passphrases, scanner.Text())
	}

	valid := 0
	for _, passphrase := range passphrases {
		if Validate(passphrase) {
			valid++
		}
	}
	fmt.Printf("Part 1: %d\n", valid)
}
