package main

import (
	"reflect"
	"testing"
)

func TestRotLeft(t *testing.T) {
	result := rotLeft([]int32{1, 2, 3, 4, 5}, 4)
	if !reflect.DeepEqual(result, []int32{5, 1, 2, 3, 4}) {
		t.Fail()
	}
}
