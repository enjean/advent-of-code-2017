package main

import (
	"bytes"
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

func (s StringOfMarks) TieKnots(lengths []int, currentPosition, skipSize int) (int, int) {
	lengthOfString := len(s)
	for _, length := range lengths {
		s.TieKnot(currentPosition, length)
		currentPosition = (currentPosition + length + skipSize) % lengthOfString
		skipSize++
		//fmt.Printf("%v %d %d\n", s, currentPosition, skipSize)
	}
	return currentPosition, skipSize
}

func ToLengths(input string) []int {
	var lengths []int
	for _, r := range input {
		lengths = append(lengths, int(r))
	}
	lengths = append(lengths, 17, 31, 73, 47, 23)
	return lengths
}

type SparseHash [256]int

func CreateSparseHash(input string) SparseHash {
	s := New(256)
	lengths := ToLengths(input)
	current, skip := 0, 0
	for round := 0; round < 64; round++ {
		current, skip = s.TieKnots(lengths, current, skip)
	}

	var sparseHash SparseHash
	copy(sparseHash[:], s)
	return sparseHash
}

func (sh SparseHash) ToDenseHash() string {
	var denseHash [16]int
	for i := 0; i < 16; i++ {
		nums := sh[i*16 : (i+1)*16]
		denseHash[i] = nums[0]
		for j := 1; j < 16; j++ {
			denseHash[i] ^= nums[j]
		}
	}

	var b bytes.Buffer
	for _, val := range denseHash {
		b.WriteString(fmt.Sprintf("%02x", val))
	}
	return b.String()
}

func main() {
	lines := adventutil.Parse(10)
	var lengths []int
	for _, lengthString := range strings.Split(lines[0], ",") {
		val, _ := strconv.Atoi(lengthString)
		lengths = append(lengths, val)
	}

	s := New(256)
	s.TieKnots(lengths, 0, 0)

	fmt.Printf("Part 1: %d\n", s[0]*s[1])

	sparseHash := CreateSparseHash(lines[0])
	knotHash := sparseHash.ToDenseHash()
	fmt.Printf("Part 2: %s\n", knotHash)
}
