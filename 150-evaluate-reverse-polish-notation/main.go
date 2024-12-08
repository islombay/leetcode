package main

import (
	"container/list"
	"fmt"
	"strconv"
)

func main() {
	cases := []struct {
		id     int
		tokens []string
		want   int
	}{
		{
			id:     1,
			tokens: []string{"2", "1", "+", "3", "*"},
			want:   9,
		},
		{
			id:     2,
			tokens: []string{"4", "13", "5", "/", "+"},
			want:   6,
		},
		{
			id:     3,
			tokens: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			want:   22,
		},
	}

	for _, c := range cases {
		if res := evalRPN(c.tokens); res != c.want {
			fmt.Printf("Test %d failed, expected %d, got %d\n", c.id, c.want, res)
		}
	}
}

func evalRPN(tokens []string) int {
	stack := list.New()

	for _, token := range tokens {
		if token == "+" || token == "-" || token == "*" || token == "/" {
			last := stack.Remove(stack.Back()).(int)
			prelast := stack.Remove(stack.Back()).(int)
			if token == "+" {
				stack.PushBack(prelast + last)
			} else if token == "-" {
				stack.PushBack(prelast - last)
			} else if token == "/" {
				stack.PushBack(prelast / last)
			} else {
				stack.PushBack(prelast * last)
			}
			continue
		}
		n, _ := strconv.Atoi(token)
		stack.PushBack(n)
	}
	return stack.Back().Value.(int)
}
