package main

import (
	"fmt"
	"io/ioutil"
)

func Captcha(input string) int {
	sum := 0

	for i := 0; i < len(input); i++ {
		j := i + 1
		if i == len(input) - 1 {
			j = 0
		}
		if input[i] == input[j] {
			val := int(input[i]) - '0'
			sum += val
		}
	}


	return sum
}

func Captcha2(input string) int {
	sum := 0

	half := len(input) / 2
	for i := 0; i < len(input); i++ {
		j := (i + half) % len(input)
		if input[i] == input[j] {
			val := int(input[i]) - '0'
			sum += val
		}
	}


	return sum
}

func main() {
	bytes, _ := ioutil.ReadFile("day1/input.txt")
	input := string(bytes)

	captcha := Captcha(input)
	fmt.Printf("Part 1: %d\n", captcha)

	captcha2 := Captcha2(input)
	fmt.Printf("Part 2: %d\n", captcha2)
}
