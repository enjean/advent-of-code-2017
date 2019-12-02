package main

import "testing"

func TestTripSeverity(t *testing.T) {
	//0: 3
	//1: 2
	//4: 4
	//6: 4
	input := []Layer{
		{0, 3},
		{1, 2},
		{4, 4},
		{6, 4},
	}

	result := TripSeverity(input)

	if result != 24 {
		t.Errorf("Expected 24, was %d", result)
	}
}
