package main

import (
	"reflect"
	"testing"
)

func TestMemory_Reallocate(t *testing.T) {
	tests := []struct {
		allocation    Memory
		expectedAfter Memory
	}{
		{
			Memory{[]int{0, 2, 7, 0}},
			Memory{[]int{2, 4, 1, 2}},
		},
		{
			Memory{[]int{2, 4, 1, 2}},
			Memory{[]int{3, 1, 2, 3}},
		},
		{
			Memory{[]int{3, 1, 2, 3}},
			Memory{[]int{0, 2, 3, 4}},
		},
		{
			Memory{[]int{0, 2, 3, 4}},
			Memory{[]int{1, 3, 4, 1}},
		},
		{
			Memory{[]int{1, 3, 4, 1}},
			Memory{[]int{2, 4, 1, 2}},
		},
	}

	for _, test := range tests {
		m := Memory{make([]int, len(test.allocation.banks))}
		copy(m.banks, test.allocation.banks)
		m.Reallocate()
		if !reflect.DeepEqual(m, test.expectedAfter) {
			t.Errorf("Reallocate %v expected %v got %v", test.allocation, test.expectedAfter, m)
		}
	}
}

func TestCyclesBeforeLoop(t *testing.T) {
	mem := Memory{[]int{0, 2, 7, 0}}
	result := CyclesBeforeLoop(mem)
	if result != 5 {
		t.Errorf("Expected to take 5 cycles, took %d", result)
	}
}
