package main

import "testing"

func TestFindRoot(t *testing.T) {
	//pbga (66)
	//xhth (57)
	//ebii (61)
	//havc (66)
	//ktlj (57)
	//fwft (72) -> ktlj, cntj, xhth
	//qoyq (66)
	//padx (45) -> pbga, havc, qoyq
	//tknk (41) -> ugml, padx, fwft
	//jptl (61)
	//ugml (68) -> gyxo, ebii, jptl
	//gyxo (61)
	//cntj (57)
	input := []programSpec{
		{"pbga", nil},
		{"xhth", nil},
		{"ebii", nil},
		{"havc", nil},
		{"ktlj", nil},
		{"fwft", []string{"ktlj", "cntj", "xhth"}},
		{"qoyq", nil},
		{"padx", []string{"pbga", "havc", "qoyq"}},
		{"tknk", []string{"ugml", "padx", "fwft"}},
		{"jptl", nil},
		{"tknk", []string{"gyxo", "ebii", "jptl"}},
		{"gyxo", nil},
		{"cntj", nil},
	}

	root := FindRoot(input)
	if root != "tknk" {
		t.Errorf("Expected root of tknk, got %s", root)
	}
}