package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

type program []int

func (p program) Print(ip int) string {
	var b bytes.Buffer

	for i, ins := range p {
		if i == ip {
			b.WriteString("(")
		}
		b.WriteString(fmt.Sprintf("%d", ins))
		if i == ip {
			b.WriteString(")")
		}
		b.WriteString(" ")
	}
	return b.String()
}

func Execute(p program) int {
	steps := 0
	ip := 0
	for ip >= 0 && ip < len(p) {
		//fmt.Printf("%s\n", p.Print(ip))
		steps++
		jump := p[ip]
		p[ip]++
		ip += jump
	}

	//fmt.Printf("%s\n", p.Print(ip))
	return steps
}

func Execute2(p program) int {
	steps := 0
	ip := 0
	for ip >= 0 && ip < len(p) {
		//fmt.Printf("%s\n", p.Print(ip))
		steps++
		jump := p[ip]
		if p[ip] >= 3 {
			p[ip]--
		} else {
			p[ip]++
		}
		ip += jump
	}

	//fmt.Printf("%s\n", p.Print(ip))
	return steps
}

func main() {
	file, err := os.Open("day5/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var p []int
	for scanner.Scan() {
		lineAsInt, _ := strconv.Atoi(scanner.Text())
		p = append(p, lineAsInt)
	}

	program1 := make(program, len(p))
	copy(program1, p)
	steps := Execute(program1)
	fmt.Printf("Part 1: %d\n", steps)

	program2 := make(program, len(p))
	copy(program2, p)
	steps2 := Execute2(program2)
	fmt.Printf("Part 2: %d\n", steps2)
}
