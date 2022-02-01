package patholes

import (
	"sort"
)

func patholes(road string, budget int) int {
	x := 'x'
	counter := 0
	values := []int{}
	result := 0
	for _, l := range road {
		if l == x {
			counter++
		} else {
			if counter != 0 {
				values = append(values, counter)
			}
			counter = 0
		}
	}
	if counter != 0 {
		values = append(values, counter)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(values)))

	for _, v := range values {
		b := v + 1
		if b < budget {
			budget -= b
			result += v
		} else {
			result += budget - 1
			break
		}
	}
	return result
}
