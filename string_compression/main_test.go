package stringcompression

import (
	"reflect"
	"testing"
)

func TestStringCompression(t *testing.T) {
	cases := []struct {
		result   int
		expected int
	}{
		{compress([]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}), 6},
		{compress([]byte{'a'}), 1},
		{compress(
			[]byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}),
			4,
		},
		{compress([]byte{'a', 'a', 'a', 'b', 'b', 'a', 'a'}), 6},
		{compress([]byte{'a', 'b', 'c'}), 3},
	}
	for _, c := range cases {
		if !reflect.DeepEqual(c.expected, c.result) {
			t.Errorf("Error got: %v, want: %v.", c.result, c.expected)
		}
	}
}
