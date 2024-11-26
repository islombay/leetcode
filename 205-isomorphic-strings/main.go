package main

import "fmt"

func main() {
	cases := []struct {
		id   int
		s, t string
		want bool
	}{
		{
			id:   1,
			s:    "egg",
			t:    "add",
			want: true,
		},
		{
			id:   2,
			s:    "paper",
			t:    "title",
			want: true,
		},
		{
			id:   3,
			s:    "foo",
			t:    "bar",
			want: false,
		},
	}

	for _, c := range cases {
		if isIsomorphic(c.s, c.t) != c.want {
			fmt.Printf("Test %d failed, expected %T, got %T\n", c.id, c.want, !c.want)
		}
	}
	fmt.Println("Finished")
}

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mapST := make(map[rune]rune)
	mapTS := make(map[rune]rune)

	for i, r := range s {
		charT := rune(t[i])

		if val, ok := mapST[r]; ok {
			if val != charT {
				return false
			}
		} else {
			mapST[r] = charT
		}

		if val, ok := mapTS[charT]; ok {
			if val != r {
				return false
			}
		} else {
			mapTS[charT] = r
		}
	}
	return true
}
