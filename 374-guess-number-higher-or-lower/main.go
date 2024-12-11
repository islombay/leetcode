package main

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	l, r := 1, n
	for l <= r {
		mid := l + (r-l)/2
		res := guess(mid)
		if res == 0 {
			return mid
		} else if res == 1 {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return 1
}

func guess(num int) int {
	return num
}
