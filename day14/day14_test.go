package main

import "testing"

func TestUsedSquares(t *testing.T) {
	result := UsedSquares("flqrgnkx")
	if result != 8108 {
		t.Errorf("Incorrect, expected %d, got %d", 8108, result)
	}
}
