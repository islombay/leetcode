package main

import "fmt"

func main() {
	cases := []struct {
		name   string
		height []int
		result int
	}{
		{
			name:   "1",
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			result: 49,
		},
		{
			name:   "2",
			height: []int{1, 1},
			result: 1,
		},
	}

	for _, c := range cases {
		if res := maxArea(c.height); res != c.result {
			fmt.Printf("- Test %s failed, expected %d, got %d\n", c.name, c.result, res)
		} else {
			fmt.Printf("Test %s passed, expected %d, got %d\n", c.name, c.result, res)
		}
	}
}

func maxArea(height []int) int {
	var (
		max   = 0
		left  = 0
		right = len(height) - 1
	)

	for left < right {
		if max < calc(height, left, right) {
			max = calc(height, left, right)
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return max
}

func calc(height []int, left, right int) int {
	if height[left] > height[right] {
		return (right - left) * height[right]
	}
	return (right - left) * height[left]
}
