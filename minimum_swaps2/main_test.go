package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	first := minimumSwaps([]int32{4, 3, 1, 2})
	second := minimumSwaps([]int32{2, 3, 4, 1, 5})
	if first != 3 {
		t.Fail()
		t.Errorf("Error got: %d, want: %d.", first, 3)
	}

	if second != 3 {
		t.Fail()
		t.Errorf("Error got: %d, want: %d.", second, 3)
	}
}
