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