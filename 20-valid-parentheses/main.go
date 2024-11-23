package main

import (
	"fmt"
)

func main() {
	a := "()"
	fmt.Println(isValid(a))
}

type Stack struct {
	Val  rune
	Next *Stack
	Head *Stack
}

func (s *Stack) push(r rune) {
	s.Head = &Stack{
		Val:  r,
		Next: s.Head,
	}
}

func (s *Stack) pop() (rune, bool) {
	cur := s.Head

	if cur == nil {
		return 0, false
	}

	s.Head = s.Head.Next
	return cur.Val, true
}

func isValid(s string) bool {
	charList := map[rune]bool{
		'(': true,
		'{': true,
		'[': true,
	}

	charListOpposite := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := Stack{}

	for _, ch := range s {
		if charList[ch] {
			stack.push(ch)
			continue
		}

		prev, ok := stack.pop()
		if !ok {
			return false
		}
		if prev != charListOpposite[ch] {
			return false
		}
	}

	if _, ok := stack.pop(); ok {
		return false
	}

	return true
}
