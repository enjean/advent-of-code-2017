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
		stream string
		score  int
	}{
		{"{}", 1},
		{"{{{}}}", 6},
		{"{{},{}}", 5},
		{"{{{},{},{{}}}}", 16},
		{"{<a>,<a>,<a>,<a>}", 1},
		{"{{<ab>},{<ab>},{<ab>},{<ab>}}", 9},
		{"{{<!!>},{<!!>},{<!!>},{<!!>}}", 9},
		{"{{<a!>},{<a!>},{<a!>},{<ab>}}", 3},
	}

	for _, test := range tests {
		result := Score(test.stream)
		if result != test.score {
			t.Errorf("Score(%s) expected %d got %d", test.stream, test.score, result)
		}
	}
}
