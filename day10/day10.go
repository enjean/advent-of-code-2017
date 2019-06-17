package main

import (
	"fmt"
	"github.com/enjean/advent-of-code-2017/adventutil"
	"strconv"
	"strings"
)

type StringOfMarks []int

func New(numMarks int) StringOfMarks {
	var s StringOfMarks
	for i := 0; i < numMarks; i++ {
		s = append(s, i)
	}
	return s
}

func (s StringOfMarks) TieKnot(position, length int) {
	lengthOfString := len(s)
	for left, right := position, position+length-1; left < right; left, right = left+1, right-1 {
		s[left%lengthOfString], s[right%lengthOfString] = s[right%lengthOfString], s[left%lengthOfString]
	}
}

func (s StringOfMarks) TieKnots(lengths []int) {
	lengthOfString := len(s)
	currentPosition := 0
	skipSize := 0
	for _, length := range lengths {
		s.TieKnot(currentPosition, length)
		currentPosition = (currentPosition + length + skipSize) % lengthOfString
		skipSize++
		//fmt.Printf("%v %d %d\n", s, currentPosition, skipSize)
	}
}

func main() {
	lines := adventutil.Parse(10)
	var lengths []int
	for _, lengthString := range strings.Split(lines[0],",") {
		val, _ := strconv.Atoi(lengthString)
		lengths = append(lengths, val)
	}

	s := New(256)
	s.TieKnots(lengths)

	fmt.Printf("Part 1: %d", s[0] * s[1])
}
