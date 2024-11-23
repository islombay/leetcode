package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "/.../a/../b/c/../d/./"
	fmt.Println(simplifyPath(s))
}

type Stack struct {
	Val  string
	Head *Stack
	Next *Stack
}

func (s *Stack) Pop() {
	if s.Head != nil {
		s.Head = s.Head.Next
	}
}

func (s *Stack) Push(path string) {
	s.Head = &Stack{
		Val:  path,
		Next: s.Head,
	}
}

func (s *Stack) Top() (string, bool) {
	if s.Head != nil {
		return s.Head.Val, true
	}
	return "", false
}

func simplifyPath(path string) string {
	pathArr := strings.Split(path, "/")

	stack := Stack{}

	for _, i := range pathArr {
		if i == "" || i == "." {
			continue
		}

		if i == ".." {
			stack.Pop()
			continue
		}

		stack.Push(i)
	}

	path = ""

	val, ok := stack.Top()
	for ok {
		path = "/" + val + path
		stack.Pop()
		val, ok = stack.Top()
	}

	if path == "" {
		path = "/"
	}

	return path
}
