package main

import "fmt"

func main() {
	cases := []struct {
		id   int
		x    int
		want int
	}{
		{
			id:   1,
			x:    4,
			want: 2,
		},
		{
			id:   2,
			x:    8,
			want: 2,
		},
		{
			id:   3,
			x:    3,
			want: 1,
		},
		{
			id:   4,
			x:    9,
			want: 3,
		},
	}

	for _, c := range cases {
		if res := mySqrt(c.x); res != c.want {
			fmt.Printf("Test %d failed, expected %d, got %d\n", c.id, c.want, res)
		}
	}
}

func mySqrt(x int) int {
	if x == 0 {
		return x
	}
	l, r := 0, 46340
	nearestNumber := 0
	nearest := 0
	for l <= r {
		mid := l + (r-l)/2
		midValue := mid * mid
		//if abs(x-midValue) < nearest {
		//	nearestNumber = mid
		//	nearest = abs(x - midValue)
		//}
		if midValue == x {
			return mid
		} else if midValue > x {
			r = mid - 1
		} else if midValue < x {
			if midValue > nearest {
				nearest = midValue
				nearestNumber = mid
			}
			l = mid + 1
		}
	}
	return nearestNumber
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
