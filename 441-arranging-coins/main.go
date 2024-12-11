package main

func arrangeCoins(n int) int {
	l, r := 1, n
	ans := -1
	for l <= r {
		mid := l + (r-l)/2
		midVal := (mid * (mid + 1)) / 2
		if midVal > n {
			r = mid - 1
		} else {
			ans = mid
			l = mid + 1
		}
	}
	return ans
}
