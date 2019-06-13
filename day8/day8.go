package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Instruction struct {
	registerToModify string
	increase         bool
	amount           int
	condition        Condition
}

type Condition struct {
	register string
	op       string
	val      int
}

func (c Condition) evaluate(registers map[string]int) bool {
	registerVal := registers[c.register]
	switch c.op {
	case ">":
		return registerVal > c.val
	case "<":
		return registerVal < c.val
	case ">=":
		return registerVal >= c.val
	case "<=":
		return registerVal <= c.val
	case "==":
		return registerVal == c.val
	case "!=":
		return registerVal != c.val
	}
	panic(fmt.Errorf("unknown operation %s", c.op))
}

func Execute(instructions []Instruction) map[string]int {
	registers := make(map[string]int)

	for _, instruction := range instructions {
		if !instruction.condition.evaluate(registers) {
			continue
		}
		if instruction.increase {
			registers[instruction.registerToModify] += instruction.amount
		} else {
			registers[instruction.registerToModify] -= instruction.amount
		}

	}

	return registers
}

func LargestValue(registers map[string]int) int {
	max := -int(^uint(0) >> 1) - 1
	for _, rValue := range registers {
		if rValue > max {
			max = rValue
		}
	}
	return max
}

//b inc 5 if a > 1
var instructionRegex = regexp.MustCompile(`([a-z]+) (inc|dec) (-?\d+) if ([a-z]+) (.*) (-?\d+)`)
func main() {
	file, err := os.Open("day8/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var instructions []Instruction
	for scanner.Scan() {
		line := scanner.Text()
		instructionMatches := instructionRegex.FindStringSubmatch(line)
		amount, _ := strconv.Atoi(instructionMatches[3])
		conditionVal, _ := strconv.Atoi(instructionMatches[6])
		instruction := Instruction{
			registerToModify: instructionMatches[1],
			increase: instructionMatches[2] == "inc",
			amount: amount,
			condition:Condition{
				register:instructionMatches[4],
				op: instructionMatches[5],
				val: conditionVal,
			},
		}
		instructions = append(instructions, instruction)
	}

	part1 := LargestValue(Execute(instructions))
	fmt.Printf("Part 1: %d\n", part1)
}
