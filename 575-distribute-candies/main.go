package main

func distributeCandies(candyType []int) int {
	types := map[int]bool{}
	count := 0
	for _, e := range candyType {
		if !types[e] {
			types[e] = true
			if count < len(candyType)/2 {
				count++
			} else {
				break
			}
		}
	}
	return count
}
