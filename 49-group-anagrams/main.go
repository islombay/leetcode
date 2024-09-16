package main

import "fmt"

func main() {
	cases := []struct {
		name   string
		strs   []string
		result [][]string
	}{
		{
			name:   "1",
			strs:   []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			result: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
	}

Out:
	for _, c := range cases {
		result := groupAnagrams(c.strs)
		for i := range c.result {
			for j := range c.result[i] {
				if c.result[i][j] != result[i][j] {
					fmt.Printf("- Test %s failed\n", c.name)
					fmt.Println(result)
					fmt.Println(c.result)
					continue Out
				}
			}
		}
		fmt.Printf("Test %s passed\n", c.name)
	}
}

type Key [26]byte

func makeHash(str string) Key {
	key := Key{}
	for i := range str {
		key[str[i]-'a']++
	}
	return key
}

func groupAnagrams(strs []string) [][]string {
	list := map[Key][]string{}

	for _, str := range strs {
		hash := makeHash(str)
		list[hash] = append(list[hash], str)
	}

	var result [][]string
	for _, v := range list {
		result = append(result, v)
	}
	return result
}
