package main

import "testing"

func TestScore(t *testing.T) {
	//{}, score of 1.
	//{{{}}}, score of 1 + 2 + 3 = 6.
	//{{},{}}, score of 1 + 2 + 2 = 5.
	//{{{},{},{{}}}}, score of 1 + 2 + 3 + 3 + 3 + 4 = 16.
	//{<a>,<a>,<a>,<a>}, score of 1.
	//{{<ab>},{<ab>},{<ab>},{<ab>}}, score of 1 + 2 + 2 + 2 + 2 = 9.
	//{{<!!>},{<!!>},{<!!>},{<!!>}}, score of 1 + 2 + 2 + 2 + 2 = 9.
	//{{<a!>},{<a!>},{<a!>},{<ab>}}, score of 1 + 2 = 3
	tests := []struct {
		stream              string
		score, garbageCount int
	}{
		{"{}", 1, 0},
		{"{{{}}}", 6, 0},
		{"{{},{}}", 5, 0},
		{"{{{},{},{{}}}}", 16, 0},
		{"{<a>,<a>,<a>,<a>}", 1, 4},
		{"{{<ab>},{<ab>},{<ab>},{<ab>}}", 9, 8},
		{"{{<!!>},{<!!>},{<!!>},{<!!>}}", 9, 0},
		{"{{<a!>},{<a!>},{<a!>},{<ab>}}", 3, 17},
	}

	for _, test := range tests {
		score, garbageCount := Score(test.stream)
		if score != test.score {
			t.Errorf("Score(%s) expected %d got %d", test.stream, test.score, score)
		}
		if garbageCount != test.garbageCount {
			t.Errorf("GarbageCount(%s) expected %d got %d", test.stream, test.garbageCount, garbageCount)
		}
	}
}
