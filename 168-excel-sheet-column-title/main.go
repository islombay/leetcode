package main

import "fmt"

func main() {
	cases := []struct {
		id           int
		columnNumber int
		want         string
	}{
		{
			id:           1,
			columnNumber: 1,
			want:         "A",
		},
		{
			id:           2,
			columnNumber: 28,
			want:         "AB",
		},
		{
			id:           3,
			columnNumber: 701,
			want:         "ZY",
		},
		{
			id:           4,
			columnNumber: 2147483647,
			want:         "FXSHRXW",
		},
		{
			id:           5,
			columnNumber: 52,
			want:         "AZ",
		},
	}

	for _, c := range cases {
		if res := convertToTitle(c.columnNumber); res != c.want {
			fmt.Printf("Test %d failed, expected %s, got %s\n", c.id, c.want, res)
		}
	}
}

func convertToTitle(columnNumber int) string {
	m := map[int]rune{
		1:  'A',
		2:  'B',
		3:  'C',
		4:  'D',
		5:  'E',
		6:  'F',
		7:  'G',
		8:  'H',
		9:  'I',
		10: 'J',
		11: 'K',
		12: 'L',
		13: 'M',
		14: 'N',
		15: 'O',
		16: 'P',
		17: 'Q',
		18: 'R',
		19: 'S',
		20: 'T',
		21: 'U',
		22: 'V',
		23: 'W',
		24: 'X',
		25: 'Y',
		26: 'Z',
	}

	res := ""
	for {
		tmp := string(m[columnNumber%26])
		if tmp == "\x00" {
			tmp = string(m[26])
		}
		res = tmp + res
		if columnNumber == 26 {
			break
		}
		if columnNumber%26 == 0 {
			columnNumber = columnNumber/26 - 1
		} else {
			columnNumber = columnNumber / 26
		}
		if columnNumber == 0 {
			break
		}
	}
	return res
}
