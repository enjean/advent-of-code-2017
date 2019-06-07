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

func TestValidate2(t *testing.T) {
	//  abcde fghij is a valid passphrase.
	//	abcde xyz ecdab is not valid - the letters from the third word can be rearranged to form the first word.
	//	a ab abc abd abf abj is a valid passphrase, because all letters need to be used when forming another word.
	//	iiii oiii ooii oooi oooo is valid.
	//	oiii ioii iioi iiio is not valid - any of these words can be rearranged to form any other word.
	tests := []struct {
		passphrase string
		expected bool
	}{
		{"abcde fghij", true},
		{"abcde xyz ecdab",false},
		{"a ab abc abd abf abj",true},
		{"iiii oiii ooii oooi oooo",true},
		{"oiii ioii iioi iiio",false},
		{"aeggz eljcry buqdeog dvjzn ilvw arz vep kxdzm mvh szkf", true},
	}
	for _, test := range tests {
		result := Validate2(test.passphrase)
		if result != test.expected {
			t.Errorf("Validate2(%s) expected %v, got %v", test.passphrase, test.expected, result)
		}
	}
}
