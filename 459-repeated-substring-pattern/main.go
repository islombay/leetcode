package main

func repeatedSubstringPattern(s string) bool {
	if len(s) == 1 {
		return false
	}

	// abab

	for i := 1; i <= len(s)/2; i++ { // i = 1
		found := false
		if len(s)%i == 0 {
			j := i + i
			for ; j <= len(s); j += i { // j = 4
				if s[0:i] != s[j-i:j] {
					found = false
					break
				}
				found = true
			}
			if j-i != len(s) {
				found = false
			}
		}

		if found {
			return true
		}
	}
	return false
}
