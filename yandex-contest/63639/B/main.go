package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var N, Q int
	fmt.Fscan(reader, &N, &Q)

	words := make([]string, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &words[i])
	}

	for i := 0; i < Q; i++ {
		var k int
		var prefix string
		fmt.Fscan(reader, &k, &prefix)

		start := lowerBound(words, prefix)
		end := upperBound(words, prefix)

		if start+k-1 < end {
			fmt.Println(start + k)
		} else {
			fmt.Println(-1)
		}
	}
}

func lowerBound(words []string, prefix string) int {
	low, high := 0, len(words)
	for low < high {
		mid := (low + high) / 2
		if strings.HasPrefix(words[mid], prefix) {
			high = mid
		} else if words[mid] < prefix {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}

func upperBound(words []string, prefix string) int {
	low, high := 0, len(words)
	for low < high {
		mid := (low + high) / 2
		if strings.HasPrefix(words[mid], prefix) {
			low = mid + 1
		} else if words[mid] < prefix {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}
