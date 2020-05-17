package main

// https://leetcode.com/problems/trapping-rain-water/
// https://leetcode.com/problems/trapping-rain-water/Figures/42/trapping_rain_water.png
func trap(height []int) int {
	sum := 0
	l := len(height)
	if l == 0 {
		return sum
	}
	leftMax := make([]int, l)
	rightMax := make([]int, l)

	leftMax[0] = height[0]
	for i := 1; i < l; i++ {
		leftMax[i] = max(height[i], leftMax[i-1])
	}
	rightMax[l-1] = height[l-1]
	for i := l - 2; i >= 0; i-- {
		rightMax[i] = max(height[i], rightMax[i+1])
	}
	for i := 1; i < l; i++ {
		sum += min(leftMax[i], rightMax[i]) - height[i]
	}
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
