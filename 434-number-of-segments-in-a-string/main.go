package main

func countSegments(s string) int {
	if s == "" {
		return 0
	}
	prev := s[0]
	count := 0
	for i := 1; i < len(s); i++ {
		if s[i] == ' ' && prev != ' ' {
			count++
		}
		prev = s[i]
	}

	if prev != ' ' {
		count++
	}

	return count
}
