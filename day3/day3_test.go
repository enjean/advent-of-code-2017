package main

import "testing"

func TestStepsToPort(t *testing.T) {
	//Data from square 1 is carried 0 steps, since it's at the access port.
	//	Data from square 12 is carried 3 steps, such as: down, left, left.
	//Data from square 23 is carried only 2 steps: up twice.
	//	Data from square 1024 must be carried 31 steps.
	tests := []struct {
		fromSquare int
		expected int
	}{
		{1, 0},
		{12,3},
		{23,2},
		{1024, 31},
	}
	for _, test := range tests {
		result := StepsToPort(test.fromSquare)
		if result != test.expected {
			t.Errorf("StepsToPort(%d) expected %d, got %d", test.fromSquare, test.expected, result)
		}
	}
}
