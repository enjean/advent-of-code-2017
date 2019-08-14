package main

import "testing"

func TestStepsAway(t *testing.T) {
	//ne,ne,ne is 3 steps away.
	//	ne,ne,sw,sw is 0 steps away (back where you started).
	//ne,ne,s,s is 2 steps away (se,se).
	//	se,sw,se,sw,sw is 3 steps away (s,s,sw).
	tests := []struct {
		path string
		expected int
	}{
		{"ne,ne,ne", 3},
		{"ne,ne,sw,sw", 0},
		{"ne,ne,s,s", 2},
		{"se,sw,se,sw,sw", 3},
	}
	for _, test := range tests {
		result := StepsAway(test.path)
		if result != test.expected {
			t.Errorf("StepsAway(%v) expected %d got %d", test.path, test.expected, result)
		}
	}
}
