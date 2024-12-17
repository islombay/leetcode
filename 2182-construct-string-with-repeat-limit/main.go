package main

import (
	"fmt"
)

func main() {
	cases := []struct {
		id          int
		s           string
		repeatLimit int
		want        string
	}{
		{
			id:          1,
			s:           "cczazcc",
			repeatLimit: 3,
			want:        "zzcccac",
		},
		{
			id:          2,
			s:           "aababab",
			repeatLimit: 2,
			want:        "bbabaa",
		},
		{
			id:          3,
			s:           "zccczacc",
			repeatLimit: 3,
			want:        "zzcccacc",
		},
		{
			id:          4,
			s:           "yxxvvvuusrrqqppppoonliihgfeeddcbba",
			repeatLimit: 2,
			want:        "yxxvvuvusrrqqppopponliihgfeeddcbba",
		},
		{
			id:          5,
			s:           "zzzxxxnnn",
			repeatLimit: 1,
			want:        "zxzxzxn",
		},
		{
			id:          6,
			s:           "jjjjjjii",
			repeatLimit: 2,
			want:        "jjijjijj",
		},
	}

	for _, c := range cases {
		if res := repeatLimitedString(c.s, c.repeatLimit); res != c.want {
			fmt.Printf("Test %d failed, expected %s, got %s\n", c.id, c.want, res)
		}
	}
}

func repeatLimitedString(s string, repeatLimit int) string {

	freq := make([]int, 26)

	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
	}

	index := 25
	result := []byte{}

	for index >= 0 {
		if freq[index] == 0 {
			index--
			continue
		}

		k := min(repeatLimit, freq[index])

		for i := 0; i < k; i++ {
			result = append(result, byte(index+'a'))
		}

		freq[index] -= k
		lettersLeft := freq[index]
		if lettersLeft > 0 {
			prevIndex := index - 1

			for prevIndex >= 0 && freq[prevIndex] == 0 {
				prevIndex--
			}

			if prevIndex == -1 {
				break
			}
			result = append(result, byte(prevIndex+'a'))
			freq[prevIndex]--
		}

	}

	return string(result)
}
