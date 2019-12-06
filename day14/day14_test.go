package main

import "testing"

func TestUsedSquares(t *testing.T) {
	result := UsedSquares("flqrgnkx")
	if result != 8108 {
		t.Errorf("Incorrect, expected %d, got %d", 8108, result)
	}
}

func TestCountRegions(t *testing.T) {
	result := CountRegions("flqrgnkx")
	if result != 1242 {
		t.Errorf("Incorrect, expected %d, got %d", 1242, result)
	}
}
