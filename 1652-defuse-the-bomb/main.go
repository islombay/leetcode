package main

func decrypt(code []int, k int) []int {
	result := make([]int, len(code))
	if k == 0 {
		return result
	}

	n := len(code)
	// 2 4 9 3

	for i := 0; i < len(code); i++ {
		l, r := i+1, i+k
		if k < 0 {
			l, r = n+i+k, n+i-1
		}

		for l <= r {
			result[i] += code[l%n]
			l++
		}
	}

	return result
}
