package main

import "testing"

func TestValidate(t *testing.T) {
	//  aa bb cc dd ee is valid.
	//	aa bb cc dd aa is not valid - the word aa appears more than once.
	//	aa bb cc dd aaa is valid - aa and aaa count as different words.
	tests := []struct {
		passphrase string
		expected bool
	}{
		{"aa bb cc dd ee", true},
		{"aa bb cc dd aa",false},
		{"aa bb cc dd aaa",true},
	}
	for _, test := range tests {
		result := Validate(test.passphrase)
		if result != test.expected {
			t.Errorf("Validate(%s) expected %v, got %v", test.passphrase, test.expected, result)
		}
	}
}
