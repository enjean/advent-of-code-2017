package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func Validate(passphrase string) bool {
	return doValidation(passphrase, func(w1, w2 string) bool {
		return w1 == w2
	})
}

func Validate2(passphrase string) bool {
	return doValidation(passphrase, func(w1, w2 string) bool {
		if len(w1) != len(w2) {
			return false
		}
		runesIn1 := []rune(w1)
		sort.Slice(runesIn1, func(i, j int) bool {
			return runesIn1[i] < runesIn1[j]
		})
		runesIn2 := []rune(w2)
		sort.Slice(runesIn2, func(i, j int) bool {
			return runesIn2[i] < runesIn2[j]
		})
		for i, runeIn1 := range runesIn1 {
			if runesIn2[i] != runeIn1 {
				return false
			}
		}
		return true
	})
}

func doValidation(passphrase string, wordComparator func(string, string) bool) bool {
	words := strings.Split(passphrase, " ")
	for w1 := 0; w1 < len(words) - 1; w1++ {
		for w2 := w1 + 1; w2 < len(words); w2++ {
			if wordComparator(words[w1], words[w2]) {
				fmt.Printf("%s invalid: %s and %s are matches\n", passphrase, words[w1], words[w2])
				return false
			}
		}
	}
	return true
}



func CountValid(passphrases []string, f func(string) bool) int {
	valid := 0
	for _, passphrase := range passphrases {
		if f(passphrase) {
			valid++
		}
	}
	return valid
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

	fmt.Printf("Part 1: %d\n", CountValid(passphrases, Validate))
	fmt.Printf("Part 2: %d\n", CountValid(passphrases, Validate2))
}
