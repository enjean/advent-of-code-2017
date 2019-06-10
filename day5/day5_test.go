package main

import "testing"

func TestExecute(t *testing.T) {
	program := []int{
		0,
		3,
		0,
		1,
		-3,
	}

	steps := Execute(program)

	if steps != 5 {
		t.Errorf("Expected 5 to Execute, was %d", steps)
	}
}
