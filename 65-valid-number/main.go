package main

import "fmt"

func main() {
	cases := []struct {
		id   int
		s    string
		want bool
	}{
		{id: 1, s: "2", want: true},
		{id: 2, s: "0089", want: true},
		{id: 3, s: "-0.1", want: true},
		{id: 4, s: "+3.14", want: true},
		{id: 5, s: "4.", want: true},
		{id: 6, s: "-.9", want: true},
		{id: 7, s: "2e10", want: true},
		{id: 8, s: "-90E3", want: true},
		{id: 10, s: "+6e-1", want: true},
		{id: 11, s: "53.5e93", want: true},
		{id: 12, s: "-123.456e789", want: true},
		{id: 13, s: "abc", want: false},
		{id: 14, s: "1a", want: false},
		{id: 15, s: "1e", want: false},
		{id: 16, s: "e3", want: false},
		{id: 17, s: "99e2.5", want: false},
		{id: 18, s: "--6", want: false},
		{id: 19, s: "-+3", want: false},
		{id: 20, s: "95a54e53", want: false},
		{id: 21, s: ".e1", want: false},
		{id: 22, s: "+.", want: false},
		{id: 23, s: ".-4", want: false},
		{id: 24, s: "+E3", want: false},
		{id: 25, s: "46.e3", want: true},
	}

	for _, c := range cases {
		if res := isNumber(c.s); res != c.want {
			fmt.Printf("Test %d failed, expected %t, got %t\n", c.id, c.want, res)
		}
	}
}

func isNumber(s string) bool {
	m := map[rune]bool{
		'0': true,
		'1': true,
		'2': true,
		'3': true,
		'4': true,
		'5': true,
		'6': true,
		'7': true,
		'8': true,
		'9': true,
		'e': true,
		'E': true,
		'-': true,
		'+': true,
		'.': true,
	}
	nums := map[rune]bool{
		'0': true,
		'1': true,
		'2': true,
		'3': true,
		'4': true,
		'5': true,
		'6': true,
		'7': true,
		'8': true,
		'9': true,
	}

	if len(s) == 1 {
		return nums[rune(s[0])]
	}

	if rune(s[0]) == 'e' || rune(s[0]) == 'E' || rune(s[len(s)-1]) == 'e' || rune(s[len(s)-1]) == 'E' || rune(s[len(s)-1]) == '+' || rune(s[len(s)-1]) == '-' {
		return false
	}

	signOccur := false
	signOccurIndex := 0
	eOccur := false
	eOccurIndex := 0
	dotOccur := false
	digitOccur := false
	dotOccurIndex := 0
	for i, r := range s {
		if !m[r] {
			return false
		} else if nums[r] {
			digitOccur = true
			continue
		} else if r == '-' || r == '+' {
			if i == 0 {
				signOccur = true
				signOccurIndex = i
				continue
			} else if (signOccur && i-signOccurIndex == 1) || (!eOccur && digitOccur) || (eOccur && i-eOccurIndex != 1) || (dotOccur && i-dotOccurIndex == 1) {
				return false
			}
			continue
		} else if r == '.' {
			if eOccur || dotOccur {
				return false
			}
			dotOccurIndex = i
			dotOccur = true
		} else if r == 'e' || r == 'E' {
			if eOccur || !digitOccur || (signOccur && i-signOccurIndex == 1) {
				return false
			}
			eOccur = true
			eOccurIndex = i
		}
	}
	return !(dotOccur && !digitOccur)
	//if dotOccur && !digitOccur {
	//	return false
	//}
	//return true
}
