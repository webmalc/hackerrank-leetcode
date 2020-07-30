package main

import (
	"testing"
)

func TestActivityNotifications(t *testing.T) {
	cases := []struct {
		result   int
		expected int
	}{
		{
			activityNotifications([]int{1, 1, 1, 1, 1}, 3), 0,
		},
		{
			activityNotifications([]int{10, 20, 30, 40, 50}, 3), 1,
		},
		{
			activityNotifications([]int{2, 3, 4, 2, 3, 6, 8, 4, 5}, 5), 2,
		},
		{
			activityNotifications([]int{1, 2, 3, 4, 4}, 4), 0,
		},
	}
	for _, c := range cases {
		if c.result != c.expected {
			t.Errorf("Error got: %v, want: %v.", c.result, c.expected)
		}
	}
}
