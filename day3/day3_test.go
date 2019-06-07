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

func TestStressTest(t *testing.T) {
	//Once a square is written, its value does not change. Therefore, the first few squares would receive the following values:
	//
	//147  142  133  122   59
	//304    5    4    2   57
	//330   10    1    1   54
	//351   11   23   25   26
	//362  747  806--->   ...
	tests := []struct {
		limit int
		expected int
	}{
		{1, 2},
		{12,23},
		{55,57},
		{800, 806},
	}
	for _, test := range tests {
		result := StressTest(test.limit)
		if result != test.expected {
			t.Errorf("StressTest(%d) expected %d, got %d", test.limit, test.expected, result)
		}
	}
}
