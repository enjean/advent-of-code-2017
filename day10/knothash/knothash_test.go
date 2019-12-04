package knothash

import (
	"reflect"
	"testing"
)

func TestStringOfMarks_TieKnot(t *testing.T) {
	tests := []struct {
		before           StringOfMarks
		position, length int
		after            StringOfMarks
	}{
		{StringOfMarks{0, 1, 2, 3, 4}, 0, 3, StringOfMarks{2, 1, 0, 3, 4}},
		{StringOfMarks{2, 1, 0, 3, 4}, 3, 4, StringOfMarks{4, 3, 0, 1, 2}},
		{StringOfMarks{4, 3, 0, 1, 2}, 3, 1, StringOfMarks{4, 3, 0, 1, 2}},
		{StringOfMarks{4, 3, 0, 1, 2}, 1, 5, StringOfMarks{3, 4, 2, 1, 0}},
	}
	for _, test := range tests {
		b := make(StringOfMarks, len(test.before))
		copy(b, test.before)
		b.TieKnot(test.position, test.length)
		if !reflect.DeepEqual(b, test.after) {
			t.Errorf("%v.TieKnot(%d,%d) expected %v got %v", test.before, test.position, test.length, test.after, b)
		}
	}
}

func TestStringOfMarks_TieKnots(t *testing.T) {
	s := StringOfMarks{0, 1, 2, 3, 4}
	s.TieKnots([]int{3, 4, 1, 5}, 0, 0)
	if !reflect.DeepEqual(s, StringOfMarks{3, 4, 2, 1, 0}) {
		t.Errorf("TieKnots expected [3,4,2,1,0] got %v", s)
	}
}

func TestToLengths(t *testing.T) {
	result := ToLengths("1,2,3")
	expected := []int{49, 44, 50, 44, 51, 17, 31, 73, 47, 23}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Lengths incorrect, got %v", result)
	}
}

func TestKnotHash(t *testing.T) {
	//The empty string becomes a2582a3a0e66e6e86e3812dcb672a272.
	//	AoC 2017 becomes 33efeb34ea91902bb2f59c9920caa6cd.
	//1,2,3 becomes 3efbe78a8d82f29979031a4aa0b16a9d.
	//1,2,4 becomes 63960835bcdc130f0b66d7ff4f6a5a8e.
	tests := []struct {
		input, expected string
	}{
		{"", "a2582a3a0e66e6e86e3812dcb672a272"},
		{"AoC 2017", "33efeb34ea91902bb2f59c9920caa6cd"},
		{"1,2,3", "3efbe78a8d82f29979031a4aa0b16a9d"},
		{"1,2,4", "63960835bcdc130f0b66d7ff4f6a5a8e"},
	}

	for _, test := range tests {
		sparseHash := CreateSparseHash(test.input)
		result := sparseHash.ToDenseHash()
		if result != test.expected {
			t.Errorf("KnotHash(%s) expected %s got %s", test.input, test.expected, result)
		}
	}
}
