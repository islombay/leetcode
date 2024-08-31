package main

import "fmt"

func main() {
	cases := []struct {
		name    string
		numbers []int
		target  int
		result  []int
	}{
		{
			name:    "1",
			numbers: []int{2, 7, 11, 15},
			target:  9,
			result:  []int{1, 2},
		},
		{
			name:    "2",
			numbers: []int{2, 3, 4},
			target:  6,
			result:  []int{1, 3},
		},
		{
			name:    "3",
			numbers: []int{-1, 0},
			target:  -1,
			result:  []int{1, 2},
		},
	}
out:
	for _, c := range cases {
		r := twoSum(c.numbers, c.target)
		for i := 0; i < len(r); i++ {
			if r[i] != c.result[i] {
				fmt.Printf("- Test %s failed, expected %v, got %v\n", c.name, c.result, r)
				continue out
			}
		}
		fmt.Printf("Test %s passed, expected %v, got %v\n", c.name, c.result, r)
	}
}

func twoSum(numbers []int, target int) []int {
	if len(numbers) == 2 {
		return []int{1, 2}
	}

	right := len(numbers) - 1
	for left := 0; left < len(numbers); left++ {
		for right > left && numbers[left]+numbers[right] != target {
			right--
		}

		if numbers[left]+numbers[right] == target {
			return []int{left + 1, right + 1}
		}
		right = len(numbers) - 1
	}

	return []int{-1, -1}
}
