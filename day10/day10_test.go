package main

import (
	"reflect"
	"testing"
)

func TestStringOfMarks_TieKnot(t *testing.T) {
	tests := []struct{
		before StringOfMarks
		position, length int
		after StringOfMarks
	}{
		{StringOfMarks{0,1,2,3,4},0, 3, StringOfMarks{2,1,0,3,4}},
		{StringOfMarks{2,1,0,3,4},3, 4, StringOfMarks{4,3,0,1,2}},
		{StringOfMarks{4,3,0,1,2},3, 1, StringOfMarks{4,3,0,1,2}},
		{StringOfMarks{4,3,0,1,2},1, 5, StringOfMarks{3,4,2,1,0}},
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
	s := StringOfMarks{0,1,2,3,4}
	s.TieKnots([]int{3,4,1,5})
	if !reflect.DeepEqual(s, StringOfMarks{3,4,2,1,0}) {
		t.Errorf("TieKnots expected [3,4,2,1,0] got %v", s)
	}
}
