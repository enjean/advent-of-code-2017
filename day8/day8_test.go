package main

import "testing"

func TestExecute(t *testing.T) {
	//b inc 5 if a > 1
	//	a inc 1 if b < 5
	//	c dec -10 if a >= 1
	//	c inc -20 if c == 10
	instructions := []Instruction{
		{"b", true, 5, Condition{"a", ">", 1}},
		{"a", true, 1, Condition{"b", "<", 5}},
		{"c", false, -10, Condition{"a", ">=", 1}},
		{"c", true, -20, Condition{"c", "==", 10}},
	}

	result, largestSeen := Execute(instructions)
	largest := LargestValue(result)
	if largest != 1 {
		t.Errorf("Wrong largest value, expected 1, was %d registers were %v", largest, result)
	}
	if largestSeen != 10 {
		t.Errorf("Wrong largest seen, expected 10, was %d", largestSeen)
	}
}
