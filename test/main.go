// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strconv"
)

func main() {
	kratnoe()
}

func kratnoe() {
	numbers := map[rune]int{
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'0': 0,
	}
	for i := 0; i < 1001; i++ {
		if i%3 == 0 && i%5 != 0 {
			tmp := strconv.Itoa(i)
			sum := 0
			for _, e := range tmp {
				sum += numbers[e]
			}

			if sum < 10 {
				fmt.Println(i)
			}
		}
	}
}
