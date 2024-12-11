package main

func isPerfectSquare(num int) bool {
	l, r := 1, num
	for l <= r {
		mid := l + (r-l)/2
		midSquare := mid * mid
		if midSquare == num {
			return true
		} else if midSquare < num {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return false
}
