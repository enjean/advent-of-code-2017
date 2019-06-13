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
		{"pbga", 66,nil},
		{"xhth", 57,nil},
		{"ebii", 61,nil},
		{"havc", 66,nil},
		{"ktlj", 57,nil},
		{"fwft", 72,[]string{"ktlj", "cntj", "xhth"}},
		{"qoyq", 66,nil},
		{"padx", 45,[]string{"pbga", "havc", "qoyq"}},
		{"tknk", 41,[]string{"ugml", "padx", "fwft"}},
		{"jptl", 61,nil},
		{"ugml", 41,[]string{"gyxo", "ebii", "jptl"}},
		{"gyxo", 61,nil},
		{"cntj", 57,nil},
	}

	root := FindRoot(input)
	if root != "tknk" {
		t.Errorf("Expected root of tknk, got %s", root)
	}
}

func TestFindProgramToAdjust(t *testing.T) {
	input := []programSpec{
		{"pbga", 66,nil},
		{"xhth", 57,nil},
		{"ebii", 61,nil},
		{"havc", 66,nil},
		{"ktlj", 57,nil},
		{"fwft", 72,[]string{"ktlj", "cntj", "xhth"}},
		{"qoyq", 66,nil},
		{"padx", 45,[]string{"pbga", "havc", "qoyq"}},
		{"tknk", 41,[]string{"ugml", "padx", "fwft"}},
		{"jptl", 61,nil},
		{"ugml", 41,[]string{"gyxo", "ebii", "jptl"}},
		{"gyxo", 61,nil},
		{"cntj", 57,nil},
	}

	//As you can see, tknk's disc is unbalanced: ugml's stack is heavier than the other two.
	// Even though the nodes above ugml are balanced, ugml itself is too heavy:
	// it needs to be 8 units lighter for its stack to weigh 243 and keep the towers balanced.
	// If this change were made, its weight would be 60.

	node, targetWeight := FindProgramToAdjust(input)
	if node != "ugml" {
		t.Errorf("Wrong target node, expected %s got %s", "ugml", node)
	}
	if targetWeight != 60 {
		t.Errorf("Wrong target weight, expected %d got %d", 60, targetWeight)
	}
}