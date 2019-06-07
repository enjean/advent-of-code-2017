package main

import "testing"

func TestCaptcha(t *testing.T) {
	tests := []struct {
		input string
		expected int
	}{
		{"1122", 3},
		{"1111", 4},
		{"1234", 0},
		{"91212129", 9},
	}

	for _, test := range tests {
		result := Captcha(test.input)
		if result != test.expected {
			t.Errorf("Captcha(%s) expected %d, got %d", test.input, test.expected, result)
		}
	}
}

func TestCaptcha2(t *testing.T) {
	//1212 produces 6: the list contains 4 items, and all four digits match the digit 2 items ahead.
	//1221 produces 0, because every comparison is between a 1 and a 2.
	//123425 produces 4, because both 2s match each other, but no other digit has a match.
	//123123 produces 12.
	//12131415 produces 4.
	tests := []struct {
		input string
		expected int
	}{
		{"1212", 6},
		{"1221", 0},
		{"123425", 4},
		{"123123", 12},
		{"12131415", 4},
	}

	for _, test := range tests {
		result := Captcha2(test.input)
		if result != test.expected {
			t.Errorf("Captcha(%s) expected %d, got %d", test.input, test.expected, result)
		}
	}
}