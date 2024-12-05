package main

import "fmt"

func main() {
	cases := []struct {
		id    int
		words []string
		want  []string
	}{
		{
			id:    1,
			words: []string{"Hello", "Alaska", "DAD", "Peace"},
			want:  []string{"Alaska", "DAD"},
		},
		{
			id:    2,
			words: []string{"okm"},
			want:  []string{},
		},
		{
			id:    3,
			words: []string{"adsdf", "sdf"},
			want:  []string{"adsdf", "sdf"},
		},
	}

	for _, c := range cases {
		res := findWords(c.words)
		if len(res) != len(c.want) {
			fmt.Printf("Test %d failed len\n", c.id)
			continue
		}

		for i := 0; i < len(c.want); i++ {
			if res[i] != c.want[i] {
				fmt.Printf("Test %d failed len\n", c.id)
				break
			}
		}
	}
}

func findWords(words []string) []string {
	keyboardMap := map[rune]int{
		'q': 1, 'w': 1, 'e': 1, 'r': 1, 't': 1, 'y': 1, 'u': 1, 'i': 1, 'o': 1, 'p': 1,
		'Q': 1, 'W': 1, 'E': 1, 'R': 1, 'T': 1, 'Y': 1, 'U': 1, 'I': 1, 'O': 1, 'P': 1,
		'a': 2, 's': 2, 'd': 2, 'f': 2, 'g': 2, 'h': 2, 'j': 2, 'k': 2, 'l': 2,
		'A': 2, 'S': 2, 'D': 2, 'F': 2, 'G': 2, 'H': 2, 'J': 2, 'K': 2, 'L': 2,
		'z': 3, 'x': 3, 'c': 3, 'v': 3, 'b': 3, 'n': 3, 'm': 3,
		'Z': 3, 'X': 3, 'C': 3, 'V': 3, 'B': 3, 'N': 3, 'M': 3,
	}

	result := []string{}

Outer:
	for _, word := range words {
		row := keyboardMap[rune(word[0])]
		for _, r := range word {
			if keyboardMap[r] != row {
				continue Outer
			}
		}
		result = append(result, word)
	}
	return result
}
