package main

import "testing"

func TestChecksum(t *testing.T) {
	//For example, given the following spreadsheet:
	//
	//5 1 9 5
	//7 5 3
	//2 4 6 8
	//The first row's largest and smallest values are 9 and 1, and their difference is 8.
	//The second row's largest and smallest values are 7 and 3, and their difference is 4.
	//The third row's difference is 6.
	//In this example, the spreadsheet's checksum would be 8 + 4 + 6 = 18.
	example := [][]int{
		{5, 1, 9, 5},
		{7, 5, 3},
		{2, 4, 6, 8},
	}

	checksum := Checksum(example)
	if checksum != 18 {
		t.Errorf("Checksum(%v) expected %d got %d", example, 18, checksum)
	}
}
