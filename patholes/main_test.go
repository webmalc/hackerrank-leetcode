package patholes

import (
	"testing"
)

func TestPatholes(t *testing.T) {
	cases := []struct {
		result   int
		expected int
	}{
		{patholes("x.x..x.x.x.x.x.x.x", 7), 3},
		{patholes("...xxxxxx..x..x.x.x.x.", 7), 6},
		{patholes("...xxxx.....", 7), 4},
		{patholes("...xxxx....xxx.", 7), 5},
		{patholes("...xxx..x....xxx.", 7), 5},
		{patholes("...xxx..x....xxx.", 4), 3},
		{patholes("...xxx..x....xxx.", 3), 2},
		{patholes("..xxxxx", 4), 3},
		{patholes("..", 4), 0},
	}
	for _, c := range cases {
		if c.expected != c.result {
			t.Errorf("Error got: %d, want: %d.", c.result, c.expected)
		}
	}
}
