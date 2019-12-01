package main

import "testing"

func TestProgramsInGroup(t *testing.T) {
	//0 <-> 2
	//1 <-> 1
	//2 <-> 0, 3, 4
	//3 <-> 2, 4
	//4 <-> 2, 3, 6
	//5 <-> 6
	//6 <-> 4, 5
	input := map[int][]int{
		0: {2},
		1: {1},
		2: {0, 3, 4},
		3: {2, 4},
		4: {2, 3, 6},
		5: {6},
		6: {4, 5},
	}
	result := ProgramsInGroup(0, input)
	if result != 6 {
		t.Errorf("Expected 6 got %d", result)
	}
}

func TestNumberOfGroups(t *testing.T) {
	input := map[int][]int{
		0: {2},
		1: {1},
		2: {0, 3, 4},
		3: {2, 4},
		4: {2, 3, 6},
		5: {6},
		6: {4, 5},
	}
	result := NumberOfGroups(input)
	if result != 2 {
		t.Errorf("Expected 2 got %d", result)
	}
}
